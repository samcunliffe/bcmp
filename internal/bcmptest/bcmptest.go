// Test helpers for the bcmp internal subpackages
package bcmptest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Utility test helper to recreate a moved file after testing is done.
// Usually deferred.
func PutFileBack(t *testing.T, path string) {
	t.Helper()
	f, err := os.Create(path)
	if err != nil {
		panic("unable to put back file: " + err.Error())
	}
	defer f.Close()

	f.WriteString("Just a non-empty test file.")
}

// Walks through the directory and counts files and directories.
func DirCount(path string) (int, error) {
	count := -1 // exclude the root directory
	err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		count++
		return nil
	})
	return count, err
}

// Convenience function when testing: the directory at 'path' must exist and be empty.
func AssertDirEmpty(t *testing.T, path string, msgAndArgs ...interface{}) {
	t.Helper()
	assert.DirExists(t, path)
	count, err := DirCount(path)
	assert.NoError(t, err, "AssertDirEmpty failed to read directory %q: %v", path, err)
	assert.Zero(t, count, msgAndArgs...)
}
