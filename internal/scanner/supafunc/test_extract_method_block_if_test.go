//ff:func feature=scan type=test topic=supafunc control=sequence
//ff:what extractMethodBlockFromIf req.method 조건 if → 메서드별 블록 매핑 테스트
package supafunc

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestExtractMethodBlockFromIf(t *testing.T) {
	// method-based condition -> registers consequence
	fi := mustParse(t, []byte(`if (req.method === 'POST') { handlePost(); }`))
	ifs := findAllByType(fi.Root, "if_statement")
	if len(ifs) == 0 {
		t.Fatal("no if_statement")
	}
	result := map[string]*sitter.Node{}
	extractMethodBlockFromIf(ifs[0], fi.Src, result)
	if result["POST"] == nil {
		t.Errorf("POST block not registered: %v", result)
	}

	// non-method condition -> nothing
	fi2 := mustParse(t, []byte(`if (x === 1) { y(); }`))
	ifs2 := findAllByType(fi2.Root, "if_statement")
	r2 := map[string]*sitter.Node{}
	extractMethodBlockFromIf(ifs2[0], fi2.Src, r2)
	if len(r2) != 0 {
		t.Errorf("non-method should register nothing: %v", r2)
	}

	// existing entry must not be overwritten
	pre := ifs[0]
	r3 := map[string]*sitter.Node{"POST": pre}
	extractMethodBlockFromIf(ifs[0], fi.Src, r3)
	if r3["POST"] != pre {
		t.Error("existing POST entry must not be overwritten")
	}
}
