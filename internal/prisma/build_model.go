//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 모델명+본문 라인을 model 구조체(필드 + 블록속성 + @@map 테이블명)로 조립
package prisma

import "strings"

// buildModel assembles a model from its name and body lines.
func buildModel(name string, body []string) model {
	m := model{name: name, tableName: name}
	for _, line := range body {
		if strings.HasPrefix(line, "@@") {
			applyBlockAttr(&m, line)
			continue
		}
		m.fields = append(m.fields, parseField(line))
	}
	return m
}
