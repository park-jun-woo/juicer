//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestFormatType 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestFormatType(t *testing.T) {
	tests := []struct {
		name string
		typ  types.Type
		want string
	}{
		{"basic int", types.Typ[types.Int], "int"},
		{"basic string", types.Typ[types.String], "string"},
		{"pointer", types.NewPointer(types.Typ[types.Int]), "*int"},
		{"slice", types.NewSlice(types.Typ[types.String]), "[]string"},
		{"array", types.NewArray(types.Typ[types.Int], 5), "[]int"},
		{"map", types.NewMap(types.Typ[types.String], types.Typ[types.Int]), "map[string]int"},
		{"interface", types.NewInterfaceType(nil, nil), "any"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatType(tt.typ)
			if got != tt.want {
				t.Errorf("formatType() = %q, want %q", got, tt.want)
			}
		})
	}
}
