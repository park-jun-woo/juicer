//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what extractAutoloadCall: register(autoload,{dir,options:{prefix}})에서 dir/prefix 추출을 검증
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestExtractAutoloadCall(t *testing.T) {
	src := []byte(`
const app = Fastify();
app.register(autoload, { dir: join(__dirname, "routes"), options: { prefix: "/api" } });
`)
	fi := mustParse(t, src)
	names := map[string]bool{"autoload": true}
	for _, call := range findAllByType(fi.Root, "call_expression") {
		dir, prefix, ok := extractAutoloadCall(call, fi.Src, names)
		if !ok {
			continue
		}
		if dir != "routes" {
			t.Errorf("dir: want routes, got %q", dir)
		}
		if prefix != "/api" {
			t.Errorf("prefix: want /api, got %q", prefix)
		}
		return
	}
	t.Fatal("autoload call not detected")
}

func acFirstCall(t *testing.T, src string) (*fileInfo, []*sitter.Node) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	return fi, findAllByType(fi.Root, "call_expression")
}

func TestExtractAutoloadCall_NoMemberExpression(t *testing.T) {
	// plain function call, no member_expression -> false
	fi, calls := acFirstCall(t, "foo(autoload, {});\n")
	names := map[string]bool{"autoload": true}
	for _, c := range calls {
		if _, _, ok := extractAutoloadCall(c, fi.Src, names); ok {
			t.Fatal("plain call should not be autoload")
		}
	}
}

func TestExtractAutoloadCall_NotRegister(t *testing.T) {
	// member call but property is not "register"
	fi, calls := acFirstCall(t, "app.listen(3000);\n")
	names := map[string]bool{"autoload": true}
	for _, c := range calls {
		if _, _, ok := extractAutoloadCall(c, fi.Src, names); ok {
			t.Fatal("listen should not be autoload")
		}
	}
}

func TestExtractAutoloadCall_NotAutoloadFirstArg(t *testing.T) {
	// register with a non-autoload plugin as first arg
	fi, calls := acFirstCall(t, `app.register(somePlugin, { dir: join(d, "r") });`+"\n")
	names := map[string]bool{"autoload": true}
	for _, c := range calls {
		if _, _, ok := extractAutoloadCall(c, fi.Src, names); ok {
			t.Fatal("non-autoload plugin should not match")
		}
	}
}

func TestExtractAutoloadCall_NoDir(t *testing.T) {
	// register(autoload, {...}) but no dir -> false
	fi, calls := acFirstCall(t, `app.register(autoload, { options: { prefix: "/api" } });`+"\n")
	names := map[string]bool{"autoload": true}
	for _, c := range calls {
		if _, _, ok := extractAutoloadCall(c, fi.Src, names); ok {
			t.Fatal("missing dir should yield false")
		}
	}
}

func TestExtractAutoloadCall_TooFewArgs(t *testing.T) {
	fi, calls := acFirstCall(t, "app.register(autoload);\n")
	names := map[string]bool{"autoload": true}
	for _, c := range calls {
		if _, _, ok := extractAutoloadCall(c, fi.Src, names); ok {
			t.Fatal("single-arg register should yield false")
		}
	}
}
