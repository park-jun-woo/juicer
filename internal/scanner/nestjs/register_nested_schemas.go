//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 중첩 DTO/enum named 타입을 재귀적으로 컴포넌트 스키마로 등록한다
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// registerNestedSchemas walks the named-type references ($ref) reachable from
// the given DTO requests and registers each referenced DTO/enum as a separate
// component schema in `schemas`. It resolves nested types recursively: a DTO's
// own field refs are followed using that DTO file's imports. Already-processed
// types are skipped (cycle-safe). Top-level request/response types are not
// registered here (they are emitted via the endpoint TypeName) but their nested
// refs are still followed.
func registerNestedSchemas(dtoReqs []dtoRequest, cache map[string][]scanner.Field, schemas map[string]any) {
	r := &schemaRegistry{
		cache:     cache,
		schemas:   schemas,
		processed: make(map[string]bool),
		topLevel:  make(map[string]bool),
	}
	for _, dr := range dtoReqs {
		r.queue = append(r.queue, schemaJob{dr.typeName, dr.imports, dr.referrer, dr.projectRoot})
		r.topLevel[dr.typeName] = true
	}
	for len(r.queue) > 0 {
		j := r.queue[0]
		r.queue = r.queue[1:]
		r.process(j)
	}
}
