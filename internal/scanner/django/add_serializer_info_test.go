//ff:func feature=scan type=test control=sequence topic=django
//ff:what addSerializerInfo — 쓰기 메서드 serializer body 추가 분기를 검증
package django

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAddSerializerInfo_Adds(t *testing.T) {
	ep := &scanner.Endpoint{}
	am := actionMethod{action: "create", method: "POST"}
	serializers := map[string]serializerInfo{
		"UserSerializer": {name: "User", fields: []scanner.Field{{Name: "id"}}},
	}
	addSerializerInfo(ep, am, "UserSerializer", serializers)
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatal("expected body added")
	}
	if ep.Request.Body.TypeName != "User" {
		t.Errorf("TypeName = %q, want User", ep.Request.Body.TypeName)
	}
}

func TestAddSerializerInfo_NotWriteMethod(t *testing.T) {
	ep := &scanner.Endpoint{}
	am := actionMethod{action: "list", method: "GET"}
	addSerializerInfo(ep, am, "UserSerializer", map[string]serializerInfo{
		"UserSerializer": {name: "User"},
	})
	if ep.Request != nil {
		t.Fatal("expected no body for read method")
	}
}

func TestAddSerializerInfo_EmptyClass(t *testing.T) {
	ep := &scanner.Endpoint{}
	addSerializerInfo(ep, actionMethod{method: "POST"}, "", nil)
	if ep.Request != nil {
		t.Fatal("expected no body for empty serializer class")
	}
}

func TestAddSerializerInfo_NotFound(t *testing.T) {
	ep := &scanner.Endpoint{}
	am := actionMethod{method: "POST"}
	addSerializerInfo(ep, am, "Missing", map[string]serializerInfo{})
	if ep.Request != nil {
		t.Fatal("expected no body when serializer not in map")
	}
}

func TestAddSerializerInfo_ExistingRequest(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{}}
	am := actionMethod{method: "PUT"}
	addSerializerInfo(ep, am, "S", map[string]serializerInfo{"S": {name: "T"}})
	if ep.Request.Body == nil || ep.Request.Body.TypeName != "T" {
		t.Fatalf("expected body set on existing request, got %+v", ep.Request.Body)
	}
}
