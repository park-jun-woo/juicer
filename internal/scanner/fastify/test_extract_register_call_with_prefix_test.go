//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractRegisterCall_WithPrefix 테스트
package fastify

import "testing"

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
