package agent

import (
	"encoding/json"
	"testing"

	"github.com/nextlevelbuilder/goclaw/internal/config"
	"github.com/nextlevelbuilder/goclaw/internal/store"
)

func TestResolveToolBudget_OverrideBeatsDefaults(t *testing.T) {
	cfg := config.Default()
	ag := &store.AgentData{OtherConfig: json.RawMessage(`{
		"max_tool_calls": 40,
		"max_parallel_tool_calls": 2,
		"tool_result_max_tokens": 800,
		"tool_loop_same_call_warning": 10,
		"tool_loop_same_call_critical": 12,
		"tool_loop_same_result_warning": 8,
		"tool_loop_same_result_critical": 10
	}`)}
	got := resolveToolBudget(cfg, "agent-a", ag.ParseToolBudget())
	if got.MaxToolCalls != 40 || got.MaxParallelToolCalls != 2 || got.ToolResultMaxTokens != 800 {
		t.Fatalf("override = %+v", got)
	}
	if got.ToolLoopSameCallWarning != 10 || got.ToolLoopSameCallCritical != 12 {
		t.Fatalf("same-call limits = %d/%d", got.ToolLoopSameCallWarning, got.ToolLoopSameCallCritical)
	}
	if got.ToolLoopSameResultWarning != 8 || got.ToolLoopSameResultCritical != 10 {
		t.Fatalf("same-result limits = %d/%d", got.ToolLoopSameResultWarning, got.ToolLoopSameResultCritical)
	}
}

func TestResolveToolBudget_MissingOverrideUsesDefaults(t *testing.T) {
	ag := &store.AgentData{OtherConfig: json.RawMessage(`{"prompt_mode":"full"}`)}
	got := resolveToolBudget(nil, "agent-a", ag.ParseToolBudget())
	if got.MaxToolCalls != config.DefaultMaxToolCalls {
		t.Fatalf("MaxToolCalls = %d, want %d", got.MaxToolCalls, config.DefaultMaxToolCalls)
	}
	if got.MaxParallelToolCalls != config.DefaultMaxParallelToolCalls {
		t.Fatalf("MaxParallelToolCalls = %d", got.MaxParallelToolCalls)
	}
	if got.ToolResultMaxTokens != config.DefaultToolResultMaxTokens {
		t.Fatalf("ToolResultMaxTokens = %d", got.ToolResultMaxTokens)
	}
	if got.ToolLoopSameCallWarning != 3 || got.ToolLoopSameCallCritical != 5 {
		t.Fatalf("same-call defaults = %d/%d", got.ToolLoopSameCallWarning, got.ToolLoopSameCallCritical)
	}
	if got.ToolLoopSameResultWarning != 4 || got.ToolLoopSameResultCritical != 6 {
		t.Fatalf("same-result defaults = %d/%d", got.ToolLoopSameResultWarning, got.ToolLoopSameResultCritical)
	}

	loop := NewLoop(LoopConfig{
		MaxToolCalls:               got.MaxToolCalls,
		MaxParallelToolCalls:       got.MaxParallelToolCalls,
		ToolResultMaxTokens:        got.ToolResultMaxTokens,
		ToolLoopSameCallWarning:    got.ToolLoopSameCallWarning,
		ToolLoopSameCallCritical:   got.ToolLoopSameCallCritical,
		ToolLoopSameResultWarning:  got.ToolLoopSameResultWarning,
		ToolLoopSameResultCritical: got.ToolLoopSameResultCritical,
	})
	if loop.maxToolCalls != 25 || loop.maxParallelToolCalls != 3 || loop.toolResultMaxTokens != config.DefaultToolResultMaxTokens {
		t.Fatalf("loop budget = calls %d parallel %d tokens %d", loop.maxToolCalls, loop.maxParallelToolCalls, loop.toolResultMaxTokens)
	}
}
