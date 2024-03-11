package ftc

import (
	"testing"
)

// TestGetLeagues calls ftc.GetEvents to test the results of the call
func TestGetLeagues(t *testing.T) {
	leagues, err := GetLeagues("2023")
	if err != nil {
		t.Fatalf("TestGetEvents: %s", err.Error())
	}
	t.Log(leagues)
}

// TestGetLeaguesWithRegionCode calls ftc.GetLeaguesWithRegionCode to test the results of the call
func TestGetLeaguesWithRegionCode(t *testing.T) {
	qparms := map[string]string{
		"regionCode": "USTX",
	}
	leagues, err := GetLeagues("2023", qparms)
	if err != nil {
		t.Fatalf("TestGetLeaguesWithRegionCode: %s", err.Error())
	}
	t.Log(leagues)
}

// TestGetLeaguesWithLeagueCode calls ftc.GetLeaguesWithLeagueCode to test the results of the call
func TestGetLeaguesWithLeagueCode(t *testing.T) {
	qparms := map[string]string{
		"leagueCode": "CG",
	}
	leagues, err := GetLeagues("2023", qparms)
	if err != nil {
		t.Fatalf("TestGetLeaguesWithLeagueCode: %s", err.Error())
	}
	t.Log(leagues)
}

// TestGetLeaguesWithLRegionCodeAndeagueCode calls ftc.GetLeaguesWithLRegionCodeAndeagueCode to test the results of the call
func TestGetLeaguesWithLRegionCodeAndeagueCode(t *testing.T) {
	qparms := map[string]string{
		"regionCode": "USTX",
		"leagueCode": "CG",
	}
	leagues, err := GetLeagues("2023", qparms)
	if err != nil {
		t.Fatalf("TestGetLeaguesWithLRegionCodeAndeagueCode: %s", err.Error())
	}
	t.Log(leagues)
}

// TestGetLeagueMembers calls ftc.GetLeagueMembers to test the results of the call
func TestGetLeagueMembers(t *testing.T) {
	members, err := GetLeagueMembers("2023", "USTX", "CG")
	if err != nil {
		t.Fatalf("TestGetLeagueMembers: %s", err.Error())
	}
	t.Log(members)
}
