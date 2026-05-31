//ff:func feature=ddl type=test control=sequence
//ff:what sortedTableNames 맵 키 알파벳 정렬 반환 테스트
package ddl

import (
	"reflect"
	"testing"
)

func TestSortedTableNames(t *testing.T) {
	tables := map[string]*Table{"users": {}, "accounts": {}, "orgs": {}}
	got := sortedTableNames(tables)
	want := []string{"accounts", "orgs", "users"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	if got := sortedTableNames(map[string]*Table{}); len(got) != 0 {
		t.Errorf("empty: %v", got)
	}
}
