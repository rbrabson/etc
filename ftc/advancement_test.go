package ftc

import (
	"testing"
)

// TestGetAdvacementFrom calls ftc.GetAdvancementFrom to test the results of the call
func TestGetAdvacementFrom(t *testing.T) {
	advancements, err := GetAdvancementsFrom("2023", "USNCCMP")
	if err != nil {
		t.Fatalf("TestGetAdvacementFrom: %s", err.Error())
	}
	t.Log(advancements)
}

// TestGetAdvacementTo calls ftc.GetAdvancementTo to test the results of the call
func TestGetAdvacementTo(t *testing.T) {
	advancements, err := GetAdvancementsTo("2023", "USNCCMP")
	if err != nil {
		t.Fatalf("TestGetAdvacementTo: %s", err.Error())
	}
	t.Log(advancements)
}
