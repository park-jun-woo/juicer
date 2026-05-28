//ff:func feature=scan type=parse control=iteration dimension=1 topic=django
//ff:what Python 파일 목록을 모두 파싱하여 fileInfo 목록을 반환한다
package django

// parseAllFiles parses all Python files and returns fileInfo list.
func parseAllFiles(absRoot string, pyFiles []string) []fileInfo {
	var files []fileInfo
	for _, f := range pyFiles {
		fi, err := parseFile(absRoot, f)
		if err != nil {
			continue
		}
		files = append(files, *fi)
	}
	return files
}
