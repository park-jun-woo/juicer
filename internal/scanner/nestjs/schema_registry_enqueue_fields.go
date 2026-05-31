//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 필드 목록의 named ref들을 해석 큐에 추가한다
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// enqueueFields appends a resolution job for each field that references a named
// type, using the given imports/referrer so nested lookups use the right file.
func (r *schemaRegistry) enqueueFields(fields []scanner.Field, imports map[string]string, referrer, projectRoot string) {
	for _, f := range fields {
		if f.Ref != "" {
			r.queue = append(r.queue, schemaJob{f.Ref, imports, referrer, projectRoot})
		}
	}
}
