//ff:type feature=scan type=model topic=actix
//ff:what 핸들러 함수 노드와 소속 파일 정보
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type handlerInfo struct {
	funcNode *sitter.Node
	file     *fileInfo
}
