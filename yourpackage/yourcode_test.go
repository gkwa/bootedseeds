package yourpackage_test

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestYourFunction(t *testing.T) {
	dataPath := filepath.Join("testdata", "data.json")

	fmt.Println(dataPath)
	// Use dataPath in your test function
}
