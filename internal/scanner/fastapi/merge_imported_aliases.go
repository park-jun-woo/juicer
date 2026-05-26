//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what import 경로를 추적하여 cross-file 타입 별칭을 글로벌 맵에 병합한다
package fastapi

import "path/filepath"

// mergeImportedAliases resolves cross-file imports for a single file and adds
// any imported type aliases to the global map.
func mergeImportedAliases(fi fileInfo, perFile map[string]map[string]string, global map[string]string) {
	importMap := buildImportMap(fi.imports, filepath.Dir(fi.absPath))
	for _, imp := range fi.imports {
		srcFile := importMap[imp.name]
		if srcFile == "" {
			continue
		}
		if fn, ok := perFile[srcFile][imp.name]; ok {
			global[imp.name] = fn
		}
	}
}
