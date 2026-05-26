//ff:func feature=scan type=test control=sequence
//ff:what forwardRouterCalls가 ctx.info 대신 info(callee TypesInfo)를 rescan에 전달하는지 검증한다
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

// TestForwardRouterCalls_UsesCalleeInfo verifies that forwardRouterCalls
// propagates the callee's TypesInfo (info param) into the rescan context.
// Limitation: single-package tests cannot trigger real cross-package mismatch;
// we simulate it with two distinct *types.Info having different Uses maps.
func TestForwardRouterCalls_UsesCalleeInfo(t *testing.T) {
	innerIdent := &ast.Ident{Name: "inner"}

	fset := token.NewFileSet()
	file := fset.AddFile("callee.go", -1, 100)
	targetPos := file.Pos(10)
	targetObj := types.NewFunc(targetPos, nil, "inner",
		types.NewSignatureType(nil, nil, nil, nil, nil, false))

	// calleeInfo knows about the inner call; callerInfo does not.
	calleeInfo := &types.Info{
		Uses: map[*ast.Ident]types.Object{innerIdent: targetObj},
	}
	callerInfo := &types.Info{
		Uses: map[*ast.Ident]types.Object{},
	}

	innerFnDecl := &ast.FuncDecl{
		Name: &ast.Ident{Name: "inner"},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{{
					Names: []*ast.Ident{{Name: "r"}},
					Type:  &ast.Ident{Name: "Router"},
				}},
			},
		},
		Body: &ast.BlockStmt{},
	}

	idx := &funcIndex{
		byPos: map[token.Pos]*ast.FuncDecl{targetPos: innerFnDecl},
		info:  map[token.Pos]*types.Info{targetPos: calleeInfo},
	}

	call := &ast.CallExpr{
		Fun:  innerIdent,
		Args: []ast.Expr{&ast.Ident{Name: "router"}},
	}
	stmts := []ast.Stmt{&ast.ExprStmt{X: call}}

	ctx := &groupArgCtx{
		ginAlias: "gin",
		routers:  map[string]*routerInfo{},
		info:     callerInfo, // caller's TypesInfo — no entry for innerIdent
		fset:     fset,
		idx:      idx,
		root:     "/tmp",
		epIndex:  map[struct{ file string; line int }]int{},
		hmap:     map[int][]ast.Expr{},
	}
	parent := &routerInfo{prefix: "/api", middleware: []string{"auth"}}

	// Verify the precondition: calleeInfo resolves, callerInfo does not.
	if pos := resolveCallTarget(call, calleeInfo); !pos.IsValid() {
		t.Fatal("resolveCallTarget should succeed with calleeInfo")
	}
	if pos := resolveCallTarget(call, callerInfo); pos.IsValid() {
		t.Fatal("resolveCallTarget should fail with callerInfo")
	}

	// After the fix, forwardRouterCalls creates fwdCtx with info=calleeInfo,
	// so rescanCalleeWithPrefixDepth resolves the inner call via calleeInfo.
	forwardRouterCalls(stmts, "router", "/api", parent, calleeInfo, ctx, 0)
}
