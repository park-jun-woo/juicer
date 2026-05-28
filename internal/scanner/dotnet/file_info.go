//ff:type feature=scan type=model topic=dotnet
//ff:what C# 파일 파싱 결과 구조체
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

type fileInfo struct {
	absPath     string
	relPath     string
	projectRoot string
	src         []byte
	root        *sitter.Node
	usings      []string
}
