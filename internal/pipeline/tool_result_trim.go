package pipeline

import (
	"fmt"
	"unicode/utf8"

	"github.com/nextlevelbuilder/goclaw/internal/providers"
)

const (
	charsPerTokenEstimate = 4
	// defaultToolResultTokens matches soft-trim maxChars 6000.
	defaultToolResultTokens = 1500
	deferredToolResultText  = "Deferred by GoClaw tool-call budget. Re-issue this call on the next turn if still required."
	compactionToolPlaceholder = "[Tool result omitted after compaction to fit context budget.]"
)

// mediaToolNames matches agent.mediaToolNames. Media descriptions keep a larger
// budget (8000 chars vs 6000) because regenerating them costs another model call.
var mediaToolNames = map[string]bool{
	"read_image":    true,
	"read_document": true,
	"read_audio":    true,
	"read_video":    true,
}

func splitToolCallsByCap(calls []providers.ToolCall, cap int) (execute, deferred []providers.ToolCall) {
	if cap <= 0 || len(calls) <= cap {
		return calls, nil
	}
	return calls[:cap], calls[cap:]
}

func deferredToolMessage(tc providers.ToolCall) providers.Message {
	return providers.Message{
		Role:       "tool",
		Content:    deferredToolResultText,
		ToolCallID: tc.ID,
		ToolName:   tc.Name,
	}
}

func (s *ToolStage) appendDeferredPlaceholders(state *RunState, deferred []providers.ToolCall) {
	for _, tc := range deferred {
		state.Messages.AppendPending(deferredToolMessage(tc))
	}
}

func (s *ToolStage) trimToolMessages(state *RunState, msgs []providers.Message) []providers.Message {
	limit := s.deps.Config.ToolResultMaxTokens
	if limit <= 0 || len(msgs) == 0 {
		return msgs
	}
	out := make([]providers.Message, len(msgs))
	model := ""
	if state != nil {
		model = state.Model
	}
	for i, msg := range msgs {
		if msg.Role == "tool" && msg.Content != "" {
			msg.Content = trimToolResultContent(msg.Content, msg.ToolName, limit, s.deps, model)
		}
		out[i] = msg
	}
	return out
}

func trimToolResultContent(content, toolName string, maxTokens int, deps *PipelineDeps, model string) string {
	if maxTokens <= 0 || content == "" {
		return content
	}
	limit := tokenLimitForTool(toolName, maxTokens)
	if estimateContentTokens(content, deps, model) <= limit {
		return content
	}
	return headTailTrim(content, limit)
}

func tokenLimitForTool(toolName string, maxTokens int) int {
	if mediaToolNames[toolName] {
		// 8000/6000 of the configured budget.
		scaled := maxTokens * 4 / 3
		if scaled < maxTokens {
			return maxTokens
		}
		return scaled
	}
	return maxTokens
}

func estimateContentTokens(content string, deps *PipelineDeps, model string) int {
	if deps != nil && deps.TokenCounter != nil {
		return deps.TokenCounter.Count(model, content)
	}
	n := utf8.RuneCountInString(content)
	if n == 0 {
		return 0
	}
	tokens := n / charsPerTokenEstimate
	if tokens < 1 {
		return 1
	}
	return tokens
}

func headTailTrim(content string, maxTokens int) string {
	charBudget := maxTokens * charsPerTokenEstimate
	if charBudget < 64 {
		charBudget = 64
	}
	headN := charBudget / 2
	tailN := charBudget - headN
	total := utf8.RuneCountInString(content)
	trimmed := fmt.Sprintf("%s\n...\n%s\n\n[Tool result trimmed: kept first %d chars and last %d chars of %d chars.]",
		takeHead(content, headN), takeTail(content, tailN), headN, tailN, total)
	return trimmed
}

func takeHead(s string, n int) string {
	if n <= 0 {
		return ""
	}
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n])
}

func takeTail(s string, n int) string {
	if n <= 0 {
		return ""
	}
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[len(r)-n:])
}

// shrinkToolResults reduces recent tool results one step after compaction is
// still over the history budget. Returns true when any content changed.
func shrinkToolResults(state *RunState, configuredMaxTokens int) bool {
	if state == nil || state.Messages == nil {
		return false
	}
	limit := configuredMaxTokens
	if limit <= 0 {
		limit = defaultToolResultTokens
	}
	limit /= 2
	if limit < 64 {
		limit = 64
	}
	changed := false
	changed = shrinkToolSlice(state.Messages.Pending(), limit) || changed
	changed = shrinkToolSlice(state.Messages.History(), limit) || changed
	return changed
}

func shrinkToolSlice(msgs []providers.Message, limit int) bool {
	changed := false
	for i := len(msgs) - 1; i >= 0; i-- {
		if msgs[i].Role != "tool" || msgs[i].Content == "" {
			continue
		}
		next := shrinkOneToolResult(msgs[i].Content, msgs[i].ToolName, limit)
		if next != msgs[i].Content {
			msgs[i].Content = next
			changed = true
		}
	}
	return changed
}

func shrinkOneToolResult(content, toolName string, limit int) string {
	if content == compactionToolPlaceholder {
		return content
	}
	toolLimit := tokenLimitForTool(toolName, limit)
	runeLimit := toolLimit * charsPerTokenEstimate
	if utf8.RuneCountInString(content) > runeLimit {
		return headTailTrim(content, toolLimit)
	}
	return compactionToolPlaceholder
}
