//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestIsTestFile 테스트
package scanner

import "testing"

func TestIsTestFile(t *testing.T) {
	yes := []string{
		"hooks.test.ts", "users.spec.ts", "app.test.js",
		"views_test.py", "handler_spec.rs",
		"UserControllerTest.java", "UserControllerTests.java",
		"UserControllerTest.cs", "UserControllerTests.cs",
	}
	for _, name := range yes {
		if !IsTestFile(name) {
			t.Errorf("IsTestFile(%q) = false, want true", name)
		}
	}
	no := []string{
		"users.controller.ts", "app.ts", "views.py",
		"UserController.java", "UserController.cs", "main.rs",
		"latest.go", "contest.ts",
	}
	for _, name := range no {
		if IsTestFile(name) {
			t.Errorf("IsTestFile(%q) = true, want false", name)
		}
	}
}
