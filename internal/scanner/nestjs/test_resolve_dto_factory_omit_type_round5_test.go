//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestResolveDTOFactory_OmitType_Round5 테스트
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveDTOFactory_OmitType_Round5(t *testing.T) {
	b := []byte("class C extends OmitType(BaseDto, ['name']) {}")
	root, _ := parseTypeScript(b)
	calls := findAllByType(root, "call_expression")
	cache := map[string][]scanner.Field{
		"BaseDto": {
			{Name: "id", Type: "number"},
			{Name: "name", Type: "string"},
		},
	}
	got := resolveDTOFactory(calls[0], b, "/proj/c.dto.ts", map[string]string{}, "/proj", cache)
	for _, f := range got {
		if f.name == "name" {
			t.Errorf("name should be omitted: %+v", got)
		}
	}
}
