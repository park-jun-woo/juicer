//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 클래스 body에서 @action 데코레이터 메서드를 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// extractActions finds @action decorated methods in a class body.
func extractActions(body *sitter.Node, src []byte, relPath string) []actionInfo {
	var actions []actionInfo
	for _, decDef := range findAllByType(body, "decorated_definition") {
		ai := parseDecoratedAction(decDef, src)
		if ai != nil {
			actions = append(actions, *ai)
		}
	}
	return actions
}
