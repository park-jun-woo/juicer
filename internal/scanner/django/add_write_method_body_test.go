//ff:func feature=scan type=test control=sequence topic=django
//ff:what addWriteMethodBody — APIView 쓰기 메서드 serializer body 추가 분기를 검증
package django

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAddWriteMethodBody_Adds(t *testing.T) {
	ep := &scanner.Endpoint{}
	serializers := map[string]serializerInfo{
		"S": {name: "User", fields: []scanner.Field{{Name: "id"}}},
	}
	addWriteMethodBody(ep, "POST", "S", serializers)
	if ep.Request == nil || ep.Request.Body == nil || ep.Request.Body.TypeName != "User" {
		t.Fatalf("expected body added, got %+v", ep.Request)
	}
}

func TestAddWriteMethodBody_NotWrite(t *testing.T) {
	ep := &scanner.Endpoint{}
	addWriteMethodBody(ep, "GET", "S", map[string]serializerInfo{"S": {name: "User"}})
	if ep.Request != nil {
		t.Fatal("expected no body for GET")
	}
}

func TestAddWriteMethodBody_EmptyClass(t *testing.T) {
	ep := &scanner.Endpoint{}
	addWriteMethodBody(ep, "POST", "", nil)
	if ep.Request != nil {
		t.Fatal("expected no body for empty class")
	}
}

func TestAddWriteMethodBody_NotFound(t *testing.T) {
	ep := &scanner.Endpoint{}
	addWriteMethodBody(ep, "PUT", "Missing", map[string]serializerInfo{})
	if ep.Request != nil {
		t.Fatal("expected no body when serializer missing")
	}
}

func TestAddWriteMethodBody_ExistingRequest(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{}}
	addWriteMethodBody(ep, "PATCH", "S", map[string]serializerInfo{"S": {name: "T"}})
	if ep.Request.Body == nil || ep.Request.Body.TypeName != "T" {
		t.Fatalf("expected body set, got %+v", ep.Request.Body)
	}
}
