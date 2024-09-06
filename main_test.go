// program_test.go
package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

// Function to run the bash script
func runBashScript() (string, error) {
	cmd := exec.Command("./bin/math-skills")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

// Function to run the Go program
func runGoProgram() (string, error) {
	cmd := exec.Command("./mathskils", "data.txt")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

// The test function
func TestProgramOutput(t *testing.T) {
	// Run the Bash script
	bashOutput, err := runBashScript()
	if err != nil {
		t.Fatalf("Failed to run bash script: %v", err)
	}

	// Run the Go program
	goOutput, err := runGoProgram()
	if err != nil {
		t.Fatalf("Failed to run Go program: %v", err)
	}

	// Compare the outputs

	if bashOutput != goOutput {
		t.Errorf("Outputs don't match.\nBash output:\n%s\nGo output:\n%s", bashOutput, goOutput)
	} else {
		t.Logf("Success: Outputs match!\nBash output:\n%s\nGo output:\n%s", bashOutput, goOutput)
	}
}
