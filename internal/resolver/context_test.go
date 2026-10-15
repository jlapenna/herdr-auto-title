package resolver

import (
	"testing"

	"github.com/kryptamine/herdr-auto-title/internal/herdr"
	"github.com/kryptamine/herdr-auto-title/internal/state"
)

func hiddenContextChain() *Deterministic {
	return Default(Options{
		MaxLength:     DefaultMaxLength,
		BranchMax:     DefaultBranchMaxLength,
		HideAgentName: true,
		HideContext:   true,
	})
}

func TestAHiddenContextLeavesOnlyTheWork(t *testing.T) {
	t.Parallel()

	got := hiddenContextChain().Resolve(tabWithPane(&state.PaneState{
		Dir:         dashboard,
		Agent:       "claude",
		AgentStatus: herdr.AgentStatusWorking,
		AgentTitle:  "Implement OAuth scopes",
	}))

	if want := "Implement OAuth scopes"; got.Name != want {
		t.Errorf("name = %q, want %q", got.Name, want)
	}
}

func TestAHiddenContextNeverFallsBackToTheDirectory(t *testing.T) {
	t.Parallel()

	got := hiddenContextChain().Resolve(tabWithPane(&state.PaneState{Dir: dashboard}))

	if want := "Shell"; got.Name != want {
		t.Errorf("name = %q, want %q", got.Name, want)
	}
}
