//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what import된 이름에 대응하는 원본 파일의 라우터 변수명을 찾는다
package fastapi

// findOriginalVarName finds the original variable name in the source file for
// an imported name. Returns "" if the imported name does not match any router
// variable in the source file.
func findOriginalVarName(importedName string, srcFI *fileInfo) string {
	if _, ok := srcFI.prefixes[importedName]; ok {
		return importedName
	}
	return ""
}
