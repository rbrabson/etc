package ftc

import (
	"fmt"
	"testing"
)

// TestGetTeams calls ftc.GetTeams to test the results of the call
func TestGetTeams(t *testing.T) {
	results, err := GetTeams("2023")
	if err != nil {
		t.Fatalf("TestGetTeams: %s", err.Error())
	}
	fmt.Println(results)
	t.Log(results)
}

// TestGetTeamsWithTeamNumber calls ftc.GetTeams to test the results of the call
func TestGetTeamsWithTeamNumber(t *testing.T) {
	results, err := GetTeams("2023", "7083")
	if err != nil {
		t.Fatalf("TestGetTeamsWithTeamNumber: %s", err.Error())
	}
	t.Log(results)
}
