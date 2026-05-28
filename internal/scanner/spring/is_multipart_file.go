//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 타입명이 MultipartFile인지 확인한다
package spring

import "strings"

func isMultipartFile(typeName string) bool {
	return typeName == "MultipartFile" || strings.HasSuffix(typeName, ".MultipartFile")
}
