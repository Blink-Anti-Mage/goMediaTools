package mkdir

import (
	"os"
	"testing"
)

func TestMkdir(t *testing.T) {
	dirName := "testDir"
	// // Clean up the directory after the test
	// defer func() {
	// 	err := os.Remove(dirName)
	// 	if err != nil {
	// 		t.Errorf("Failed to remove directory: %v", err)
	// 	}
	// }()
	Mkdir(dirName)

	// Check if the directory was created
	_, err := os.Stat(dirName)
	if err != nil {
		t.Errorf("Failed to create directory: %v", err)
	}
}

func TestRedir(t *testing.T) {
	dirName := "testDir"

	// Create the directory to be removed
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	// Call the function being tested
	Redir(dirName)

	// Check if the directory was removed
	_, err = os.Stat(dirName)
	if !os.IsNotExist(err) {
		t.Errorf("Directory was not removed: %v", err)
	}
}
