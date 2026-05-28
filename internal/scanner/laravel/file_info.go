//ff:type feature=scan type=model topic=laravel
//ff:what 파싱된 PHP 파일 정보 구조체
package laravel

import sitter "github.com/smacker/go-tree-sitter"

type fileInfo struct {
	absPath string
	relPath string
	src     []byte
	root    *sitter.Node
}
