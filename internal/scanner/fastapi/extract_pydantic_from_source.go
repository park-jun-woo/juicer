//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 소스 바이트에서 Pydantic 모델 필드를 추출한다
package fastapi

import "github.com/park-jun-woo/juicer/internal/scanner"

// extractPydanticModelFromSource extracts Pydantic model fields from source bytes.
func extractPydanticModelFromSource(src []byte, className string) ([]scanner.Field, error) {
	root, err := parsePython(src)
	if err != nil {
		return nil, err
	}
	return findPydanticClass(root, src, className), nil
}
