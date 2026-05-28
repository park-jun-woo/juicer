//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 인터페이스 이름으로 Java 소스 파일 경로를 찾는다
package spring

func findInterfaceFile(ifaceName string, imports map[string]string, referrerPath, projectRoot string) string {
	// 1. same file check
	filePath := resolveSameFileInterface(referrerPath, ifaceName)
	if filePath != "" {
		return filePath
	}
	// 2. imports map
	if fqcn, ok := imports[ifaceName]; ok {
		filePath = resolveImportPath(projectRoot, fqcn)
		if filePath != "" {
			return filePath
		}
	}
	// 3. same package
	filePath = resolveSamePackageClass(referrerPath, ifaceName)
	return filePath
}
