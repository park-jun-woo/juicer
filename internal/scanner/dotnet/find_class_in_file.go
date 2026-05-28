//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 파일에서 지정된 이름의 클래스를 찾아 프로퍼티를 추출한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func findClassInFile(fi *fileInfo, className string) []scanner.Field {
	classes := findAllByType(fi.root, "class_declaration")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "identifier")
		if nameNode == nil {
			continue
		}
		if nodeText(nameNode, fi.src) == className {
			return extractClassProps(cls, fi.src)
		}
	}
	records := findAllByType(fi.root, "record_declaration")
	for _, rec := range records {
		nameNode := findChildByType(rec, "identifier")
		if nameNode == nil {
			continue
		}
		if nodeText(nameNode, fi.src) == className {
			return extractRecordParams(rec, fi.src)
		}
	}
	return nil
}
