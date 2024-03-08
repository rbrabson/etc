package ftc

import (
	"testing"
)

// TestGetEvents calls ftc.GetEvents to test the results of the call
func TestGetEvents(t *testing.T) {
	events, err := GetEvents("2023")
	if err != nil {
		t.Fatalf("TestGetEvents: %s", err.Error())
	}
	t.Log(events)
}

// TestGetEvents calls ftc.GetEvents to test the results of the call
func TestGetEventsWithEventCode(t *testing.T) {
	qparms := map[string]string{
		"eventCode": "USNCCMP",
	}
	events, err := GetEvents("2023", qparms)
	if err != nil {
		t.Fatalf("TestGetEventsWithEventCode: %s", err.Error())
	}
	t.Log(events)
}

// TestGetEvents calls ftc.GetEvents to test the results of the call
func TestGetEventsWithTeamNumber(t *testing.T) {
	qparms := map[string]string{
		"teamNumber": "7083",
	}
	events, err := GetEvents("2023", qparms)
	if err != nil {
		t.Fatalf("TestGetEventsWithTeamNumber: %s", err.Error())
	}
	t.Log(events)
}
