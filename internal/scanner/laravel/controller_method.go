//ff:type feature=scan type=model topic=laravel
//ff:what 컨트롤러 메서드 정보(이름/파라미터/FormRequest 참조/return 노드)
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// controllerMethod holds information about a controller method.
type controllerMethod struct {
	name           string
	params         []methodParam
	formRequestRef string         // FormRequest type hint if any
	returnNodes    []*sitter.Node // return statement nodes
	src            []byte
}
