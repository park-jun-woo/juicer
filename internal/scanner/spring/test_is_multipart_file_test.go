//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestIsMultipartFile 테스트
package spring

import "testing"

func TestIsMultipartFile(t *testing.T) {
	if !isMultipartFile("MultipartFile") || !isMultipartFile("org.x.MultipartFile") || isMultipartFile("String") {
		t.Fatal("multipart")
	}
}
