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
	// Tests all files in tests/valid_pascal_programs/*.pas with his expected jasm output.
	testPath := "tests/valid_pascal_programs/"
	inputFiles, err := filepath.Glob(testPath + "*.pas")
	if err != nil {
		t.Errorf("reading inputFiles = %v", err)
	}

	dmp := diffmatchpatch.New()

	for _, inputFile := range inputFiles {
		t.Run(inputFile, func(t *testing.T) {
			inputFile = path.Base(inputFile)
			inputFile = inputFile[:len(inputFile)-4]

			got, _, _, err := genCode(testPath + inputFile)
			if err != nil {
				t.Errorf("genCode() error = %v", err)
				return
			}

			expectedOutputFile := testPath + inputFile + ".jasm"

			expectedOutput, err := os.ReadFile(expectedOutputFile)
			if err != nil {
				t.Errorf("reading expectedOutputFile = %v", err)
				return
			}

			diffs := dmp.DiffMain(string(expectedOutput), got, false)
			if len(diffs) > 1 {
				t.Errorf("diff = \n%v", dmp.DiffPrettyText(diffs))
			}
		})
	}
}

func Test_InvalidPascalPrograms(t *testing.T) {
	// Tests all files in tests/invalid_pascal_programs/*.pas with his expected output errors.
	testPath := "tests/invalid_pascal_programs/"
	invalidInputFiles, err := filepath.Glob(testPath + "*.pas")
	if err != nil {
		t.Errorf("reading invalidInputFiles = %v", err)
	}

	dmp := diffmatchpatch.New()

	for _, invalidInputFile := range invalidInputFiles {
		t.Run(invalidInputFile, func(t *testing.T) {
			invalidInputFile = path.Base(invalidInputFile)
			invalidInputFile = invalidInputFile[:len(invalidInputFile)-4]

			got, lexerErrors, parserErrors, err := genCode(testPath + invalidInputFile)
			assert.NotNil(t, err)
			assert.Empty(t, got)

			expectedOutputFileErrors := testPath + invalidInputFile + ".errors"

			expectedOutputError, err := os.ReadFile(expectedOutputFileErrors)
			if err != nil {
				t.Errorf("reading expectedOutputFileErrors = %v", err)
				return
			}

			expectedErrors := lexerErrors.String() + parserErrors.String()

			diffs := dmp.DiffMain(string(expectedOutputError), expectedErrors, false)
			if len(diffs) > 1 {
				t.Errorf("diff = \n%v", dmp.DiffPrettyText(diffs))
				return
			}
		})
	}
}
