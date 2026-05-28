//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what URL path parameter를 엔드포인트에 추가한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// addPathParams adds URL path parameters to an endpoint.
func addPathParams(ep *scanner.Endpoint, params []urlParam) {
	if len(params) == 0 {
		return
	}
	if ep.Request == nil {
		ep.Request = &scanner.Request{}
	}
	for _, p := range params {
		oaType := djangoConverterToOpenAPI(p.converter)
		ep.Request.PathParams = append(ep.Request.PathParams, scanner.Param{
			Name: p.name,
			Type: oaType.Type,
		})
	}
}
