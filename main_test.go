package main

import (
	"os"
	"testing"
)

func Test_genCode(t *testing.T) {
	// Files in examples/*.pas and examples/expected/*.jasm.
	tests := []struct {
		inputFile   string
		expectedErr error
	}{
		{
			inputFile:   "hello_world",
			expectedErr: nil,
		},
		{
			inputFile:   "concat_two_strings",
			expectedErr: nil,
		},
		{
			inputFile:   "concat_three_strings",
			expectedErr: nil,
		},
		{
			inputFile:   "add_two_numbers",
			expectedErr: nil,
		},
		{
			inputFile:   "add_three_numbers",
			expectedErr: nil,
		},
		{
			inputFile:   "sub_two_numbers",
			expectedErr: nil,
		},
		{
			inputFile:   "sub_three_numbers",
			expectedErr: nil,
		},
		// // examples/fatorial.pas
		// // examples/name_and_age.pas
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			got, err := genCode("examples/" + tt.inputFile)
			if err != tt.expectedErr {
				t.Errorf("genCode() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}

			expectedOutputFile := "examples/expected/" + tt.inputFile + ".jasm"

			expectedOutput, err := os.ReadFile(expectedOutputFile)
			if err != nil {
				t.Errorf("reading expectedOutputFile = %v", err)
				return
			}

			if got != string(expectedOutput) {
				t.Errorf("genCode() = %v, want %v", got, string(expectedOutput))
			}
		})
	}
}
