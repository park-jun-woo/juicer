//ff:func feature=scan type=test control=sequence
//ff:what TestGoTypeFormat_Int64 테스트
package scanner

import "testing"

func TestGoTypeFormat_Int64(t *testing.T) {
	if goTypeFormat("int64", Field{}) != "int64" {
		t.Fatal("int64")
	}
	if goTypeFormat("int32", Field{}) != "int32" {
		t.Fatal("int32")
	}
	if goTypeFormat("float64", Field{}) != "double" {
		t.Fatal("double")
	}
	if goTypeFormat("float32", Field{}) != "float" {
		t.Fatal("float")
	}
	if goTypeFormat("time.Time", Field{}) != "date-time" {
		t.Fatal("date-time")
	}
	if goTypeFormat("string", Field{Validate: "required,email"}) != "email" {
		t.Fatal("email")
	}
	if goTypeFormat("string", Field{Validate: "url"}) != "uri" {
		t.Fatal("uri")
	}
	if goTypeFormat("string", Field{}) != "" {
		t.Fatal("expected empty")
	}
}

