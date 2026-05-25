//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what 슬라이스에 중복 없이 값 추가
package sqls

// appendUnique appends val to slice if not already present.
//
func appendUnique(slice []string, val string) []string {
	for _, s := range slice {
		if s == val {
			return slice
		}
	}
	return append(slice, val)
}

