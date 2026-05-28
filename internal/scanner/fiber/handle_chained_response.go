//ff:func feature=scan type=extract control=selection
//ff:what c.Status(N).JSON(body) 패턴의 체이닝된 응답을 추출한다
package fiber

import (
	"go/ast"
	"go/types"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// handleChainedResponse handles c.Status(N).JSON(body) pattern
func handleChainedResponse(ep *scanner.Endpoint, statusCall *ast.CallExpr, outerCall *ast.CallExpr, methodName string, info *types.Info, source string) {
	resp := scanner.Response{}

	if source != "handler" {
		resp.Source = source
	}

	// Extract status code from c.Status(N)
	if len(statusCall.Args) >= 1 {
		resp.Status = resolveStatusCode(statusCall.Args[0], info)
	}

	switch methodName {
	case "JSON":
		resp.Kind = "json"
		if len(outerCall.Args) >= 1 {
			resp.Body = exprString(outerCall.Args[0])
			typeName, fields, confidence := resolveResponseType(outerCall.Args[0], info)
			resp.TypeName = typeName
			resp.Fields = fields
			resp.Confidence = confidence
		}
	case "SendString":
		resp.Kind = "string"
	case "Send":
		resp.Kind = "data"
	default:
		resp.Kind = methodName
	}

	if resp.Status == "" {
		resp.Status = "(unknown)"
	}

	ep.Responses = append(ep.Responses, resp)
}
