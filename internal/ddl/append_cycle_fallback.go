//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 미출력(순환) 노드를 알파벳순으로 결과에 추가하고 경고 로그 출력
package ddl

import "log"

// appendCycleFallback appends any unemitted (cyclic) nodes in alphabetical
// order and logs a warning naming them.
func appendCycleFallback(names []string, emitted map[string]bool, result []string) []string {
	var cyclic []string
	for _, name := range names {
		if !emitted[name] {
			cyclic = append(cyclic, name)
		}
	}
	if len(cyclic) > 0 {
		log.Printf("ddl: foreign-key cycle detected; emitting %v in alphabetical order (FK may fail at apply time)", cyclic)
		result = append(result, cyclic...)
	}
	return result
}
