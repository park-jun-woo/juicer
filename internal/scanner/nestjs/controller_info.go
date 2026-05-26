//ff:type feature=scan type=model topic=nestjs
//ff:what 컨트롤러 추출 중간 결과 구조체
package nestjs

// controllerInfo holds intermediate route extraction results.
type controllerInfo struct {
	prefix          string
	version         string
	classMiddleware []string // class-level @UseGuards guards
	classRoles      []string // class-level @Roles values
	endpoints       []endpointInfo
	imports         map[string]string // typeName -> relative import path
}
