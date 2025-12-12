package bcmptest

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the test helpers because testing all the way down 🐢🐢🐢
func TestPutFileBack(t *testing.T) {
	path := path.Join(t.TempDir(), "testfile.txt")
	assert.NoFileExists(t, path, "Test setup error: test file %q already exists", path)

	PutFileBack(t, path)

	assert.FileExists(t, path, "PutFileBack did not create file %q", path)
}
