package ftc

import "testing"

// TestGetRankings calls ftc.GetRankings to test the results of the call
func TestGetRankings(t *testing.T) {
	members, err := GetRankings("2023", "USNCCMP")
	if err != nil {
		t.Fatalf("TestGetLeagueRankings: %s", err.Error())
	}
	t.Log(members)
}
