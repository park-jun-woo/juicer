//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what base에만 있는 보존 대상 필드를 result에 추가한다
package scanner

import "gopkg.in/yaml.v3"

func appendBaseOnlyPreserved(result *yaml.Node, baseOp *yaml.Node, added map[string]bool) {
	for i := 0; i+1 < len(baseOp.Content); i += 2 {
		key := baseOp.Content[i].Value
		if added[key] || !preserveFromBase[key] {
			continue
		}
		result.Content = append(result.Content, baseOp.Content[i], baseOp.Content[i+1])
	}
}
