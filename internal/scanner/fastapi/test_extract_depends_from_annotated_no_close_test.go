//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDependsFromAnnotated_NoClose 테스트
package fastapi

import "testing"

func TestExtractDependsFromAnnotated_NoClose(t *testing.T) {
	if got := extractDependsFromAnnotated("Annotated[x, Depends("); got != "" {
		t.Fatalf("got %q", got)
	}
}
