package fileops

import (
	"errors"
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

func TestIsPipedInput(t *testing.T) {
	t.Run("Piped Input", func(t *testing.T) {
		// Redirect stdin to a named pipe for testing
		read, write, _ := os.Pipe()
		origStdin := os.Stdin
		os.Stdin = read
		defer func() {
			os.Stdin = origStdin
			read.Close()
			write.Close()
		}()

		piped, err := IsPipedInput()

		assert.Nil(t, err)
		assert.True(t, piped)
	})

	t.Run("Terminal Input", func(t *testing.T) {
		origStdin := os.Stdin
		defer func() {
			os.Stdin = origStdin
		}()

		terminal, err := IsPipedInput()

		// Assert that there are no errors
		assert.Nil(t, err)

		// Assert that terminal is false
		assert.False(t, terminal)
	})

	t.Run("Error", func(t *testing.T) {
		// Simulate an error by closing the stdin
		origStdin := os.Stdin
		os.Stdin = nil
		defer func() {
			os.Stdin = origStdin
		}()

		_, err := IsPipedInput()

		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, os.ErrInvalid))
	})
}
