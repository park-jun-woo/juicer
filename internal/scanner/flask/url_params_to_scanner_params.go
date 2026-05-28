//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what urlParam 슬라이스를 scanner.Param 슬라이스로 변환한다
package flask

import "github.com/park-jun-woo/codistill/internal/scanner"

// urlParamsToScannerParams converts urlParam slice to scanner.Param slice.
func urlParamsToScannerParams(params []urlParam) []scanner.Param {
	var result []scanner.Param
	for _, p := range params {
		oaType := flaskConverterToOpenAPI(p.converter)
		result = append(result, scanner.Param{
			Name: p.name,
			Type: oaType.Type,
		})
	}
	return result
}
