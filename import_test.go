package collada

import (
    "testing"
    "strings"
)

var emptyCollada string = `
<COLLADA version="1.5.0">
</COLLADA>
`

func TestingEmptyDocument(t *testing.T) {
    reader := strings.NewReader(emptyCollada)
    collada, err := LoadDocumentFromReader(reader)
    if err != nil {
        t.Error(err)
        t.FailNow()
    }
    if collada.Version != Version1_5_0 {
        t.Error("wrong version", collada.Version)
    }
}

func TestingCubeDocument(t *testing.T) {
    reader := strings.NewReader("/Users/glen/Documents/blends/cube.dae")
    collada, err := LoadDocumentFromReader(reader)
    if err != nil {
        t.Error(err)
        t.FailNow()
    }
    err = collada.Export("~/out.dae")
    if err != nil {
        t.Error(err)
    }
}