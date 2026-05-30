//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractRegisterCall_NoArgs 테스트
package fastify

import "testing"

func TestExtractRegisterCall_NoArgs(t *testing.T) {
	inst := map[string]bool{"app": true}
	if pm := firstRegisterMount(t, `app.register();`+"\n", inst); pm != nil {
		t.Fatalf("no-arg register should be nil, got %+v", pm)
	}
}
