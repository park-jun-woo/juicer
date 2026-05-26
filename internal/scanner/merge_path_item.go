//ff:func feature=scan type=extract control=sequence
//ff:what scan과 base의 path item을 operation 수준에서 병합한다
package scanner

import "gopkg.in/yaml.v3"

func mergePathItem(scanOps *yaml.Node, baseOps *yaml.Node) *yaml.Node {
	if scanOps.Kind != yaml.MappingNode || baseOps.Kind != yaml.MappingNode {
		return scanOps
	}
	result := &yaml.Node{Kind: yaml.MappingNode}
	added := mergeBaseOpsInto(result, baseOps, scanOps)
	appendScanOnlyOps(result, scanOps, added)
	return result
}
