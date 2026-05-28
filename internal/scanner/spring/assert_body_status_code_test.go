//ff:func feature=scan type=test control=sequence topic=spring
//ff:what assertBodyStatusCode — Java 소스에서 첫 엔드포인트의 statusCode를 검증한다
package spring

import "testing"

func assertBodyStatusCode(t *testing.T, javaSrc string, want string) {
	t.Helper()
	src := []byte(javaSrc)
	root, err := parseJava(src)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	fi := &fileInfo{root: root, src: src, relPath: "C.java", absPath: "/test/C.java"}
	fi.imports = extractImports(root, src)
	controllers := extractControllers(fi)
	if len(controllers) == 0 || len(controllers[0].endpoints) == 0 {
		if want == "" {
			return
		}
		t.Fatal("expected controller with endpoint")
	}
	ep := controllers[0].endpoints[0]
	if ep.statusCode != want {
		t.Errorf("status code: want %q, got %q", want, ep.statusCode)
	}
}
