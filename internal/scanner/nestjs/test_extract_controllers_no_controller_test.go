//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractControllers_NoController 테스트
package nestjs

import "testing"

func TestExtractControllers_NoController(t *testing.T) {
	src := []byte(`export class SomeService {}`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	controllers := extractControllers(root, src, "test.ts")
	if len(controllers) != 0 {
		t.Fatalf("expected 0, got %d", len(controllers))
	}
}
