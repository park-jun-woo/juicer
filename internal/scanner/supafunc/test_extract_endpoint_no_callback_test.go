//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractEndpoint_NoCallback 테스트
package supafunc

import "testing"

func TestExtractEndpoint_NoCallback(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "hello/index.ts", `const x = 1;`)
	fi, err := parseFile(dir + "/hello/index.ts")
	if err != nil {
		t.Fatal(err)
	}
	eps := extractEndpoint(fi, dir)
	if len(eps) != 1 || eps[0].Method != "POST" || eps[0].Path != "/hello" {
		t.Fatalf("got %+v", eps)
	}
}
