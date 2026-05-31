//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 단일 작업의 DTO 파일을 열어 enum/DTO 스키마를 등록하고 중첩 ref를 큐잉한다
package nestjs

import "os"

// process resolves one schemaJob: it opens the referenced type's source file,
// registers the type as an enum or DTO schema (unless it is a top-level type),
// and enqueues any nested field references for further resolution.
func (r *schemaRegistry) process(j schemaJob) {
	if r.processed[j.typeName] {
		return
	}
	r.processed[j.typeName] = true

	absPath := resolveJobFile(j)
	if absPath == "" {
		return
	}
	src, err := os.ReadFile(absPath)
	if err != nil {
		return
	}
	root, err := parseTypeScript(src)
	if err != nil {
		return
	}
	fileImports := mergeImports(j.imports, extractImports(root, src))
	r.registerType(j, absPath, root, src, fileImports)
}
