//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what Pydantic BaseModel 서브클래스를 파일에서 읽어 필드를 추출한다
package fastapi

import (
	"os"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// extractPydanticModel parses a source file and extracts fields from a Pydantic
// BaseModel subclass with the given name.
func extractPydanticModel(filePath, className string) ([]scanner.Field, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return extractPydanticModelFromSource(src, className)
}
