//ff:func feature=scan type=extract control=selection
//ff:what c.JSON(body), c.SendStatus(N) 등의 응답 호출에서 상태 코드와 body를 추출한다
package fiber

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
		// Fiber: c.JSON(body) — single argument, status 200 implied
		resp.Status = "200"
		if len(call.Args) >= 1 {
			resp.Body = exprString(call.Args[0])
			typeName, fields, confidence := resolveResponseType(call.Args[0], info)
			resp.TypeName = typeName
			resp.Fields = fields
			resp.Confidence = confidence
		}

	case "string":
		// c.SendString(s)
		resp.Status = "200"

	case "data":
		// c.Send(data)
		resp.Status = "200"

	case "file":
		// c.SendFile(path)
		resp.Status = "200"

	case "redirect":
		// c.Redirect(url, status) — Fiber redirect takes URL first, status second
		resp.Status = "302" // default
		if len(call.Args) >= 2 {
			resp.Status = resolveStatusCode(call.Args[1], info)
		}

	case "status":
		// c.Status(code) — just sets status, or c.SendStatus(code) — sends with empty body
		if len(call.Args) >= 1 {
			resp.Status = resolveStatusCode(call.Args[0], info)
		}
		// c.Status() is a chaining call, not a terminal response by itself
		// Only record if it's SendStatus (terminal)
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "Status" {
			// c.Status() is typically followed by .JSON(), .SendString(), etc
			// Don't record it as a standalone response
			return
		}
	}

	if resp.Status == "" {
		resp.Status = "(unknown)"
	}

	ep.Responses = append(ep.Responses, resp)
}

