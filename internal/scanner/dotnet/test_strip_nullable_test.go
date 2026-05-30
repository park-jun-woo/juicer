//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestStripNullable 테스트
package dotnet

import "testing"

func TestStripNullable(t *testing.T) {
	n, ok := stripNullable("int?")
	if n != "int" || !ok {
		t.Fatalf("nullable: %q %v", n, ok)
	}
	n2, ok2 := stripNullable("int")
	if n2 != "int" || ok2 {
		t.Fatalf("non-nullable: %q %v", n2, ok2)
	}
}
