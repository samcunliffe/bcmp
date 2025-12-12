// Test helpers for the bcmp internal subpackages
package bcmptest

import (
	"os"
	"testing"
)

func putFileBack(t *testing.T, path string, empty bool) {
	t.Helper()
	f, err := os.Create(path)
	if err != nil {
		panic("unable to put back file: " + err.Error())
	}
	defer f.Close()

	if !empty {
		if _, err := f.WriteString("Just a non-empty test file."); err != nil {
			panic("unable to write test file: " + err.Error())
		}
	}
}

// Utility test helper to recreate a moved file after testing is done.
func PutFileBack(t *testing.T, path string) {
	t.Helper()
	putFileBack(t, path, false)
}

func PutEmptyFileBack(t *testing.T, path string) {
	t.Helper()
	putFileBack(t, path, true)
}
