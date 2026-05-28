//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what DTO 타입을 해석하여 필드를 반환한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func resolveDTOType(dr dtoRequest, files []*fileInfo, cache map[string][]scanner.Field) []scanner.Field {
	if cached, ok := cache[dr.typeName]; ok {
		return cached
	}

	lookupName := stripGeneric(dr.typeName)

	for _, fi := range files {
		fields := findClassInFile(fi, lookupName)
		if fields != nil {
			cache[dr.typeName] = fields
			return fields
		}
	}
	return nil
}
