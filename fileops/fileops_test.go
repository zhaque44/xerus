package fileops

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {
	// Create a temporary test file
	content := "Line 1\nLine 2\nLine 3"
	tmpFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // Cleanup

	_, err = tmpFile.WriteString(content)
	if err != nil {
		tmpFile.Close()
		t.Fatal(err)
	}
	tmpFile.Close()

	// Test the LoadFile function
	lines, err := LoadFile(tmpFile.Name())

	// Assert that there are no errors
	assert.Nil(t, err)

	// Assert the expected output
	expectedLines := []string{"Line 1", "Line 2", "Line 3"}
	assert.Equal(t, expectedLines, lines)
}
