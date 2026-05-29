//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 함수 시그니처에서 web::Path, web::Json, web::Query, web::Form extractor를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func extractExtractors(funcNode *sitter.Node, src []byte) []extractorInfo {
	params := findChildByType(funcNode, "parameters")
	if params == nil {
		return nil
	}

	var result []extractorInfo
	for _, param := range childrenOfType(params, "parameter") {
		ext := parseExtractor(param, src)
		if ext != nil {
			result = append(result, *ext)
		}
	}
	return result
}
