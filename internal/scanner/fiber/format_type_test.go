//ff:func feature=scan type=test control=selection
//ff:what formatType — 타입 문자열화 테스트
package fiber

import "testing"

func TestFormatType(t *testing.T) {
	src := `package m
import "time"
type Inner struct{ X int }
type T struct {
	B  int
	P  *int
	Sl []string
	Ar [3]byte
	Mp map[string]int
	Nm time.Time
	If interface{}
	St Inner
	Lo error
	An struct{ Y int }
	Ch chan int
}
`
	st, _ := structFields(t, src, "T")
	want := map[string]string{
		"B":  "int",
		"P":  "*int",
		"Sl": "[]string",
		"Ar": "[]byte",
		"Mp": "map[string]int",
		"Nm": "time.Time",
		"If": "any",
		"St": "m.Inner", // named type with package
		"Lo": "error",   // named with nil pkg (universe scope)
		"An": "object",  // anonymous struct
		"Ch": "chan int", // default: t.String()
	}
	for i := 0; i < st.NumFields(); i++ {
		f := st.Field(i)
		got := formatType(f.Type())
		if w, ok := want[f.Name()]; ok && got != w {
			t.Errorf("formatType(%s) = %q, want %q", f.Name(), got, w)
		}
	}
}
