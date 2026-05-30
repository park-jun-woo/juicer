//ff:func feature=scan type=test control=sequence
//ff:what TestResolveType_Variants 테스트
package fiber

import "testing"

func TestResolveType_Variants(t *testing.T) {
	src := `package m
import "time"
type Item struct { ID int ` + "`json:\"id\"`" + ` }
type Holder struct {
	Ptr   *Item
	List  []Item
	PList []*Item
	Arr   [2]Item
	When  time.Time
	Whens []time.Time
	PArr  [2]*Item
	Num   int
}
`
	st, _ := structFields(t, src, "Holder")
	get := func(i int) (string, int) {
		tn, f := resolveType(st.Field(i).Type())
		return tn, len(f)
	}

	if tn, n := get(0); tn != "Item" || n != 1 {
		t.Errorf("Ptr: %q %d", tn, n)
	}
	if tn, n := get(1); tn != "[]Item" || n != 1 {
		t.Errorf("List: %q %d", tn, n)
	}
	if tn, n := get(2); tn != "[]Item" || n != 1 {
		t.Errorf("PList: %q %d", tn, n)
	}
	if tn, n := get(3); tn != "[]Item" || n != 1 {
		t.Errorf("Arr: %q %d", tn, n)
	}
	if tn, _ := get(4); tn != "time.Time" {
		t.Errorf("When: %q", tn)
	}
	if tn, _ := get(5); tn != "[]time.Time" {
		t.Errorf("Whens: %q", tn)
	}
	if tn, n := get(6); tn != "[]Item" || n != 1 {
		t.Errorf("PArr: %q %d", tn, n)
	}
	if tn, n := get(7); tn != "" || n != 0 {
		t.Errorf("Num: %q %d", tn, n)
	}
}
