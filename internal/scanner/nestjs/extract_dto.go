//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what DTO 소스 파일에서 지정 클래스의 프로퍼티를 추출한다
package nestjs

import (
	"os"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// extractDTO parses a DTO source file and extracts the class with the given name.
func extractDTO(filePath, className string) ([]scanner.Field, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	root, err := parseTypeScript(src)
	if err != nil {
		return nil, err
	}
	classes := findAllByType(root, "class_declaration")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "type_identifier")
		if nameNode == nil {
			continue
		}
		if nodeText(nameNode, src) != className {
			continue
		}
		fields := extractClassProperties(cls, src)
		return dtoFieldsToScannerFields(fields), nil
	}
	return nil, nil
}
