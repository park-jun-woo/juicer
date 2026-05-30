//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractFirstParamName_Typed 테스트
package fastify

import "testing"

func TestExtractFirstParamName_Typed(t *testing.T) {

	fi := mustParse(t, []byte("function f(app: FastifyInstance) {}\n"))
	params := findAllByType(fi.Root, "formal_parameters")[0]
	if got := extractFirstParamName(params, fi.Src); got != "app" {
		t.Fatalf("expected app, got %q", got)
	}
}
