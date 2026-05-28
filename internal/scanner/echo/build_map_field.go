//ff:func feature=scan type=extract control=sequence
//ff:what map[string]any 리터럴의 단일 키-값 쌍에서 Field를 생성한다
package echo

import (
	"go/ast"
	"go/token"
	"go/types"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

// buildMapField creates a Field from a map key-value pair. Returns nil if the element is not a KV or has no key.
func buildMapField(elt ast.Expr, info *types.Info) *scanner.Field {
	kv, ok := elt.(*ast.KeyValueExpr)
	if !ok {
		return nil
	}

	key := ""
	if lit, ok := kv.Key.(*ast.BasicLit); ok && lit.Kind == token.STRING {
		key = unquote(lit.Value)
	} else if id, ok := kv.Key.(*ast.Ident); ok {
		key = id.Name
	}
	if key == "" {
		return nil
	}

	field := scanner.Field{
		Name: key,
		JSON: key,
		Type: inferValueType(kv.Value, info),
	}

	if nested, ok := kv.Value.(*ast.CompositeLit); ok && isMapStringAny(nested, info) {
		field.Fields = extractMapFields(nested, info)
		field.Type = "object"
	}

	return &field
}
