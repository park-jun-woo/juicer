//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractRegisterCall_NotRegister 테스트
package fastify

import "testing"

func TestExtractRegisterCall_NotRegister(t *testing.T) {
	inst := map[string]bool{"app": true}
	if pm := firstRegisterMount(t, `app.get("/x", h);`+"\n", inst); pm != nil {
		t.Fatalf("non-register should be nil, got %+v", pm)
	}
}
