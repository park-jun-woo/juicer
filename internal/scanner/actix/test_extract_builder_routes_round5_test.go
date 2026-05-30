//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractBuilderRoutes_Round5 테스트
package actix

import "testing"

func TestExtractBuilderRoutes_Round5(t *testing.T) {
	fi := aFi(t, builderSrc)
	routes := extractBuilderRoutes(fi)
	if len(routes) == 0 {
		t.Fatalf("expected builder routes, got %d", len(routes))
	}
}
