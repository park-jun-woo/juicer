//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 컨트롤러 클래스명에서 파일 경로를 추적하고 메서드를 파싱한다 (PSR-4)
package laravel

import (
	"os"
	"path/filepath"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// controllerMethod holds information about a controller method.
type controllerMethod struct {
	name           string
	params         []methodParam
	formRequestRef string           // FormRequest type hint if any
	returnNodes    []*sitter.Node   // return statement nodes
	src            []byte
}

// methodParam holds a parameter of a controller method.
type methodParam struct {
	name     string
	typeName string
}

// resolveController finds and parses a controller file from its class name.
// It searches common Laravel paths: app/Http/Controllers/**/*.php
func resolveController(absRoot, className string, parsedFiles map[string]*fileInfo) *fileInfo {
	// Try to find the file among already parsed files
	for _, fi := range parsedFiles {
		if classMatches(fi, className) {
			return fi
		}
	}

	// Try PSR-4 convention: App\Http\Controllers\UserController -> app/Http/Controllers/UserController.php
	candidates := []string{
		filepath.Join(absRoot, "app", "Http", "Controllers", className+".php"),
		filepath.Join(absRoot, "app", "Http", "Controllers", "Api", className+".php"),
		filepath.Join(absRoot, "app", "Http", "Controllers", "API", className+".php"),
	}

	for _, candidate := range candidates {
		if _, err := os.Stat(candidate); err == nil {
			fi, err := parseFile(absRoot, candidate)
			if err == nil {
				return fi
			}
		}
	}
	return nil
}

// classMatches checks whether a fileInfo contains a class declaration matching the name.
func classMatches(fi *fileInfo, className string) bool {
	classes := findAllByType(fi.root, "class_declaration")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "name")
		if nameNode != nil && nodeText(nameNode, fi.src) == className {
			return true
		}
	}
	return false
}

// extractControllerMethod finds and extracts a specific method from a controller file.
func extractControllerMethod(fi *fileInfo, methodName string) *controllerMethod {
	classes := findAllByType(fi.root, "class_declaration")
	for _, cls := range classes {
		declList := findChildByType(cls, "declaration_list")
		if declList == nil {
			continue
		}
		methods := childrenOfType(declList, "method_declaration")
		for _, method := range methods {
			mNameNode := findChildByType(method, "name")
			if mNameNode == nil || nodeText(mNameNode, fi.src) != methodName {
				continue
			}
			cm := &controllerMethod{
				name: methodName,
				src:  fi.src,
			}
			// Extract parameters
			formalParams := findChildByType(method, "formal_parameters")
			if formalParams != nil {
				cm.params = extractMethodParams(formalParams, fi.src)
				cm.formRequestRef = findFormRequestParam(cm.params)
			}
			// Extract return statements
			cm.returnNodes = findAllByType(method, "return_statement")
			return cm
		}
	}
	return nil
}

// extractMethodParams extracts parameters from a formal_parameters node.
func extractMethodParams(formalParams *sitter.Node, src []byte) []methodParam {
	var params []methodParam
	simpleParams := childrenOfType(formalParams, "simple_parameter")
	for _, sp := range simpleParams {
		mp := methodParam{}
		// Check for type hint
		typeNode := findChildByType(sp, "named_type")
		if typeNode != nil {
			nameNode := findChildByType(typeNode, "name")
			if nameNode != nil {
				mp.typeName = nodeText(nameNode, src)
			}
		}
		// Check for primitive type
		primType := findChildByType(sp, "primitive_type")
		if primType != nil {
			mp.typeName = nodeText(primType, src)
		}
		// Variable name
		varName := findChildByType(sp, "variable_name")
		if varName != nil {
			mp.name = strings.TrimPrefix(nodeText(varName, src), "$")
		}
		params = append(params, mp)
	}
	return params
}

// findFormRequestParam finds a FormRequest type-hinted parameter.
func findFormRequestParam(params []methodParam) string {
	for _, p := range params {
		if p.typeName != "" && p.typeName != "Request" &&
			p.typeName != "int" && p.typeName != "string" && p.typeName != "float" &&
			p.typeName != "bool" && p.typeName != "array" {
			// Likely a FormRequest subclass if it's not a built-in type
			if strings.HasSuffix(p.typeName, "Request") {
				return p.typeName
			}
		}
	}
	return ""
}
