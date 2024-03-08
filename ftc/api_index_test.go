package ftc

import (
	"testing"
)

// TestGetApiIndex calls ftc.GetApiIndex to test the results of the call
func TestGetApiIndex(t *testing.T) {
	apis, err := GetApiIndex()
	if err != nil {
		t.Fatalf("TestGetApiIndex: %s", err.Error())
	}
	t.Log(apis)
}
