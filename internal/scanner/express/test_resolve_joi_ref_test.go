//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveJoiRef import 매핑→대상 파일 const object를 Joi 요청 스키마로 해석 테스트
package express

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveJoiRef(t *testing.T) {
	dir := t.TempDir()
	// schema file with a const object
	schemaPath := filepath.Join(dir, "schemas.ts")
	if err := os.WriteFile(schemaPath, []byte(
		`const userBody = { body: Joi.object({ name: Joi.string() }) };
module.exports = { userBody };`), 0o644); err != nil {
		t.Fatal(err)
	}
	// route file importing the schema module
	routePath := filepath.Join(dir, "routes.ts")
	if err := os.WriteFile(routePath, []byte(
		`const schemas = require('./schemas');
router.post('/u', validate(schemas.userBody), h);`), 0o644); err != nil {
		t.Fatal(err)
	}
	routeFi, err := parseFile(routePath)
	if err != nil {
		t.Fatal(err)
	}
	ctx := &scanContext{parsed: map[string]*fileInfo{}, absRoot: dir, pathAliases: map[string]string{}}

	rs := resolveJoiRef(joiValidatorRef{ImportName: "schemas", Member: "userBody"}, routeFi, ctx)
	if rs.Empty() {
		t.Fatalf("expected resolved schema, got empty")
	}

	// unknown import -> empty
	if rs2 := resolveJoiRef(joiValidatorRef{ImportName: "unknown", Member: "x"}, routeFi, ctx); !rs2.Empty() {
		t.Errorf("unknown import should be empty: %+v", rs2)
	}
}
