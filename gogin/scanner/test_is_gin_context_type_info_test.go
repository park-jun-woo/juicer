//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinContextTypeInfo 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestIsGinContextTypeInfo(t *testing.T) {
	// Not a pointer
	basic := types.Typ[types.Int]
	if isGinContextTypeInfo(basic) {
		t.Error("expected false for non-pointer type")
	}

	// Pointer to non-named type
	ptrToBasic := types.NewPointer(types.Typ[types.String])
	if isGinContextTypeInfo(ptrToBasic) {
		t.Error("expected false for pointer to basic type")
	}

	// Pointer to named type that's not gin.Context
	pkg := types.NewPackage("example.com/test", "test")
	named := types.NewNamed(types.NewTypeName(0, pkg, "MyType", nil), types.NewStruct(nil, nil), nil)
	ptrToNamed := types.NewPointer(named)
	if isGinContextTypeInfo(ptrToNamed) {
		t.Error("expected false for non-gin named type")
	}

	// Pointer to a named type called "Context" but wrong package
	wrongPkg := types.NewPackage("example.com/wrong", "wrong")
	ctxNamed := types.NewNamed(types.NewTypeName(0, wrongPkg, "Context", nil), types.NewStruct(nil, nil), nil)
	ptrToCtx := types.NewPointer(ctxNamed)
	if isGinContextTypeInfo(ptrToCtx) {
		t.Error("expected false for Context from wrong package")
	}

	// Pointer to a named type called "Context" in gin-gonic/gin package
	ginPkg := types.NewPackage("github.com/gin-gonic/gin", "gin")
	ginCtxNamed := types.NewNamed(types.NewTypeName(0, ginPkg, "Context", nil), types.NewStruct(nil, nil), nil)
	ptrToGinCtx := types.NewPointer(ginCtxNamed)
	if !isGinContextTypeInfo(ptrToGinCtx) {
		t.Error("expected true for *gin.Context")
	}

	// Named type with nil package
	noPkgNamed := types.NewNamed(types.NewTypeName(0, nil, "Context", nil), types.NewStruct(nil, nil), nil)
	ptrToNoPkg := types.NewPointer(noPkgNamed)
	if isGinContextTypeInfo(ptrToNoPkg) {
		t.Error("expected false for Context with nil package")
	}
}
