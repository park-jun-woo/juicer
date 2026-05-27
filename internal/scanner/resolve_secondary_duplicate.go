//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what seen 맵에서 중복 키를 숫자 접미사로 해소한다
package scanner

import "strconv"

func resolveSecondaryDuplicate(prefixed string, seen map[string]bool) string {
	if !seen[prefixed] {
		return prefixed
	}
	for n := 2; ; n++ {
		candidate := prefixed + strconv.Itoa(n)
		if !seen[candidate] {
			return candidate
		}
	}
}
