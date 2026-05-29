//ff:func feature=scan type=parse control=iteration dimension=1 topic=laravel
//ff:what 모든 PHP 파일을 파싱해 relPath → fileInfo 맵을 만든다
package laravel

func parseAllPHPFiles(absRoot string, phpFiles []string) map[string]*fileInfo {
	parsedFiles := make(map[string]*fileInfo)
	for _, f := range phpFiles {
		fi, err := parseFile(absRoot, f)
		if err != nil {
			continue
		}
		parsedFiles[fi.relPath] = fi
	}
	return parsedFiles
}
