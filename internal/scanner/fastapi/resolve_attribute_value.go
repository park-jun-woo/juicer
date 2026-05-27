//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what dotted identifier에서 객체 속성 값을 해석한다
package fastapi

import "strings"

// resolveAttributeValue parses "objName.attrName", finds the class that objName
// is instantiated from (e.g., settings = Settings()), then looks up the default
// value of attrName in that class definition across all files.
func resolveAttributeValue(absRoot string, files []fileInfo, objAttr string) string {
	parts := strings.SplitN(objAttr, ".", 2)
	if len(parts) != 2 {
		return ""
	}
	objName, attrName := parts[0], parts[1]

	for _, fi := range files {
		// 1. Find objName = ClassName() pattern in this file
		className := findAssignedClassName(fi.root, objName, fi.src)
		if className == "" {
			continue
		}
		// 2. Look up attrName default in className's class definition
		return findClassFieldDefault(files, className, attrName)
	}
	return ""
}
