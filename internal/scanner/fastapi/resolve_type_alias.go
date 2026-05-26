//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 전체 파일에서 Annotated[T, Depends(...)] 타입 별칭 맵을 구축한다
package fastapi

// resolveTypeAliases builds a global map from type alias names to their Depends
// function names. It scans all parsed files for assignments like:
//
//	SessionDep = Annotated[Session, Depends(get_db)]
//
// and resolves cross-file imports so that a type alias defined in deps.py and
// imported in routes.py is recognized everywhere.
func resolveTypeAliases(files []fileInfo) map[string]string {
	// Pass 1: collect per-file aliases from assignment nodes.
	perFile := make(map[string]map[string]string)
	for _, fi := range files {
		local := collectFileAliases(fi.root, fi.src)
		if len(local) > 0 {
			perFile[fi.absPath] = local
		}
	}

	// Pass 2: merge into global map, resolving cross-file imports.
	global := make(map[string]string)
	for _, fi := range files {
		for alias, fn := range perFile[fi.absPath] {
			global[alias] = fn
		}
		mergeImportedAliases(fi, perFile, global)
	}
	return global
}
