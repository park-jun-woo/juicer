//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what extractAutoloadCall: register(autoload,{dir,options:{prefix}})에서 dir/prefix 추출을 검증
package fastify

import "testing"

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
