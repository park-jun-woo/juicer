//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 파일의 Pass 1: 파싱, 라우터 수집, import 해석, use-prefix 수집
package express

func scanOneFilePass1(path string, parsed map[string]*fileInfo, allRouters map[string]map[string]bool, absRoot string, aliases map[string]string) []mountEntry {
	fi, err := parseFile(path)
	if err != nil {
		return nil
	}
	parsed[path] = fi
	routers := collectRouters(fi)
	allRouters[path] = routers
	imports := resolveImports(fi, absRoot, aliases)
	mounts := resolveUsePrefixes(fi, routers, imports)
	var entries []mountEntry
	for _, m := range mounts {
		entries = append(entries, mountEntry{
			prefix:     m.Prefix,
			varName:    m.VarName,
			filePath:   m.FilePath,
			sourceFile: path,
		})
	}
	entries = append(entries, extractArrayRouteMounts(fi, routers, imports, path)...)
	return entries
}
