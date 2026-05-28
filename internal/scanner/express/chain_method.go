//ff:type feature=scan type=model topic=express
//ff:what 체인 메서드 정보 구조체
package express

import sitter "github.com/smacker/go-tree-sitter"

type chainMethod struct {
	method      string
	handler     string
	handlerNode *sitter.Node
	middleware  []string
	line        int
	authLevel   string
	roles       []string
}
