//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestAppendUniquePrefix_Round5 테스트
package express

import "testing"

func TestAppendUniquePrefix_Round5(t *testing.T) {
	m := map[routerKey][]string{}
	k := routerKey{file: "a.ts", varName: "r"}
	if !appendUniquePrefix(m, k, "/x") {
		t.Fatal("first insert should be true")
	}
	if appendUniquePrefix(m, k, "/x") {
		t.Fatal("duplicate insert should be false")
	}
	if !appendUniquePrefix(m, k, "/y") {
		t.Fatal("new value should be true")
	}
	if len(m[k]) != 2 {
		t.Fatalf("expected 2 prefixes, got %v", m[k])
	}
}
