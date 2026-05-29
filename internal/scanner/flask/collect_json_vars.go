//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 함수 본문에서 request.get_json()/request.json 으로 대입된 변수명을 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// collectJSONVars finds local variables assigned from a JSON source
// (e.g. `data = request.get_json()`), returning a set of their names.
func collectJSONVars(funcDef *sitter.Node, src []byte) map[string]bool {
	vars := make(map[string]bool)
	for _, asgn := range findAllByType(funcDef, "assignment") {
		name := assignmentJSONVar(asgn, src)
		if name != "" {
			vars[name] = true
		}
	}
	return vars
}
