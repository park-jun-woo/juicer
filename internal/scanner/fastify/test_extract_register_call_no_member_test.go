//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractRegisterCall_NoMember 테스트
package fastify

import "testing"

func TestExtractRegisterCall_NoMember(t *testing.T) {
	inst := map[string]bool{"app": true}
	if pm := firstRegisterMount(t, `register(p);`+"\n", inst); pm != nil {
		t.Fatalf("plain call should be nil, got %+v", pm)
	}
}
