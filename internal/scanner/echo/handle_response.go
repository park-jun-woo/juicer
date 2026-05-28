//ff:func feature=scan type=extract control=selection
//ff:what c.JSON(status, body) 등의 응답 호출에서 상태 코드와 body를 추출한다
package echo

import (
	"go/ast"
	"go/types"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func handleResponse(ep *scanner.Endpoint, call *ast.CallExpr, kind string, info *types.Info, source string) {
	resp := scanner.Response{
		Kind: kind,
	}

	if source != "handler" {
		resp.Source = source
	}

	switch kind {
	case "json":
		// c.JSON(status, body) / c.JSONPretty(status, body, indent)
		if len(call.Args) >= 1 {
			resp.Status = resolveStatusCode(call.Args[0], info)
		}
		if len(call.Args) >= 2 {
			resp.Body = exprString(call.Args[1])
			typeName, fields, confidence := resolveResponseType(call.Args[1], info)
			resp.TypeName = typeName
			resp.Fields = fields
			resp.Confidence = confidence
		}

	case "string", "html":
		// c.String(status, format, ...) / c.HTML(status, html)
		if len(call.Args) >= 1 {
			resp.Status = resolveStatusCode(call.Args[0], info)
		}

	case "data":
		// c.Blob(status, contentType, data) / c.Stream(status, contentType, reader)
		if len(call.Args) >= 1 {
			resp.Status = resolveStatusCode(call.Args[0], info)
		}

	case "file":
		// c.File(path) — 상태 코드 없음, 200 기본
		resp.Status = "200"

	case "redirect":
		// c.Redirect(status, url)
		if len(call.Args) >= 1 {
			resp.Status = resolveStatusCode(call.Args[0], info)
		}

	case "status":
		// c.NoContent(status)
		if len(call.Args) >= 1 {
			resp.Status = resolveStatusCode(call.Args[0], info)
		}
	}

	if resp.Status == "" {
		resp.Status = "(unknown)"
	}

	ep.Responses = append(ep.Responses, resp)
}
