//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractValidateRequest_IdentAndInline 테스트
package express

import "testing"

func TestExtractValidateRequest_IdentAndInline(t *testing.T) {
	fi := mustParse(t, []byte(`validateRequest({ body: createSchema, query: z.object({}), unknown: x });`))
	got := extractValidateRequest(firstCallExpr(t, fi), fi.Src)
	if len(got) != 2 {
		t.Fatalf("expected 2 validators (unknown key skipped), got %+v", got)
	}

	if got[0].Target != "json" || got[0].SchemaName != "createSchema" {
		t.Fatalf("unexpected body validator %+v", got[0])
	}

	if got[1].Target != "query" || got[1].SchemaNode == nil {
		t.Fatalf("unexpected query validator %+v", got[1])
	}
}
