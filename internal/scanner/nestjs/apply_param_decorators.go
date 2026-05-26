//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 데코레이터 목록을 해석하여 파라미터 종류별로 결과에 추가한다
package nestjs

import "github.com/park-jun-woo/juicer/internal/scanner"

// applyParamDecorators dispatches each decorator into the appropriate result bucket.
func applyParamDecorators(decorators []decoratorInfo, paramName, paramType string, result *methodParams) {
	for _, d := range decorators {
		if d.name == DecQuery && isQueryDTO(d.arg, paramType) {
			result.queryDTOType = paramType
			continue
		}
		name := d.arg
		if name == "" {
			name = paramName
		}
		switch d.name {
		case DecParam:
			result.pathParams = append(result.pathParams, scanner.Param{
				Name: name, Type: tsTypeToOpenAPIType(paramType),
			})
		case DecQuery:
			result.queryParams = append(result.queryParams, scanner.Param{
				Name: name, Type: tsTypeToOpenAPIType(paramType),
			})
		case DecBody:
			result.bodyType = paramType
		case DecUploadedFile:
			result.files = append(result.files, scanner.Param{
				Name: name, Type: "file",
			})
		}
	}
}
