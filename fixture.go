// Package fixture provides test assertions using test fixtures with nice line diffs, and an -update flag for updating fixtures.
package fixture

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/shibukawa/cdiff"
)

// update flag.
var update = flag.Bool("update", false, "Update test fixtures.")

// Read a test fixture from the "testdata" directory.
func Read(t testing.TB, name string) []byte {
	t.Helper()
	path := filepath.Join("testdata", name)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("error reading fixture %q: %s", name, err)
	}
	return b
}

// Write a test fixture to the "testdata" directory.
func Write(t testing.TB, name string, b []byte) {
	t.Helper()
	path := filepath.Join("testdata", name)
	err := ioutil.WriteFile(path, b, 0755)
	if err != nil {
		t.Fatalf("error writing fixture %q: %s", path, err)
	}
}

// Assert that the contents of fixture name matches the expected output.
func Assert(t testing.TB, name string, expected []byte) {
	// update fixtures
	if *update {
		t.Logf("updating test fixture %q", name)
		Write(t, name, expected)
	}

	// read fixture
	actual := Read(t, name)

	// assert
	act := string(actual)
	exp := string(expected)
	if act != exp {
		result := cdiff.Diff(exp, act, cdiff.LineByLine)
		t.Fatalf("Result does not match %q:\n%s", name, result.UnifiedWithGooKitColor("Expected", "Actual", 5, cdiff.GooKitColorTheme))
	}
}
