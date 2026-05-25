//ff:func feature=scan type=extract control=sequence
//ff:what TestBuildField_Excluded 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestBuildField_Excluded(t *testing.T) {
	v := types.NewVar(0, nil, "Secret", types.Typ[types.String])
	f := buildField(v, `json:"-"`, make(map[string]bool))
	if f != nil {
		t.Fatal("expected nil for json:-")
	}
}
