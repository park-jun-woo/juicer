//ff:func feature=hurl type=parse control=sequence
//ff:what 엔드포인트 ID("METHOD /path")에서 method와 path를 분리
package hurls

import (
	"strings"
)

// parseEndpointID splits "METHOD /path" into method and path.
func parseEndpointID(id string) (string, string) {
	parts := strings.SplitN(id, " ", 2)
	if len(parts) != 2 {
		return id, ""
	}
	return parts[0], parts[1]
}
