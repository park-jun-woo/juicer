//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what scan에만 있는 operation을 result에 추가한다
package scanner

import "gopkg.in/yaml.v3"

func appendScanOnlyOps(result *yaml.Node, scanOps *yaml.Node, added map[string]bool) {
	for i := 0; i+1 < len(scanOps.Content); i += 2 {
		method := scanOps.Content[i].Value
		if added[method] {
			continue
		}
		result.Content = append(result.Content, scanOps.Content[i], scanOps.Content[i+1])
	}
}
