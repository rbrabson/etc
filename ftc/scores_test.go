package ftc

import "testing"

// TestGetGetEventMatchResults calls ftc.GetEventMatchResults to test the results of the call
func TestGetEventMatchResults(t *testing.T) {
	results, err := GetEventMatchResults("2023", "USNCCMP", PLAYOFF)
	if err != nil {
		t.Fatalf("TestGetEventMatchResults: %s", err.Error())
	}
	t.Log(results)
}

// TestGetGetEventMatchResults calls ftc.GetEventMatchResults to test the results of the call
func TestGetEventMatchResultsWithTeamNumber(t *testing.T) {
	results, err := GetEventMatchResults("2023", "USNCCMP", PLAYOFF, "7083")
	if err != nil {
		t.Fatalf("TestGetEventMatchResultsWithTeamNumber: %s", err.Error())
	}
	t.Log(results)
}
