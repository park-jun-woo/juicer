//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractRegisterCall_NotInstance 테스트
package fastify

import "testing"

func TestExtractRegisterCall_NotInstance(t *testing.T) {
	inst := map[string]bool{"app": true}
	if pm := firstRegisterMount(t, `other.register(p);`+"\n", inst); pm != nil {
		t.Fatalf("non-instance should be nil, got %+v", pm)
	}
}
