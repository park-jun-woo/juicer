//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestMergeImports 테스트
package nestjs

import "testing"

func TestMergeImports(t *testing.T) {
	caller := map[string]string{"A": "./a", "B": "./b"}
	file := map[string]string{"B": "./b2", "C": "./c"}
	result := mergeImports(caller, file)
	if result["A"] != "./a" {
		t.Fatal("caller-only key should be preserved")
	}
	if result["B"] != "./b2" {
		t.Fatal("file import should override caller import")
	}
	if result["C"] != "./c" {
		t.Fatal("file-only key should be included")
	}
}
