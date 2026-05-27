//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 문자열 인자를 정수로 변환하여 길이 제약 포인터에 설정한다
package nestjs

import "strconv"

// applyLengthConstraint parses an integer arg and sets the target pointer.
func applyLengthConstraint(arg string, target **int) {
	v, err := strconv.Atoi(arg)
	if err != nil {
		return
	}
	*target = &v
}
