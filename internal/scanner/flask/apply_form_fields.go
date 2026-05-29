//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what form 필드 키들을 ep.Request.FormFields로 설정한다
package flask

import "github.com/park-jun-woo/codistill/internal/scanner"

// applyFormFields sets multipart/form-data fields on the endpoint request.
func applyFormFields(ep *scanner.Endpoint, keys []string) {
	if len(keys) == 0 {
		return
	}
	scanner.EnsureRequest(ep)
	for _, k := range keys {
		ep.Request.FormFields = append(ep.Request.FormFields, scanner.Param{Name: k, Type: "string"})
	}
}
