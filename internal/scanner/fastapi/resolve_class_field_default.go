//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 클래스 필드의 기본 문자열 값을 추출한다
package fastapi

// findClassFieldDefault searches all files for a class named className, then
// looks up the default string literal value for attrName in its body.
// Supports patterns like `ATTR_NAME: type = "value"`.
func findClassFieldDefault(files []fileInfo, className, attrName string) string {
	for _, fi := range files {
		result := findClassFieldDefaultInFile(fi, className, attrName)
		if result != "" {
			return result
		}
	}
	return ""
}
