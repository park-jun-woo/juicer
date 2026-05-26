//ff:func feature=scan type=convert control=iteration dimension=1 topic=fastapi
//ff:what 경로 세그먼트를 슬래시로 결합한다
package fastapi

import "strings"

// joinPath joins path segments, ensuring single slashes between segments.
func joinPath(parts ...string) string {
	var segments []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		segments = append(segments, strings.Trim(p, "/"))
	}
	if len(segments) == 0 {
		return ""
	}
	return "/" + strings.Join(segments, "/")
}
