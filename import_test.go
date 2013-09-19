package collada

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

var emptyCollada string = `
<COLLADA version="1.5.0">
</COLLADA>
`

func TestEmptyDocument(t *testing.T) {
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

//A simple scene with a cube
func TestCubeDocument(t *testing.T) {
	compareColladaFile("cube.dae", t)
}

//A more complicated geometry
func TestScrewDocument(t *testing.T) {
	compareColladaFile("screw.dae", t)
}

func compareColladaFile(filename string, t *testing.T) {
	file, err := os.Open(filename)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer file.Close()
	collada, err := LoadDocumentFromReader(file)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	buffer := &bytes.Buffer{}
	err = collada.ExportToWriter(buffer)
	if err != nil {
		t.Error(err)
	}
	file.Seek(0, os.SEEK_SET)
	reader := bytes.NewReader(buffer.Bytes())
	CompareXml(file, reader, t)
}

func failCompare(t *testing.T, message string, a, b interface{}) {
	t.Error(message, "\nExpected:", a, "\nActual :", b)
	t.FailNow()
}

func nextToken(decoder *xml.Decoder) (xml.Token, error) {
	t, err := decoder.Token()
	for err == nil {
		switch t := t.(type) {
		case xml.CharData:
			sa := strings.TrimSpace(string(t))
			if len(sa) > 0 {
				return t, err
			}
		default:
			return t, err
		}
		t, err = decoder.Token()
	}
	return t, err
}

func CompareXml(ar, br io.Reader, t *testing.T) {
	a := xml.NewDecoder(ar)
	b := xml.NewDecoder(br)
	var ea, eb error
	fmt.Println("compare")
	for true {
		var na, nb xml.Token
		na, ea = nextToken(a)
		nb, eb = nextToken(b)
		if ea != nil || eb != nil {
			break
		}
		switch na := na.(type) {
		case xml.StartElement:
			switch nb := nb.(type) {
			case xml.StartElement:
				if na.Name.Local != nb.Name.Local {
					failCompare(t, "wrong node name", na.Name, nb.Name)
				}
				lenAttr := len(na.Attr)
				if lenAttr != len(nb.Attr) {
					failCompare(t, "wrong attribute length <"+na.Name.Local+">", na.Attr, nb.Attr)
					t.FailNow()
				}
				aa := make(map[xml.Name]string, lenAttr)
				ab := make(map[xml.Name]string, lenAttr)
				for i := 0; i < lenAttr; i++ {
					aa[na.Attr[i].Name] = na.Attr[i].Value
					ab[nb.Attr[i].Name] = nb.Attr[i].Value
				}
				for k, v := range aa {
					if v != ab[k] {
						failCompare(t, "wrong attribute <"+na.Name.Local+">"+k.Local, v, ab[k])
					}
				}
				for i := 0; i < lenAttr; i++ {
					if na.Attr[i] != nb.Attr[i] {
					}
				}
			default:
				t.Error("\n", na, "\n", nb)
				failCompare(t, "wrong token", na, reflect.TypeOf(nb))
				t.FailNow()
			}
		case xml.CharData:
			switch nb := nb.(type) {
			case xml.CharData:
				sa := strings.TrimSpace(string(na))
				sb := strings.TrimSpace(string(nb))
				if sa != sb {
					failCompare(t, "wrong text", sa, sb)
				}
			default:
				t.Error("\n", na, "\n", nb)
				failCompare(t, "wrong token", reflect.TypeOf(na), reflect.TypeOf(nb))
			}
		}
	}
	fmt.Println(ea, eb)
	if ea == nil {
		t.Error("missing elements")
	}
	if eb == nil {
		t.Error("extra element")
	}
}
