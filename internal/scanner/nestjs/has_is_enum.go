//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what validators 목록에 IsEnum이 포함되어 있는지 확인한다
package nestjs

// hasIsEnum returns true if any validator starts with "IsEnum".
func hasIsEnum(validators []string) bool {
	for _, v := range validators {
		if v == "IsEnum" {
			return true
		}
	}
	return false
}
