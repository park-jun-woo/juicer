//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what scan operation의 필드를 순회하며 base에서 보존할 값을 병합한다
package scanner

import "gopkg.in/yaml.v3"

func mergeScanOpFields(result *yaml.Node, scanOp *yaml.Node, baseOp *yaml.Node) map[string]bool {
	added := map[string]bool{}
	for i := 0; i+1 < len(scanOp.Content); i += 2 {
		key := scanOp.Content[i].Value
		val := mergeOpField(key, scanOp.Content[i+1], baseOp)
		result.Content = append(result.Content, scanOp.Content[i], val)
		added[key] = true
	}
	return added
}
