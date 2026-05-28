//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestExtractBodyStatus — ResponseEntity 패턴 상태 코드 추출 테스트
package spring

import "testing"

func TestExtractBodyStatus(t *testing.T) {
	for _, tt := range bodyStatusTests {
		t.Run(tt.name, func(t *testing.T) {
			assertBodyStatusCode(t, tt.src, tt.want)
		})
	}
}
