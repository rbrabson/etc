package ftc

import "testing"

// TestGetEventScores calls ftc.GetEventScores to test the results of the call
func TestGetEventScores(t *testing.T) {
	results, err := GetEventScores("2023", "USNCCMP", PLAYOFF)
	if err != nil {
		t.Fatalf("TestGetEventMatchResults: %s", err.Error())
	}
	t.Log(results)
}

// TestGetEventScoresWithTeamNumber calls ftc.GetEventScores to test the results of the call
func TestGetEventScoresWithTeamNumber(t *testing.T) {
	results, err := GetEventScores("2023", "USNCCMP", PLAYOFF, "7083")
	if err != nil {
		t.Fatalf("TestGetEventMatchResultsWithTeamNumber: %s", err.Error())
	}
	t.Log(results)
}
