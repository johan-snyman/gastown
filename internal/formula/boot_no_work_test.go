package formula

import (
	"strings"
	"testing"
)

func TestBootTriageFormulaNoWorkIsSilentNoop(t *testing.T) {
	content, err := formulasFS.ReadFile("formulas/mol-boot-triage.formula.toml")
	if err != nil {
		t.Fatalf("reading boot triage formula: %v", err)
	}

	text := string(content)
	for _, want := range []string{
		"no work available",
		"NOTHING and exit cleanly without escalation",
		"Do not call `gt escalate` for \"No work available\"",
		"exit directly with no mail",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("boot triage formula missing %q", want)
		}
	}
}
