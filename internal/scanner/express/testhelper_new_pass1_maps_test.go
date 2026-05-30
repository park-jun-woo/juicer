//ff:func feature=scan type=test control=sequence topic=express
//ff:what newPass1Maps 테스트 헬퍼
package express

import sitter "github.com/smacker/go-tree-sitter"

func newPass1Maps() (map[string]*fileInfo, map[string]map[string]bool, map[string]*sitter.Node, map[string][]byte) {
	return map[string]*fileInfo{}, map[string]map[string]bool{}, map[string]*sitter.Node{}, map[string][]byte{}
}
