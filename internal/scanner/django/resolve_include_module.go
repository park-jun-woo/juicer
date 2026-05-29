//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what include 대상 dotted 모듈명을 수집된 모듈 맵에서 접미사 매칭으로 해석한다
package django

import "strings"

// resolveIncludeModule resolves an include("app.urls") target to a key in byModule.
// Matching is by dotted suffix: "app.urls" matches a module whose path ends with ".app.urls" or equals "app.urls".
func resolveIncludeModule(target string, byModule map[string][]urlEntry) (string, bool) {
	if target == "" {
		return "", false
	}
	if _, ok := byModule[target]; ok {
		return target, true
	}
	for mod := range byModule {
		if mod == target || strings.HasSuffix(mod, "."+target) {
			return mod, true
		}
	}
	return "", false
}
