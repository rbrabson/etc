package ftc

import (
	"testing"
)

// TestGetAwardListing calls ftc.GetAwardListing to test the results of the call
func TestGetAwardListing(t *testing.T) {
	alliances, err := GetAwardListing("2023")
	if err != nil {
		t.Fatalf("TestGetAwardListing: %s", err.Error())
	}
	t.Log(alliances)
}

// TestGetEventAwards calls ftc.func GetEventAwards to test the results of the call
func TestGetEventAwards(t *testing.T) {
	awards, err := GetEventAwards("2023", "USNCCMP")
	if err != nil {
		t.Fatalf("TestGetEventAwards: %s", err.Error())
	}
	t.Log(awards)
}

// TestGetEventAwards calls ftc.func GetEventAwards to test the results of the call
func TestGetEventAwardsWithTeam(t *testing.T) {
	awards, err := GetEventAwards("2023", "USNCCMP", "7083")
	if err != nil {
		t.Fatalf("TestGetEventAwardsWithTeam: %s", err.Error())
	}
	t.Log(awards)
}

// TestGetEventAwards calls ftc.func GetEventAwards to test the results of the call
func TestTeamAwards(t *testing.T) {
	awards, err := GetEventAwards("2023", "7083")
	if err != nil {
		t.Fatalf("TestTeamAwards: %s", err.Error())
	}
	t.Log(awards)
}

// TestGetEventAwards calls ftc.func GetEventAwards to test the results of the call
func TestGetTeamAwardsWithEvent(t *testing.T) {
	awards, err := GetEventAwards("2023", "7083", "USNCCMP")
	if err != nil {
		t.Fatalf("TestGetTeamAwardsWithEvent: %s", err.Error())
	}
	t.Log(awards)
}
