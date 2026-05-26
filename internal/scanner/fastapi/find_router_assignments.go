//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what FastAPI()/APIRouter() 인스턴스 할당문을 찾는다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findRouterAssignments finds assignments like:
//
//	app = FastAPI()
//	router = APIRouter(prefix="/users")
func findRouterAssignments(root *sitter.Node, src []byte) []routerInfo {
	var routers []routerInfo
	assignments := findAllByType(root, "assignment")
	for _, assign := range assignments {
		ri := tryParseRouterAssignment(assign, src)
		if ri != nil {
			routers = append(routers, *ri)
		}
	}
	return routers
}
