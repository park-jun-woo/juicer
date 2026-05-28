//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what Blueprint 인스턴스와 url_prefix를 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// collectBlueprints finds all Blueprint(...) assignments and extracts variable name + url_prefix.
// e.g., users_bp = Blueprint("users", __name__, url_prefix="/api/users")
func collectBlueprints(root *sitter.Node, src []byte) []blueprintInfo {
	var blueprints []blueprintInfo
	assignments := findAllByType(root, "assignment")
	for _, assign := range assignments {
		bp := tryParseBlueprintAssignment(assign, src)
		if bp != nil {
			blueprints = append(blueprints, *bp)
		}
	}
	return blueprints
}
