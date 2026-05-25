//ff:func feature=hurl type=parse control=iteration dimension=1
//ff:what 엔드포인트 경로에서 Hurl 테스트 파일명 추천
package hurls

import (
	"strings"
)

// suggestFilename generates a suggested .hurl filename from the path.
// e.g. "/api/v1/admin/buildings" -> "buildings.hurl"
// e.g. "/api/v1/admin/buildings/:id" -> "buildings_id.hurl"
// e.g. "/api/health" -> "health.hurl"
func suggestFilename(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")

	var meaningful []string
	for _, p := range parts {
		if p == "api" || p == "admin" || (len(p) <= 2 && strings.HasPrefix(p, "v")) {
			continue
		}
		if strings.HasPrefix(p, ":") {
			meaningful = append(meaningful, strings.TrimPrefix(p, ":"))
			continue
		}
		meaningful = append(meaningful, p)
	}

	if len(meaningful) == 0 {
		return "test.hurl"
	}
	return strings.Join(meaningful, "_") + ".hurl"
}
