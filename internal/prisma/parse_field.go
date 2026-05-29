//ff:func feature=prisma type=parse control=sequence topic=prisma
//ff:what 모델 본문 한 라인을 (이름, 타입, nullable, 배열, 속성들) field로 토큰화
package prisma

import "strings"

// parseField tokenizes one model body field line into a field value.
func parseField(line string) field {
	name, rest := firstToken(line)
	rawType, attrStr := firstToken(rest)

	f := field{name: name}
	f.array = strings.HasSuffix(rawType, "[]")
	rawType = strings.TrimSuffix(rawType, "[]")
	f.nullable = strings.HasSuffix(rawType, "?")
	f.baseType = strings.TrimSuffix(rawType, "?")
	f.attrs = splitAttrs(attrStr)
	return f
}
