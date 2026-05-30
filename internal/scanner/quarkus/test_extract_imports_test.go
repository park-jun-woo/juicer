//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractImports 테스트
package quarkus

import "testing"

func TestExtractImports(t *testing.T) {
	root, src := parseQ(t, `import com.example.UserDto;
import java.util.List;
class R {}`)
	imports := extractImports(root, src)
	if imports["UserDto"] != "com.example.UserDto" {
		t.Fatalf("got %v", imports)
	}
}
