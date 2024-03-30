package main

import (
	"os"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func Test_genCode(t *testing.T) {
	// Files in tests/pascal_programs/*.pas and tests/expected_jasm_files/*.jasm.
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
		{
			inputFile:   "mul_three_numbers",
			expectedErr: nil,
		},
		{
			inputFile:   "div_three_numbers",
			expectedErr: nil,
		},
		{
			inputFile:   "operator_precedence",
			expectedErr: nil,
		},
		{
			inputFile:   "hello_world_two_types",
			expectedErr: nil,
		},
		{
			inputFile:   "if_with_integers_without_and_or",
			expectedErr: nil,
		},
		{
			inputFile:   "if_with_integers_with_and_or",
			expectedErr: nil,
		},
		{
			inputFile:   "if_with_integers_with_not",
			expectedErr: nil,
		},
		// // examples/fatorial.pas
		// // examples/name_and_age.pas
	}

	dmp := diffmatchpatch.New()

	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			got, err := genCode("tests/pascal_programs/" + tt.inputFile)
			if err != tt.expectedErr {
				t.Errorf("genCode() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}

			expectedOutputFile := "tests/expected_jasm_files/" + tt.inputFile + ".jasm"

			expectedOutput, err := os.ReadFile(expectedOutputFile)
			if err != nil {
				t.Errorf("reading expectedOutputFile = %v", err)
				return
			}

			diffs := dmp.DiffMain(string(expectedOutput), got, false)
			if len(diffs) > 1 {
				t.Errorf("diff = %v", dmp.DiffPrettyText(diffs))
			}
		})
	}
}
