//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestGinCtxParamName 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestGinCtxParamName(t *testing.T) {
	tests := []struct {
		name string
		ft   *ast.FuncType
		want string
	}{
		{
			name: "nil params",
			ft:   &ast.FuncType{Params: nil},
			want: "",
		},
		{
			name: "no params",
			ft:   &ast.FuncType{Params: &ast.FieldList{}},
			want: "",
		},
		{
			name: "with *gin.Context named c",
			ft: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{{
						Names: []*ast.Ident{{Name: "c"}},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   &ast.Ident{Name: "gin"},
								Sel: &ast.Ident{Name: "Context"},
							},
						},
					}},
				},
			},
			want: "c",
		},
		{
			name: "with *gin.Context named ctx",
			ft: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{{
						Names: []*ast.Ident{{Name: "ctx"}},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   &ast.Ident{Name: "gin"},
								Sel: &ast.Ident{Name: "Context"},
							},
						},
					}},
				},
			},
			want: "ctx",
		},
		{
			name: "unnamed *gin.Context",
			ft: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{{
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   &ast.Ident{Name: "gin"},
								Sel: &ast.Ident{Name: "Context"},
							},
						},
					}},
				},
			},
			want: "_",
		},
		{
			name: "non-gin param",
			ft: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{{
						Names: []*ast.Ident{{Name: "x"}},
						Type:  &ast.Ident{Name: "int"},
					}},
				},
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ginCtxParamName(tt.ft)
			if got != tt.want {
				t.Errorf("ginCtxParamName() = %q, want %q", got, tt.want)
			}
		})
	}
}
