//ff:func feature=scan type=test control=sequence
//ff:what TestFormatType_Basic 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestFormatType_Basic(t *testing.T) {
	if formatType(types.Typ[types.String]) != "string" {
		t.Fatal("basic")
	}
	if formatType(types.NewPointer(types.Typ[types.Int])) != "*int" {
		t.Fatal("pointer")
	}
	if formatType(types.NewSlice(types.Typ[types.String])) != "[]string" {
		t.Fatal("slice")
	}
	if formatType(types.NewArray(types.Typ[types.Int], 5)) != "[]int" {
		t.Fatal("array")
	}
	if formatType(types.NewMap(types.Typ[types.String], types.Typ[types.Int])) != "map[string]int" {
		t.Fatal("map")
	}
	named := types.NewNamed(types.NewTypeName(0, types.NewPackage("time", "time"), "Time", nil), types.Typ[types.String], nil)
	if formatType(named) != "time.Time" {
		t.Fatal("named")
	}
	namedNoPkg := types.NewNamed(types.NewTypeName(0, nil, "Error", nil), types.Typ[types.String], nil)
	if formatType(namedNoPkg) != "Error" {
		t.Fatal("named no pkg")
	}
	if formatType(types.NewInterfaceType(nil, nil)) != "any" {
		t.Fatal("interface")
	}
	if formatType(types.NewStruct(nil, nil)) != "object" {
		t.Fatal("struct")
	}
}

