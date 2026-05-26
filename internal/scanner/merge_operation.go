//ff:func feature=scan type=extract control=sequence
//ff:what 단일 operation을 병합한다 (코드가 구조, base가 설명)
package scanner

import "gopkg.in/yaml.v3"

func mergeOperation(scanOp *yaml.Node, baseOp *yaml.Node) *yaml.Node {
	if scanOp == nil {
		return baseOp
	}
	if baseOp == nil {
		return scanOp
	}

	result := &yaml.Node{Kind: yaml.MappingNode}
	added := mergeScanOpFields(result, scanOp, baseOp)
	appendBaseOnlyPreserved(result, baseOp, added)
	return result
}
