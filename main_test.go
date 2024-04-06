package main

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/assert"
)

func Test_ValidPascalPrograms(t *testing.T) {
	// Test all files in tests/pascal_programs/*.pas and tests/expected_jasm_files/*.jasm.
	inputFiles, err := filepath.Glob("tests/pascal_programs/*.pas")
	if err != nil {
		t.Errorf("reading inputFiles = %v", err)
	}

	dmp := diffmatchpatch.New()

	for _, inputFile := range inputFiles {
		t.Run(inputFile, func(t *testing.T) {
			inputFile = path.Base(inputFile)
			inputFile = inputFile[:len(inputFile)-4]

			got, _, _, err := genCode("tests/pascal_programs/" + inputFile)
			if err != nil {
				t.Errorf("genCode() error = %v", err)
				return
			}

			expectedOutputFile := "tests/expected_jasm_files/" + inputFile + ".jasm"

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

func Test_InvalidPascalPrograms(t *testing.T) {
	// Test all files in tests/invalid_pascal_programs/*.pas and *.errors.
	invalidInputFiles, err := filepath.Glob("tests/invalid_pascal_programs/*.pas")
	if err != nil {
		t.Errorf("reading invalidInputFiles = %v", err)
	}

	dmp := diffmatchpatch.New()

	for _, invalidInputFile := range invalidInputFiles {
		t.Run(invalidInputFile, func(t *testing.T) {
			invalidInputFile = path.Base(invalidInputFile)
			invalidInputFile = invalidInputFile[:len(invalidInputFile)-4]

			got, lexerErrors, parserErrors, err := genCode("tests/invalid_pascal_programs/" + invalidInputFile)
			assert.NotNil(t, err)
			assert.Empty(t, got)

			expectedOutputFileErrors := "tests/invalid_pascal_programs/" + invalidInputFile + ".errors"

			expectedOutputError, err := os.ReadFile(expectedOutputFileErrors)
			if err != nil {
				t.Errorf("reading expectedOutputFileErrors = %v", err)
				return
			}

			expectedErrors := lexerErrors.String() + parserErrors.String()

			diffs := dmp.DiffMain(string(expectedOutputError), expectedErrors, false)
			if len(diffs) > 1 {
				t.Errorf("diff = %v", dmp.DiffPrettyText(diffs))
				return
			}
		})
	}
}
