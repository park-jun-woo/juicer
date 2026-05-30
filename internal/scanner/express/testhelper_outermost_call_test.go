//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what outermostCall 테스트 헬퍼
package express

import sitter "github.com/smacker/go-tree-sitter"

// outermostCall returns the call_expression that is not the object of another call.
func outermostCall(fi *fileInfo) *sitter.Node {
	calls := findAllByType(fi.Root, "call_expression")
	for _, c := range calls {
		if !isInnerCall(c) {
			return c
		}
	}
	if len(calls) > 0 {
		return calls[0]
	}
	return nil
}
