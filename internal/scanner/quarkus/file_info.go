//ff:type feature=scan type=model topic=quarkus
//ff:what Java 파일 파싱 결과 구조체
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

type fileInfo struct {
	absPath     string
	relPath     string
	projectRoot string
	src         []byte
	root        *sitter.Node
	imports     map[string]string
}
