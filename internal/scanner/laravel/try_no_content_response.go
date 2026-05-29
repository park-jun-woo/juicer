//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what response()->noContent() 형태의 204 빈 응답을 추출한다
package laravel

import (
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func tryNoContentResponse(text string) *scanner.Response {
	if !strings.Contains(text, "noContent()") {
		return nil
	}
	return &scanner.Response{
		Status: "204",
		Kind:   "empty",
	}
}
