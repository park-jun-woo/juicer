//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractRegisterCall_PluginNoOpts 테스트
package fastify

import "testing"

func TestExtractRegisterCall_PluginNoOpts(t *testing.T) {

	inst := map[string]bool{"app": true}
	pm := firstRegisterMount(t, `app.register(corsPlugin);`+"\n", inst)
	if pm == nil || pm.PluginRef != "corsPlugin" || pm.Prefix != "" {
		t.Fatalf("expected mount with no prefix, got %+v", pm)
	}
}
