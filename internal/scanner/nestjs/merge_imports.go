//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 두 import 맵을 병합한다
package nestjs

// mergeImports merges two import maps. File-level imports take precedence
// when the same name exists in both (closer scope wins).
func mergeImports(caller, file map[string]string) map[string]string {
	result := make(map[string]string, len(caller)+len(file))
	for k, v := range caller {
		result[k] = v
	}
	for k, v := range file {
		result[k] = v
	}
	return result
}
