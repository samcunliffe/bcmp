package bcmptest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the test helpers because testing all the way down 🐢🐢🐢
func TestPutFileBack(t *testing.T) {
	path := filepath.Join(t.TempDir(), "testfile.txt")
	assert.NoFileExists(t, path, "Test setup error: test file %q already exists", path)

	PutFileBack(t, path)

	assert.FileExists(t, path, "PutFileBack did not create file %q", path)
}

func TestPutFileBackExistingFile(t *testing.T) {
	path := "testdata/existingfile"
	assert.FileExists(t, path, "Test setup error: test file %q does not exist", path)

	PutFileBack(t, path)

	assert.FileExists(t, path, "PutFileBack should not effect an existing file", path)
}

func TestPutFileBackNoPermissions(t *testing.T) {
	tempdir := t.TempDir()
	err := os.Chmod(tempdir, 0000)
	assert.NoError(t, err, "Test setup error: failed to change permissions for %q: %v", tempdir, err)
	defer os.Chmod(tempdir, 0644)

	path := filepath.Join(tempdir, "no_permission_file.txt")
	assert.Panics(t, func() { PutFileBack(t, path) })
}
