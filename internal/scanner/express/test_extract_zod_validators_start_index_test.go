//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractZodValidatorsFromArgs 시작 인덱스 테스트: start=0이면 arg0 검증자 추출, start=1이면 arg0 스킵
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestExtractZodValidatorsFromArgs_StartIndex(t *testing.T) {
	src := []byte(`router.route('/items').post(validateRequest({ body: z.object({ name: z.string() }) }), createItem);`)
	fi := mustParse(t, src)

	var postArgs *sitter.Node
	walkNodes(fi.Root, func(n *sitter.Node) {
		if postArgs != nil || n.Type() != "arguments" {
			return
		}
		argNodes := collectArgNodes(n)
		if len(argNodes) >= 1 && argNodes[0].Type() == "call_expression" {
			postArgs = n
		}
	})
	if postArgs == nil {
		t.Fatal("could not locate .post() arguments node with leading call_expression")
	}

	argNodes := collectArgNodes(postArgs)

	if got := extractZodValidatorsFromArgs(argNodes, src, 0); len(got) != 1 {
		t.Fatalf("start=0: expected 1 validator (arg0), got %d", len(got))
	}
	if got := extractZodValidatorsFromArgs(argNodes, src, 1); len(got) != 0 {
		t.Fatalf("start=1: expected 0 validators (arg0 skipped), got %d", len(got))
	}
}
