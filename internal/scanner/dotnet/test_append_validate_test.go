//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestAppendValidate 테스트
package dotnet

import "testing"

func TestAppendValidate(t *testing.T) {
	if appendValidate("", "x") != "x" || appendValidate("a", "b") != "a,b" {
		t.Fatal("append")
	}
}
