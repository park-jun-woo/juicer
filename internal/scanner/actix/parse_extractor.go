//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 함수 파라미터 하나에서 extractor 정보를 파싱한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func parseExtractor(param *sitter.Node, src []byte) *extractorInfo {
	typeNode := findParamType(param)
	if typeNode == nil {
		return nil
	}
	if typeNode.Type() != "generic_type" {
		return nil
	}
	return buildGenericExtractor(typeNode, src)
}
