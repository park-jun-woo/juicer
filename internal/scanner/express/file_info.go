//ff:type feature=scan type=model topic=express
//ff:what 파싱된 TypeScript 파일 정보 구조체
package express

import sitter "github.com/smacker/go-tree-sitter"

type fileInfo struct {
	Path string
	Root *sitter.Node
	Src  []byte
}
