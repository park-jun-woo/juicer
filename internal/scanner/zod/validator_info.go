//ff:type feature=scan type=model topic=zod
//ff:what Zod validator 추출 결과 구조체
package zod

import sitter "github.com/smacker/go-tree-sitter"

// ValidatorInfo — zValidator/validateRequest 추출 결과
type ValidatorInfo struct {
	Target     string       // "json", "query", "param", "form"
	SchemaName string       // 변수명 (e.g. "createUserSchema")
	SchemaNode *sitter.Node // 인라인 스키마 노드
}
