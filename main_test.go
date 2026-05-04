package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestArguments(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		expectError    bool
	}{
		{
			name:           "Invalid flag spelling",
			args:           []string{"--alin=center", "hello"},
			expectedOutput: "Usage: go run . [OPTION] [STRING] [BANNER]",
			expectError:    true,
		},
		{
			name:           "Unsupported align type",
			args:           []string{"--align=top", "hello"},
			expectedOutput: "Usage: go run . [OPTION] [STRING] [BANNER]",
			expectError:    true,
		},
		{
			name:           "Missing equal sign",
			args:           []string{"--align center", "hello"},
			expectedOutput: "Usage: go run . [OPTION] [STRING] [BANNER]",
			expectError:    true,
		},
		{
			name:           "Empty align value",
			args:           []string{"--align=", "hello"},
			expectedOutput: "Usage: go run . [OPTION] [STRING] [BANNER]",
			expectError:    true,
		},
		{
			name:           "Unsupported align type",
			args:           []string{"--align=top", "hello"},
			expectedOutput: "Usage: go run . [OPTION] [STRING] [BANNER]",
			expectError:    true,
		},
		{
			name:           "Missing input string",
			args:           []string{"--align=left"},
			expectedOutput: "Usage: go run . [OPTION] [STRING] [BANNER]",
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run the actual main.go file
			// We use 'go run .' to execute the current directory
			cmd := exec.Command("go", append([]string{"run", "."}, tt.args...)...)
			output, _ := cmd.CombinedOutput()
			outStr := string(output)

			if tt.expectError {
				if !strings.Contains(outStr, tt.expectedOutput) {
					t.Errorf("Expected error containing '%s', got: %s", tt.expectedOutput, outStr)
				}
			} else {
				if strings.Contains(outStr, "Error:") || strings.Contains(outStr, "Usage:") {
					t.Errorf("Expected successful run, but got: %s", outStr)
				}
			}
		})
	}
}

func TestFunctionality(t *testing.T) {
	// Testing the Calculator logic directly
	t.Run("Calculator logic", func(t *testing.T) {
		got := Calculator('A')
		// 'A' is ASCII 65. (65-32)*9 + 1 = 33*9 + 1 = 297 + 1 = 298
		want := 298
		if got != want {
			t.Errorf("Calculator('A') = %d; want %d", got, want)
		}
	})
}
