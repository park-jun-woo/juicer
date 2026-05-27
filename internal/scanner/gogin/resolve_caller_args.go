//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what callee 파라미터와 caller 인자를 위치 기반으로 매핑하여 상태코드와 응답 데이터 타입을 해석한다
package gogin

import (
	"go/ast"
	"go/types"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveCallerArgs(fnDecl *ast.FuncDecl, call *ast.CallExpr, callerInfo *types.Info, calleeInfo *types.Info) (status string, typeName string, fields []scanner.Field, confidence string) {
	if fnDecl.Type.Params == nil || calleeInfo == nil {
		return
	}

	// callee 파라미터를 평탄화
	type paramEntry struct {
		name string
		typ  types.Type
	}
	var params []paramEntry
	for _, field := range fnDecl.Type.Params.List {
		var fieldType types.Type
		if len(field.Names) > 0 {
			if obj := calleeInfo.Defs[field.Names[0]]; obj != nil {
				fieldType = obj.Type()
			}
		}
		if fieldType == nil {
			if tv, ok := calleeInfo.Types[field.Type]; ok {
				fieldType = tv.Type
			}
		}

		if len(field.Names) == 0 {
			params = append(params, paramEntry{"", fieldType})
		} else {
			for _, name := range field.Names {
				params = append(params, paramEntry{name.Name, fieldType})
			}
		}
	}

	n := len(params)
	if len(call.Args) < n {
		n = len(call.Args)
	}

	for i := 0; i < n; i++ {
		if params[i].typ == nil {
			continue
		}
		r := resolveCallerArg(params[i].typ, call.Args[i], callerInfo)
		if r.status != "" {
			status = r.status
		}
		if r.typeName != "" || len(r.fields) > 0 {
			typeName = r.typeName
			fields = r.fields
			confidence = r.confidence
		}
	}

	return
}
