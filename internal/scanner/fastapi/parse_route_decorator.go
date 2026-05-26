//ff:func feature=scan type=parse control=sequence topic=fastapi
//ff:what 데코레이터 노드에서 HTTP 메서드, 경로, 라우터 변수, status_code, response_model을 파싱한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// parseRouteDecorator parses a decorator like @app.get("/users/{user_id}", status_code=200).
// Returns (method, path, routerVar, statusCode, responseModel).
func parseRouteDecorator(dec *sitter.Node, src []byte) (string, string, string, int, string) {
	callNode, attrNode := findDecoratorNodes(dec)
	if callNode != nil && attrNode == nil {
		attrNode = findChildByType(callNode, "attribute")
	}
	if attrNode == nil {
		return "", "", "", 0, ""
	}

	routerVar, httpMethod := parseAttribute(attrNode, src)
	if httpMethod == "" {
		return "", "", "", 0, ""
	}

	path, statusCode, responseModel := extractDecoratorArgs(callNode, src)
	return httpMethod, path, routerVar, statusCode, responseModel
}
