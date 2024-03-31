package main

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func Test_genCode(t *testing.T) {
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

			t.Logf("Running expected output for %s", inputFile)

			got, err := genCode("tests/pascal_programs/" + inputFile)
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

			t.Logf("Expected output for %s is fine", inputFile)
		})
	}
}
