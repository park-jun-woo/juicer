//ff:type feature=scan type=model topic=actix
//ff:what proc-macro 어트리뷰트로 추출한 라우트(method/path/handler/funcNode/file)
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type macroRoute struct {
	method   string
	path     string
	handler  string
	funcNode *sitter.Node
	file     *fileInfo
}
