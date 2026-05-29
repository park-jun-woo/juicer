//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what Resource 클래스를 담은 파일을 파싱 파일/PSR-4 경로에서 찾는다
package laravel

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
