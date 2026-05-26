//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestAddImportNames_RelativePath 테스트
package nestjs

import "testing"

func TestAddImportNames_RelativePath(t *testing.T) {
	src := []byte(`import { CreateUserDto } from './dto/create-user.dto';`)
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
	if result["CreateUserDto"] != "./dto/create-user.dto" {
		t.Fatalf("expected ./dto/create-user.dto, got %q", result["CreateUserDto"])
	}
}
