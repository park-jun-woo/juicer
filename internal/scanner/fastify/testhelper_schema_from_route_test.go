//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what schemaFromRoute 테스트 헬퍼
package fastify

import "testing"

func schemaFromRoute(t *testing.T, routeSrc string) (*schemaInfo, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(routeSrc))
	instances := collectInstances(fi)
	routes := extractRoutes(fi, instances)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	si := extractJSONSchema(routes[0].Schema, fi.Src)
	if si == nil {
		t.Fatal("nil schemaInfo")
	}
	return si, fi.Src
}
