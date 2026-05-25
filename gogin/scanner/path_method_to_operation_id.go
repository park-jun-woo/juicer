//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what path와 method에서 operationId를 조합한다
package scanner

import (
	"strings"
)

func pathMethodToOperationID(method, path string) string {
	// /api/v1/admin/buildings/:buildingId → buildings_buildingId
	segments := strings.Split(path, "/")
	var parts []string
	for _, seg := range segments {
		if seg == "" || seg == "api" || strings.HasPrefix(seg, "v") && len(seg) <= 3 {
			continue
		}
		seg = strings.TrimPrefix(seg, ":")
		seg = strings.TrimPrefix(seg, "*")
		parts = append(parts, seg)
	}

	id := strings.ToLower(method) + "_" + strings.Join(parts, "_")
	return id
}

