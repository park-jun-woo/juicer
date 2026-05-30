//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractPluginRef_CallExpression 테스트
package fastify

import "testing"

func TestExtractPluginRef_CallExpression(t *testing.T) {
	n, src := firstNodeOfType(t, `const x = require("@fastify/cors");`+"\n", "call_expression")
	if got := extractPluginRef(n, src); got != "@fastify/cors" {
		t.Fatalf("call: got %q", got)
	}
}
