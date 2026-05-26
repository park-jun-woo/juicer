//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestAddImportNames_NonRelativeSrc 테스트
package nestjs

import "testing"

func TestAddImportNames_NonRelativeSrc(t *testing.T) {
	src := []byte(`import { CreateUserDto } from 'src/users/dto/create-user.dto';`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	result := make(map[string]string)
	stmt := findChildByType(root, "import_statement")
	if stmt == nil {
		t.Fatal("expected import_statement")
	}
	addImportNames(stmt, src, result)
	if result["CreateUserDto"] != "src/users/dto/create-user.dto" {
		t.Fatalf("expected src/users/dto/create-user.dto, got %q", result["CreateUserDto"])
	}
}
