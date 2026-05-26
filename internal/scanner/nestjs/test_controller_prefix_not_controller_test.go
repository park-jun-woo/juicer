//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestControllerPrefix_NotController 테스트
package nestjs

import "testing"

func TestControllerPrefix_NotController(t *testing.T) {
	src := []byte(`export class SomeService {}`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no classes found")
	}
	_, ok := controllerPrefix(classes[0], src)
	if ok {
		t.Fatal("expected no controller prefix")
	}
}
