//ff:func feature=sql type=render control=sequence
//ff:what 결과를 YAML로 렌더링
package sqls

import "gopkg.in/yaml.v3"

// RenderYAML renders the result as YAML.
func RenderYAML(result *SkeletonResult) ([]byte, error) {
	return yaml.Marshal(result)
}
