//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveParamName 테스트
package nestjs

import "testing"

func TestResolveParamName(t *testing.T) {

	if got := resolveParamName(DecParam, "userId", "id", "/users/:id"); got != "userId" {
		t.Fatalf("got %q", got)
	}

	if got := resolveParamName(DecParam, "", "id", "/users/:id"); got != "id" {
		t.Fatalf("got %q", got)
	}

	if got := resolveParamName(DecParam, "", "p", "/a/:x/:y"); got != "p" {
		t.Fatalf("got %q", got)
	}

	if got := resolveParamName("Query", "", "q", "/x"); got != "q" {
		t.Fatalf("got %q", got)
	}
}
