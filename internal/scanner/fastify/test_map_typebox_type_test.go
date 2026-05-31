//ff:func feature=scan type=test topic=fastify control=iteration dimension=1
//ff:what mapTypeBoxType TypeBox 스칼라 → OpenAPI 타입 매핑 테스트
package fastify

import "testing"

func TestMapTypeBoxType(t *testing.T) {
	cases := map[string]string{
		"String": "string", "Number": "number", "Integer": "integer",
		"Boolean": "boolean", "Object": "", "Unknown": "",
	}
	for in, want := range cases {
		if got := mapTypeBoxType(in); got != want {
			t.Errorf("mapTypeBoxType(%q) = %q, want %q", in, got, want)
		}
	}
}
