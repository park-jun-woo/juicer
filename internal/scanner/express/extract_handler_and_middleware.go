//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 인자 노드 슬라이스에서 핸들러명과 미들웨어 목록을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractHandlerAndMiddleware(argNodes []*sitter.Node, src []byte) (string, []string) {
	if len(argNodes) < 2 {
		return "", nil
	}
	handler := extractHandlerName(argNodes[len(argNodes)-1], src)
	var middleware []string
	for i := 1; i < len(argNodes)-1; i++ {
		mw := extractMiddlewareName(argNodes[i], src)
		if mw != "" {
			middleware = append(middleware, mw)
		}
	}
	return handler, middleware
}
