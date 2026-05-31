//ff:func feature=scan type=test topic=fastify control=sequence
//ff:what extractRegisterCall app.register(plugin, {prefix}) 마운트 추출 및 비대상 무시 테스트
package fastify

import "testing"

func TestExtractRegisterCall(t *testing.T) {
	instances := map[string]bool{"app": true}

	// app.register(userRoutes, { prefix: '/users' })
	fi := mustParse(t, []byte(`app.register(userRoutes, { prefix: '/users' })`))
	call := findAllByType(fi.Root, "call_expression")[0]
	pm := extractRegisterCall(call, fi.Src, instances)
	if pm == nil || pm.PluginRef != "userRoutes" || pm.Prefix != "/users" {
		t.Fatalf("got %+v", pm)
	}

	// inline plugin
	fi2 := mustParse(t, []byte(`app.register(async (instance) => {})`))
	c2 := findAllByType(fi2.Root, "call_expression")[0]
	pm2 := extractRegisterCall(c2, fi2.Src, instances)
	if pm2 == nil || !pm2.Inline {
		t.Errorf("inline: %+v", pm2)
	}

	// not a register call
	fi3 := mustParse(t, []byte(`app.get('/x', h)`))
	c3 := findAllByType(fi3.Root, "call_expression")[0]
	if extractRegisterCall(c3, fi3.Src, instances) != nil {
		t.Error("non-register should be nil")
	}

	// unknown instance
	fi4 := mustParse(t, []byte(`other.register(p)`))
	c4 := findAllByType(fi4.Root, "call_expression")[0]
	if extractRegisterCall(c4, fi4.Src, instances) != nil {
		t.Error("unknown instance should be nil")
	}
}
