package config

// Default agent configuration values.
// These are the single source of truth — all fallback/default logic should reference these
// instead of hardcoding numeric literals.
const (
	DefaultContextWindow   = 200000
	DefaultMaxTokens       = 8192
	DefaultMaxMessageChars = 32000
	DefaultMaxIterations   = 30
	DefaultTemperature     = 0.7
	DefaultHistoryShare = 0.85

	// Tool-call budget. 0 at runtime means "no cap" for parallel calls and
	// "no per-result trim" for tokens. Loop detector zeros fall back to the
	// historical 3/5 and 4/6 constants.
	DefaultMaxToolCalls              = 25
	DefaultMaxParallelToolCalls      = 3
	DefaultToolResultMaxTokens       = 1500 // ~6000 soft-trim chars at 4 chars/token
	DefaultToolLoopSameCallWarning   = 3
	DefaultToolLoopSameCallCritical  = 5
	DefaultToolLoopSameResultWarning = 4
	DefaultToolLoopSameResultCritical = 6
)
