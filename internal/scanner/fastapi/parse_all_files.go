//ff:func feature=scan type=parse control=iteration dimension=1 topic=fastapi
//ff:what Python 파일 목록을 모두 파싱하고 크로스파일 prefix를 병합한다
package fastapi

// parseAllFiles parses all Python files and returns fileInfo list.
// After the initial per-file parse, it runs a cross-file prefix merge pass
// to resolve include_router chains that span multiple files.
func parseAllFiles(absRoot string, pyFiles []string) []fileInfo {
	var files []fileInfo
	for _, f := range pyFiles {
		fi, err := parseFile(absRoot, f)
		if err != nil {
			continue
		}
		files = append(files, *fi)
	}
	mergeCrossFilePrefixes(absRoot, files)
	return files
}
