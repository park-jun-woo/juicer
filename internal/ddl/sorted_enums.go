//ff:func feature=ddl type=render control=sequence
//ff:what enum 타입 슬라이스를 이름순으로 정렬한 복사본 반환
package ddl

import "sort"

// sortedEnums returns a copy of enums sorted by name.
func sortedEnums(enums []EnumType) []EnumType {
	out := make([]EnumType, len(enums))
	copy(out, enums)
	sort.Slice(out, func(i, j int) bool {
		return out[i].Name < out[j].Name
	})
	return out
}
