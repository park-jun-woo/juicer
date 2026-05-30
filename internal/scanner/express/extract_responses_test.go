//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractResponses: 본문없음 기본값 / 다중응답 dedup / 응답없는 본문 기본값
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstArrow(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	a := findAllByType(fi.Root, "arrow_function")
	if len(a) == 0 {
		t.Fatal("no arrow_function")
	}
	return a[0]
}

func TestExtractResponses_NoBodyDefault(t *testing.T) {
	fi := mustParse(t, []byte(`const x=1;`))
	resps := extractResponses(fi, routeInfo{Handler: "(anonymous)"})
	if len(resps) != 1 || resps[0].Status != "200" || resps[0].Kind != "json" {
		t.Fatalf("got %+v", resps)
	}
}

func TestExtractResponses_MultipleDedup(t *testing.T) {
	src := []byte(`const h = (req, res) => { res.json({}); res.json({}); res.status(404).send('x'); };`)
	fi := mustParse(t, src)
	ri := routeInfo{HandlerNode: firstArrow(t, fi)}
	resps := extractResponses(fi, ri)
	// 200/json (deduped) + 404/text => 2 distinct
	if len(resps) != 2 {
		t.Fatalf("expected 2 distinct responses, got %+v", resps)
	}
}

func TestExtractResponses_BodyNoResponsesDefault(t *testing.T) {
	src := []byte(`const h = (req, res) => { doStuff(); };`)
	fi := mustParse(t, src)
	ri := routeInfo{HandlerNode: firstArrow(t, fi)}
	resps := extractResponses(fi, ri)
	if len(resps) != 1 || resps[0].Status != "200" || resps[0].Kind != "json" {
		t.Fatalf("got %+v", resps)
	}
}
