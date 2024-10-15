package main

import (
	"testing"

	"github.com/neuronOS/cmd"
)

func TestGetSystemInfo(t *testing.T) {
	cmdr := cmd.NewCommander()
	info, err := cmdr.GetSystemInfo()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if info.Hostname == "" {
		t.Error("Expected hostname to be non-empty")
	}

	if info.IPAddress == "" {
		t.Error("Expected IP address to be non-empty")
	}
}
