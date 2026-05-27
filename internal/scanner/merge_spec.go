//ff:func feature=scan type=extract control=sequence
//ff:what 스캔 결과와 기존 base spec을 병합한다 (라우터 등록이 진실)
package scanner

import "gopkg.in/yaml.v3"

func mergeSpec(scanNode *yaml.Node, baseNode *yaml.Node, scanResult *ScanResult) *yaml.Node {
	registered := buildRegisteredSet(scanResult)
	merged := shallowCopyMapping(baseNode)

	basePaths := findMappingValue(baseNode, "paths")
	scanPaths := findMappingValue(scanNode, "paths")
	mergedPaths := mergeAllPaths(basePaths, scanPaths, registered)
	setMappingValue(merged, "paths", mergedPaths)

	scanSchemas := findComponentSchemas(scanNode)
	baseSchemas := findComponentSchemas(baseNode)
	if scanSchemas != nil || baseSchemas != nil {
		ms := mergeSchemas(scanSchemas, baseSchemas)
		setComponentSchemas(merged, ms)
	}

	scanSecSchemes := findComponentSecuritySchemes(scanNode)
	baseSecSchemes := findComponentSecuritySchemes(baseNode)
	if scanSecSchemes != nil || baseSecSchemes != nil {
		ms := mergeSchemas(scanSecSchemes, baseSecSchemes)
		setComponentSecuritySchemes(merged, ms)
	}

	return merged
}
