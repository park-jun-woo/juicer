//ff:type feature=scan type=model topic=hono
//ff:what 라우트 그룹 정보 구조체
package hono

type routeGroup struct {
	Prefix     string
	ParentVar  string
	SubAppName string
	SourceFile string
}
