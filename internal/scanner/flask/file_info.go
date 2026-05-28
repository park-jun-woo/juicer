//ff:type feature=scan type=model topic=flask
//ff:what 파싱된 Python 파일 정보 구조체
package flask

import sitter "github.com/smacker/go-tree-sitter"

// fileInfo holds parsed information for a single Python source file.
type fileInfo struct {
	absPath string
	relPath string
	src     []byte
	root    *sitter.Node
}
