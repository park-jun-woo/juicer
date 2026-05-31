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
	// 정확 일치(레거시 set)는 그대로 우선 확인한다.
	if authMiddlewareNames[lower] {
		return "auth"
	}
	if roleMiddlewareNames[lower] {
		return "role"
	}
	// 부분일치 기반 확장 (BUG-010 / Phase139).
	// extractMiddlewareNameForAuth는 member_expression을 전체 텍스트로 넘기므로
	// (예: `mw.authAdminApi` → `mw.authadminapi`) 객체 한정자(`mw.`)가 앞에 붙는다.
	// 따라서 HasPrefix가 아니라 Contains로 매칭해야 camelCase·접두형 실무 명명을 잡는다.
	//
	// 역할(role) 미들웨어를 인증(auth)보다 먼저 검사한다: `authorize`처럼 auth 부분문자열을
	// 포함하는 권한 미들웨어를 role로 정확히 분류하기 위함.
	if strings.Contains(lower, "authorize") ||
		strings.Contains(lower, "requirerole") ||
		strings.Contains(lower, "checkrole") ||
		strings.Contains(lower, "hasrole") ||
		strings.Contains(lower, "allowroles") ||
		strings.Contains(lower, "rbac") {
		return "role"
	}
	// `*guard*` (NestJS-스타일 guard 미들웨어 명명) 도 인증으로 간주.
	if strings.Contains(lower, "guard") {
		return "auth"
	}
	// `require*`/`ensure*` 접두형은 Auth/Authenticated/Logged 어근이 붙은 경우만 인증으로
	// 본다 (requireAll/ensureDir/ensureSettings 등 비-auth 헬퍼 FP 방지).
	if strings.HasPrefix(lower, "require") || strings.HasPrefix(lower, "ensure") {
		if strings.Contains(lower, "auth") || strings.Contains(lower, "logged") || strings.Contains(lower, "login") {
			return "auth"
		}
	}
	// 핵심 차단 해소: 이름에 `auth`가 포함되면 인증 미들웨어로 본다.
	// 이로써 `mw.authAdminApi`/`authAdminApiWithUrl`/`authenticatePublic` 등 Ghost admin API
	// 미들웨어가 잡힌다. 트레이드오프: `author`/`authorBio` 등 auth와 무관하게 "author"를
	// 포함하는 이름은 false positive로 auth_required가 될 수 있다(recall 우선 선택). 다만
	// `authorize`는 위에서 role로 선분류되므로 영향 없음.
	if strings.Contains(lower, "auth") {
		return "auth"
	}
	return ""
}
