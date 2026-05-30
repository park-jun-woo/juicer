//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractOneResponse: json/send/sendStatus/status체인/비res 분기 검증
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstCallExpr(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no call_expression")
	}
	return calls[0]
}

func TestExtractOneResponse_JsonDefault(t *testing.T) {
	fi := mustParse(t, []byte(`res.json({a:1});`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r == nil || r.Status != "200" || r.Kind != "json" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneResponse_JsonWithStatus(t *testing.T) {
	fi := mustParse(t, []byte(`res.status(201).json({a:1});`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r == nil || r.Status != "201" || r.Kind != "json" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneResponse_SendDefault(t *testing.T) {
	fi := mustParse(t, []byte(`res.send('ok');`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r == nil || r.Status != "200" || r.Kind != "text" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneResponse_SendStatus(t *testing.T) {
	fi := mustParse(t, []byte(`res.sendStatus(204);`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r == nil || r.Status != "204" || r.Kind != "empty" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneResponse_SendStatusNoArg(t *testing.T) {
	// sendStatus with no numeric arg -> extractSendStatusArg returns "" -> defaults 200
	fi := mustParse(t, []byte(`res.sendStatus();`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r == nil || r.Status != "200" || r.Kind != "empty" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneResponse_NotResMethod(t *testing.T) {
	fi := mustParse(t, []byte(`foo.bar(1);`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneResponse_OtherResMethod(t *testing.T) {
	// res method that is not json/send/sendStatus -> default return nil
	fi := mustParse(t, []byte(`res.render('view');`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r != nil {
		t.Fatalf("expected nil for render, got %+v", r)
	}
}
