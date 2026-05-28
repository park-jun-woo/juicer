//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what [FromForm] 파라미터를 file 또는 form field로 분류한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func classifyFormParam(ep *endpointInfo, typeName, paramName string) {
	if typeName == "IFormFile" {
		ep.files = append(ep.files, scanner.Param{
			Name: paramName,
			Type: "file",
		})
	} else {
		ep.formFields = append(ep.formFields, scanner.Param{
			Name: paramName,
			Type: csharpTypeToOpenAPIType(typeName),
		})
	}
}
