package ftc

import (
	"testing"
)

// TestGetMatcheResults calls ftc.GetMatcheResults to test the results of the call
func TestGetMatcheResults(t *testing.T) {
	results, err := GetMatchResults("2023", "USNCCMP", PLAYOFF)
	if err != nil {
		t.Fatalf("TestGetMatcheResults: %s", err.Error())
	}
	t.Log(results)
}

// TestGetMatcheResultsWithTeamNumber calls ftc.GetMatcheResults to test the results of the call
func TestGetMatcheResultsWithTeamNumber(t *testing.T) {
	results, err := GetMatchResults("2023", "USNCCMP", PLAYOFF, "7083")
	if err != nil {
		t.Fatalf("TestGetMatcheResultsWithTeamNumber: %s", err.Error())
	}
	t.Log(results)
}
