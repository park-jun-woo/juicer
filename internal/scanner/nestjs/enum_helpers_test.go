//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what collectEnumElements / collectEnumMemberValues / extractEnumArray / extractEnumAssignmentValue 테스트
package nestjs

import "testing"

func TestCollectEnumElements(t *testing.T) {
	src := []byte(`const x = ['a', 'b', 1];`)
	root, _ := parseTypeScript(src)
	arr := findAllByType(root, "array")[0]
	vals := collectEnumElements(arr, src)
	if len(vals) != 3 || vals[0] != "a" || vals[2] != "1" {
		t.Fatalf("got %v", vals)
	}
}

func TestCollectEnumMemberValues_Assignment(t *testing.T) {
	src := []byte(`enum Status { OPEN = 'open', CLOSED = 'closed' }`)
	root, _ := parseTypeScript(src)
	bodies := findAllByType(root, "enum_body")
	if len(bodies) == 0 {
		t.Fatal("no enum_body")
	}
	vals := collectEnumMemberValues(bodies[0], src)
	if len(vals) != 2 || vals[0] != "open" {
		t.Fatalf("got %v", vals)
	}
}

func TestCollectEnumMemberValues_Valueless(t *testing.T) {
	src := []byte(`enum Dir { Up, Down }`)
	root, _ := parseTypeScript(src)
	bodies := findAllByType(root, "enum_body")
	vals := collectEnumMemberValues(bodies[0], src)
	if len(vals) != 2 || vals[0] != "Up" {
		t.Fatalf("got %v", vals)
	}
}

func TestExtractEnumArray(t *testing.T) {
	src := []byte(`const o = { enum: ['a', 'b'] };`)
	root, _ := parseTypeScript(src)
	obj := findAllByType(root, "object")[0]
	vals := extractEnumArray(obj, src)
	if len(vals) != 2 {
		t.Fatalf("got %v", vals)
	}
}

func TestExtractEnumArray_NoEnumKey(t *testing.T) {
	src := []byte(`const o = { other: ['a'] };`)
	root, _ := parseTypeScript(src)
	obj := findAllByType(root, "object")[0]
	if vals := extractEnumArray(obj, src); vals != nil {
		t.Fatalf("expected nil, got %v", vals)
	}
}

func TestExtractEnumAssignmentValue_PropertyIdentifier(t *testing.T) {
	src := []byte(`enum E { Foo }`)
	root, _ := parseTypeScript(src)
	body := findAllByType(root, "enum_body")[0]
	// find a property_identifier inside
	pids := findAllByType(body, "property_identifier")
	if len(pids) == 0 {
		t.Skip("no property identifier")
	}
	v, ok := extractEnumAssignmentValue(pids[0], src)
	if !ok || v != "Foo" {
		t.Fatalf("got %q %v", v, ok)
	}
}

func TestExtractEnumAssignmentValue_StringValue(t *testing.T) {
	src := []byte(`enum E { OPEN = 'open' }`)
	root, _ := parseTypeScript(src)
	asn := findAllByType(root, "enum_assignment")
	if len(asn) == 0 {
		t.Skip("no enum_assignment")
	}
	v, ok := extractEnumAssignmentValue(asn[0], src)
	if !ok || v != "open" {
		t.Fatalf("got %q %v", v, ok)
	}
}

func TestExtractEnumAssignmentValue_NumberValueUsesName(t *testing.T) {
	src := []byte(`enum E { FIRST = 1 }`)
	root, _ := parseTypeScript(src)
	asn := findAllByType(root, "enum_assignment")
	if len(asn) == 0 {
		t.Skip("no enum_assignment")
	}
	v, ok := extractEnumAssignmentValue(asn[0], src)
	if !ok || v != "FIRST" {
		t.Fatalf("got %q %v", v, ok)
	}
}

func TestExtractEnumAssignmentValue_Irrelevant(t *testing.T) {
	src := []byte(`const x = 1;`)
	root, _ := parseTypeScript(src)
	nums := findAllByType(root, "number")
	if len(nums) == 0 {
		t.Skip("no number")
	}
	if _, ok := extractEnumAssignmentValue(nums[0], src); ok {
		t.Fatal("expected false for irrelevant node")
	}
}
