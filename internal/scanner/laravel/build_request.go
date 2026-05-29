//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what path 파라미터와 FormRequest로부터 요청 객체를 구성한다(없으면 nil)
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func buildRequest(absRoot string, pathParams []scanner.Param, cm *controllerMethod, parsedFiles map[string]*fileInfo) *scanner.Request {
	hasFormRequest := cm != nil && cm.formRequestRef != ""
	if len(pathParams) == 0 && !hasFormRequest {
		return nil
	}
	req := &scanner.Request{
		PathParams: pathParams,
	}
	if hasFormRequest {
		req.Body = buildFormRequestBody(absRoot, cm, parsedFiles)
	}
	return req
}
