package yourpackage

import (
	"fmt"
	"path/filepath"
	"testing"

	testdata "github.com/taylormonacelli/bootedseeds/testdata"
)

func TestYourFunction(t *testing.T) {
	dataPath := filepath.Join(testdata.Dir(), "data.json")

	fmt.Println(dataPath)
	// Use dataPath in your test function
}
