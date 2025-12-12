// Test helpers for the bcmp internal subpackages
package bcmptest

import (
	"os"
	"testing"
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
