package ftc

import "testing"

// TestGetHybridSchedule calls ftc.GetHybridSchedule to test the results of the call
func TestGetHybridSchedule(t *testing.T) {
	results, err := GetHybridSchedule("2023", "USNCCMP", PLAYOFF)
	if err != nil {
		t.Fatalf("GetHybridSchedule: %s", err.Error())
	}
	t.Log(results)
}

// TestGetEventSchedule calls ftc.GetEventSchedule to test the results of the call
func TestGetEventSchedule(t *testing.T) {
	results, err := GetEventSchedule("2023", "USNCCMP", PLAYOFF)
	if err != nil {
		t.Fatalf("TestGetEventSchedule: %s", err.Error())
	}
	t.Log(results)
}

// TestGetEventScheduleWithTeamNumber calls ftc.GetEventSchedule to test the results of the call
func TestGetEventScheduleWithTeamNumber(t *testing.T) {
	results, err := GetEventSchedule("2023", "USNCCMP", PLAYOFF, "7083")
	if err != nil {
		t.Fatalf("TestGetEventScheduleWithTeamNumber: %s", err.Error())
	}
	t.Log(results)
}
