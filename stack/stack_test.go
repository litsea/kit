package stack

import (
	"testing"
)

func nested(level int) *Frame {
	return SkipLevel(level)
}

func TestStack(t *testing.T) {
	tests := []struct {
		name     string
		frame    *Frame
		file     string
		line     int
		function string
	}{
		{
			name:     "stack",
			frame:    Stack(),
			file:     "stack_test.go",
			line:     21,
			function: "TestStack",
		},
		{
			name:     "stack-level1",
			frame:    nested(1),
			file:     "stack_test.go",
			line:     28,
			function: "TestStack",
		},
		{
			name:     "stack-level0",
			frame:    nested(0),
			file:     "stack_test.go",
			line:     8,
			function: "nested",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.frame.FileName() != tt.file {
				t.Errorf("TestStack FileName() = %s, want %s", tt.frame.FileName(), tt.file)
			}
			if tt.frame.Line != tt.line {
				t.Errorf("TestStack Line = %d, want %d", tt.frame.Line, tt.line)
			}
			if tt.frame.FuncName() != tt.function {
				t.Errorf("TestStack FuncName() = %s, want %s", tt.frame.FuncName(), tt.function)
			}
		})
	}
}
