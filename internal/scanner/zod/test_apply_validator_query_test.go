//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestApplyValidator_Query 테스트
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyValidator_Query(t *testing.T) {
	root, src := parseTS(t, `const qSchema = z.object({ limit: z.number() });`)
	schemas := CollectSchemas(root, src)
	req := &scanner.Request{}
	v := ValidatorInfo{Target: "query", SchemaName: "qSchema"}
	ok := ApplyValidator(req, v, schemas, src, map[string][]byte{"qSchema": src})
	if !ok || len(req.Query) != 1 {
		t.Fatalf("got ok=%v query=%+v", ok, req.Query)
	}
}
