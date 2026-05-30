//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestMatchMapMethod_Round5 테스트
package dotnet

import "testing"

func TestMatchMapMethod_Round5(t *testing.T) {
	fi := csFileInfo(t, `app.MapGet("/x", () => Results.Ok());`)
	inv := findAllByType(fi.root, "invocation_expression")[0]
	ep, ok := matchMapMethod(inv, fi, map[string]string{})
	if !ok || ep.Method != "GET" || ep.Path != "/x" {
		t.Fatalf("got %+v %v", ep, ok)
	}

	fi2 := csFileInfo(t, `Console.WriteLine("x");`)
	inv2 := findAllByType(fi2.root, "invocation_expression")[0]
	if _, ok := matchMapMethod(inv2, fi2, map[string]string{}); ok {
		t.Fatal("non-map invocation should not match")
	}
}
