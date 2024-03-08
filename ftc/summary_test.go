package ftc

import "testing"

// TestGetSeasonSummary calls ftc.GetSeasonSummary to test the results of the call
func TestGetSeasonSummary(t *testing.T) {
	results, err := GetSeasonSummary("2023")
	if err != nil {
		t.Fatalf("TestGetSeasonSummary: %s", err.Error())
	}
	t.Log(results)
}
