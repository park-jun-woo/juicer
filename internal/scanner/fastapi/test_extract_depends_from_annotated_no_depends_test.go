//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDependsFromAnnotated_NoDepends 테스트
package fastapi

import "testing"

func TestExtractDependsFromAnnotated_NoDepends(t *testing.T) {
	if got := extractDependsFromAnnotated("Annotated[int, Query()]"); got != "" {
		t.Fatalf("got %q", got)
	}
}
