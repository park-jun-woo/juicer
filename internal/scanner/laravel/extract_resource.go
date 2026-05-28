//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what JsonResource::toArray() 에서 응답 필드 키를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractResourceFields finds a JsonResource class and extracts field names from toArray().
func extractResourceFields(absRoot, className string, parsedFiles map[string]*fileInfo) []scanner.Field {
	fi := findResourceFile(absRoot, className, parsedFiles)
	if fi == nil {
		return nil
	}
	return extractToArrayFields(fi, className)
}

// findResourceFile locates the file containing the Resource class.
func findResourceFile(absRoot, className string, parsedFiles map[string]*fileInfo) *fileInfo {
	for _, fi := range parsedFiles {
		if classMatches(fi, className) {
			return fi
		}
	}
	candidates := []string{
		absRoot + "/app/Http/Resources/" + className + ".php",
	}
	for _, candidate := range candidates {
		fi, err := parseFile(absRoot, candidate)
		if err == nil {
			return fi
		}
	}
	return nil
}

// extractToArrayFields extracts field names from the toArray method's return array.
func extractToArrayFields(fi *fileInfo, className string) []scanner.Field {
	classes := findAllByType(fi.root, "class_declaration")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "name")
		if nameNode != nil && nodeText(nameNode, fi.src) != className {
			continue
		}
		declList := findChildByType(cls, "declaration_list")
		if declList == nil {
			continue
		}
		methods := childrenOfType(declList, "method_declaration")
		for _, method := range methods {
			mName := findChildByType(method, "name")
			if mName == nil || nodeText(mName, fi.src) != "toArray" {
				continue
			}
			return extractArrayKeys(method, fi.src)
		}
	}
	return nil
}

// extractArrayKeys extracts keys from the return array of toArray().
func extractArrayKeys(method *sitter.Node, src []byte) []scanner.Field {
	retStmts := findAllByType(method, "return_statement")
	if len(retStmts) == 0 {
		return nil
	}
	arrNodes := findAllByType(retStmts[0], "array_creation_expression")
	if len(arrNodes) == 0 {
		return nil
	}
	arr := arrNodes[0]
	elems := childrenOfType(arr, "array_element_initializer")
	var fields []scanner.Field
	for _, elem := range elems {
		keyNode := findChildByType(elem, "string")
		if keyNode == nil {
			continue
		}
		key := extractStringContent(keyNode, src)
		if key == "" {
			continue
		}
		fields = append(fields, scanner.Field{
			Name: key,
			JSON: key,
			Type: "string", // default since we can't infer type from $this->field
		})
	}
	return fields
}
