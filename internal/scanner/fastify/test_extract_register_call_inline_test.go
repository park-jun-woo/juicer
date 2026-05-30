//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractRegisterCall_Inline 테스트
package fastify

import "testing"

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
