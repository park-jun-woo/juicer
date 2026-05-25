//ff:func feature=sql type=parse control=sequence
//ff:what 결과를 JSON으로 렌더링
package sqls

import (
	"encoding/json"
)

// RenderJSON renders the result as JSON.
//
func RenderJSON(result *SkeletonResult) ([]byte, error) {
	return json.MarshalIndent(result, "", "  ")
}

