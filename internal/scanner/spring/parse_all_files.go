//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 모든 Java 파일을 파싱한다
package spring

func parseAllFiles(absRoot string, paths []string) []*fileInfo {
	var files []*fileInfo
	for _, p := range paths {
		fi, err := parseFile(absRoot, p)
		if err != nil {
			continue
		}
		files = append(files, fi)
	}
	return files
}
