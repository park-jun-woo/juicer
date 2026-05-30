//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractRegisterCall 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func regCalls(t *testing.T, src string) (*fileInfo, []*sitter.Node) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	return fi, findAllByType(fi.Root, "call_expression")
}

func firstRegisterMount(t *testing.T, src string, inst map[string]bool) *pluginMount {
	t.Helper()
	fi, calls := regCalls(t, src)
	for _, c := range calls {
		if pm := extractRegisterCall(c, fi.Src, inst); pm != nil {
			return pm
		}
	}
	return nil
}

func TestExtractRegisterCall_WithPrefix(t *testing.T) {
	inst := map[string]bool{"app": true}
	pm := firstRegisterMount(t, `app.register(usersPlugin, { prefix: "/users" });`+"\n", inst)
	if pm == nil {
		t.Fatal("expected mount")
	}
	if pm.PluginRef != "usersPlugin" || pm.Prefix != "/users" {
		t.Fatalf("mount = %+v", pm)
	}
}

func TestExtractRegisterCall_Inline(t *testing.T) {
	inst := map[string]bool{"app": true}
	pm := firstRegisterMount(t, "app.register(async (f) => {});\n", inst)
	if pm == nil {
		t.Fatal("expected mount")
	}
	if !pm.Inline || pm.PluginRef != inlineRef {
		t.Fatalf("expected inline mount, got %+v", pm)
	}
}

func TestExtractRegisterCall_NotInstance(t *testing.T) {
	inst := map[string]bool{"app": true}
	if pm := firstRegisterMount(t, `other.register(p);`+"\n", inst); pm != nil {
		t.Fatalf("non-instance should be nil, got %+v", pm)
	}
}

func TestExtractRegisterCall_NotRegister(t *testing.T) {
	inst := map[string]bool{"app": true}
	if pm := firstRegisterMount(t, `app.get("/x", h);`+"\n", inst); pm != nil {
		t.Fatalf("non-register should be nil, got %+v", pm)
	}
}

func TestExtractRegisterCall_NoMember(t *testing.T) {
	inst := map[string]bool{"app": true}
	if pm := firstRegisterMount(t, `register(p);`+"\n", inst); pm != nil {
		t.Fatalf("plain call should be nil, got %+v", pm)
	}
}

func TestExtractRegisterCall_PluginNoOpts(t *testing.T) {
	// register(plugin) with no options object -> mount with empty prefix
	inst := map[string]bool{"app": true}
	pm := firstRegisterMount(t, `app.register(corsPlugin);`+"\n", inst)
	if pm == nil || pm.PluginRef != "corsPlugin" || pm.Prefix != "" {
		t.Fatalf("expected mount with no prefix, got %+v", pm)
	}
}

func TestExtractRegisterCall_NoArgs(t *testing.T) {
	inst := map[string]bool{"app": true}
	if pm := firstRegisterMount(t, `app.register();`+"\n", inst); pm != nil {
		t.Fatalf("no-arg register should be nil, got %+v", pm)
	}
}
