//ff:func feature=scan type=parse control=sequence topic=zod
//ff:what 문자열을 int로 파싱하여 포인터를 반환한다
package zod

import "strconv"

func parseIntArg(s string) *int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &n
}
