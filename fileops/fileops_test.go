package fileops

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
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

func TestFindFilesByPattern(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	testFiles := []string{"file1.txt", "file2.txt", "document.pdf", "image.jpg"}
	for _, filename := range testFiles {
		filePath := filepath.Join(tmpDir, filename)
		err := ioutil.WriteFile(filePath, []byte("Test content"), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Run("Matching Files Found", func(t *testing.T) {
		pattern := filepath.Join(tmpDir, "*.txt")
		files, err := FindFilesByPattern(pattern)

		assert.Nil(t, err)
		assert.Len(t, files, 2)
	})

	t.Run("No Matching Files", func(t *testing.T) {
		pattern := filepath.Join(tmpDir, "nonexistent*.txt")
		files, err := FindFilesByPattern(pattern)

		assert.NotNil(t, err)
		assert.Len(t, files, 0)
		assert.Contains(t, err.Error(), "no files found")
	})
}
