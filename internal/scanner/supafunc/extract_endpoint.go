//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what 함수 파일 하나에서 엔드포인트를 추출한다 (디렉토리명 → 경로, 본문 → 메서드/요청/응답)
package supafunc

import (
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractEndpoint(fi *fileInfo, root string) []scanner.Endpoint {
	dirName := filepath.Base(filepath.Dir(fi.Path))
	path := "/" + dirName

	relPath, _ := filepath.Rel(root, fi.Path)
	if relPath == "" {
		relPath = fi.Path
	}

	callbackBody, handler := findServeCallback(fi)
	if callbackBody == nil {
		return []scanner.Endpoint{{
			Method:  "POST",
			Path:    path,
			Handler: "serve",
			File:    relPath,
		}}
	}

	methods := extractHTTPMethods(callbackBody, fi.Src)
	if len(methods) == 0 {
		methods = []string{"POST"}
	}

	methodBlocks := extractMethodBlock(callbackBody, fi.Src)

	// 블록이 없는 메서드에 대해 body 폴백 차단 여부 판정
	skipFallbackBody := len(methodBlocks) > 0 && allJSONInsideBlocks(callbackBody, fi.Src, methodBlocks)

	var endpoints []scanner.Endpoint
	for _, m := range methods {
		if mb, ok := methodBlocks[m]; ok {
			ep := extractEndpointForMethod(mb, fi.Src, m, path, handler, relPath)
			endpoints = append(endpoints, ep)
		} else if skipFallbackBody {
			ep := extractEndpointForMethodNoBody(callbackBody, fi.Src, m, path, handler, relPath)
			endpoints = append(endpoints, ep)
		} else {
			ep := extractEndpointForMethod(callbackBody, fi.Src, m, path, handler, relPath)
			endpoints = append(endpoints, ep)
		}
	}
	return endpoints
}
