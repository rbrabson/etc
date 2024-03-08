package ftc

import (
	"testing"
)

// TestGetEventAlliances calls ftc.TestGetEventAlliances to test the results of the call
func TestGetEventAlliances(t *testing.T) {
	alliances, err := GetEventAlliances("2023", "USNCCMP")
	if err != nil {
		t.Fatalf("TestGetEventAlliances: %s", err.Error())
	}
	t.Log(alliances)
}

// TestGetAllianceSelections calls ftc.GetAllianceSelections to test the results of the call
func TestGetAllianceSelections(t *testing.T) {
	allianceSelections, err := GetAllianceSelections("2023", "USNCCMP")
	if err != nil {
		t.Fatalf("TestGetAllianceSelections: %s", err.Error())
	}
	t.Log(allianceSelections)
}
