//ff:func feature=scan type=extract control=selection topic=nestjs
//ff:what 파싱된 타입을 enum 또는 DTO 스키마로 등록하고 중첩 ref를 큐잉한다
package nestjs

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// registerType registers the job's type as an enum schema (and stops) or as a
// DTO schema, then enqueues the DTO's nested field references. Top-level types
// are not registered (emitted via endpoint TypeName) but their refs are followed.
func (r *schemaRegistry) registerType(j schemaJob, absPath string, root *sitter.Node, src []byte, fileImports map[string]string) {
	if vals := extractEnumMembers(root, src, j.typeName); vals != nil {
		if !r.topLevel[j.typeName] {
			r.schemas[j.typeName] = scanner.EnumSchema(vals)
		}
		return
	}
	fields, err := extractDTO(absPath, j.typeName, fileImports, j.projectRoot, r.cache)
	if err != nil || fields == nil {
		return
	}
	if !r.topLevel[j.typeName] {
		r.schemas[j.typeName] = scanner.SchemaFromFields(fields)
	}
	r.enqueueFields(fields, fileImports, absPath, j.projectRoot)
}
