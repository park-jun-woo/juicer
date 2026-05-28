//ff:type feature=scan type=model topic=express
//ff:what 스캔 컨텍스트 구조체 (Pass 1 결과)
package express

type scanContext struct {
	parsed      map[string]*fileInfo
	allRouters  map[string]map[string]bool
	prefixMap   map[string]string
	absRoot     string
	pathAliases map[string]string
}
