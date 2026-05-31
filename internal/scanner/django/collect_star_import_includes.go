//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 패키지 __init__.py의 `from .X import urlpatterns` 집계를 prefix 없는 include로 변환한다
package django

import "path/filepath"

// collectStarImportIncludes detects a package __init__.py that aggregates
// submodule urlpatterns via `from .X import urlpatterns [as Y]` and turns each
// such submodule into a synthetic prefix-less include entry on the package
// module. This lets include('pkg') reach the submodules' path() entries while
// keeping the include() prefix intact, and marks the submodules as included so
// findRootURLModules no longer mistakes them for roots.
func collectStarImportIncludes(fi fileInfo) []urlEntry {
	if filepath.Base(filepath.ToSlash(fi.relPath)) != "__init__.py" {
		return nil
	}
	var entries []urlEntry
	for _, stmt := range findAllByType(fi.root, "import_from_statement") {
		if !importsName(stmt, "urlpatterns", fi.src) {
			continue
		}
		sub := relativeImportSubmodule(stmt, fi.src)
		if sub == "" {
			continue
		}
		entries = append(entries, urlEntry{
			isInclude:     true,
			includeModule: fi.module + "." + sub,
		})
	}
	return entries
}
