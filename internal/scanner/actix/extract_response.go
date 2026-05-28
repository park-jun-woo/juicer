//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 함수 본문에서 HttpResponse 상태 코드와 응답 종류를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractResponses(funcNode *sitter.Node, src []byte) []scanner.Response {
	block := findChildByType(funcNode, "block")
	if block == nil {
		return nil
	}

	seen := map[string]bool{}
	var responses []scanner.Response

	walkNodes(block, func(n *sitter.Node) {
		if n.Type() != "scoped_identifier" {
			return
		}
		text := nodeText(n, src)
		parts := splitScoped(text)
		if len(parts) != 2 || parts[0] != "HttpResponse" {
			return
		}
		statusName := parts[1]
		code, ok := httpResponseStatuses[statusName]
		if !ok {
			return
		}
		if seen[code] {
			return
		}
		seen[code] = true

		kind := detectResponseKind(n, src)
		responses = append(responses, scanner.Response{
			Status: code,
			Kind:   kind,
		})
	})

	return responses
}

func detectResponseKind(scopedID *sitter.Node, src []byte) string {
	// Walk up to find call_expression parent, then check if .json() is chained
	parent := scopedID.Parent()
	if parent == nil {
		return ""
	}
	// scopedID -> call_expression (HttpResponse::Ok()) -> field_expression (.json) -> call_expression (.json(...))
	// walk the parent chain to find field_expression with "json" or "finish"
	grandParent := parent.Parent()
	if grandParent != nil && grandParent.Type() == "field_expression" {
		fieldID := findChildByType(grandParent, "field_identifier")
		if fieldID != nil {
			fname := nodeText(fieldID, src)
			if fname == "json" {
				return "json"
			}
		}
	}
	// check one more level
	if grandParent != nil {
		ggParent := grandParent.Parent()
		if ggParent != nil && ggParent.Type() == "field_expression" {
			fieldID := findChildByType(ggParent, "field_identifier")
			if fieldID != nil {
				fname := nodeText(fieldID, src)
				if fname == "json" {
					return "json"
				}
			}
		}
	}
	return ""
}

func splitScoped(s string) []string {
	var parts []string
	current := ""
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == ':' && s[i+1] == ':' {
			parts = append(parts, current)
			current = ""
			i++
			continue
		}
		current += string(s[i])
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}
