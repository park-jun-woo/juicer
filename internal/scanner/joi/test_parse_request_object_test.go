//ff:func feature=scan type=test topic=joi control=sequence
//ff:what ParseRequestObject body/query/params 키 → RequestSchema 변환 테스트
package joi

import "testing"

func TestParseRequestObject(t *testing.T) {
	src := `const reqSchema = {
  body: Joi.object().keys({ name: Joi.string().required() }),
  query: Joi.object().keys({ page: Joi.number() }),
  params: Joi.object().keys({ id: Joi.string() })
};`
	root, b := parseJoiTS(t, src)
	obj := firstOfType(root, "object")
	rs := ParseRequestObject(obj, b)
	if len(rs.Body) != 1 || rs.Body[0].Name != "name" {
		t.Errorf("body: %+v", rs.Body)
	}
	if len(rs.Query) != 1 || rs.Query[0].Name != "page" {
		t.Errorf("query: %+v", rs.Query)
	}
	if len(rs.Params) != 1 || rs.Params[0].Name != "id" {
		t.Errorf("params: %+v", rs.Params)
	}

	// non-object node -> empty
	if !ParseRequestObject(nil, b).Empty() {
		t.Error("nil node should be empty")
	}
}
