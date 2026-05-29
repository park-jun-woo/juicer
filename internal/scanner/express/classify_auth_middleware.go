//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 미들웨어 이름을 인증/권한/무관으로 분류한다
package express

import "strings"

var authMiddlewareNames = map[string]bool{
	"auth":                true,
	"authenticate":        true,
	"requireauth":         true,
	"ensureauthenticated": true,
	"isauthenticated":     true,
	"authmiddleware":      true,
	"jwtauth":             true,
	"verifytoken":         true,
}

var roleMiddlewareNames = map[string]bool{
	"authorize":  true,
	"requirerole": true,
	"checkrole":  true,
	"hasrole":    true,
	"allowroles": true,
	"rbac":       true,
}

func classifyAuthMiddleware(name string) string {
	lower := strings.ToLower(name)
	if lower == "passport.authenticate" {
		return "auth"
	}
	if authMiddlewareNames[lower] {
		return "auth"
	}
	if roleMiddlewareNames[lower] {
		return "role"
	}
	return ""
}
