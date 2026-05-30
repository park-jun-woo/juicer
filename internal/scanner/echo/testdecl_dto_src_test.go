//ff:type feature=scan type=test topic=echo
//ff:what dtoSrc 테스트 보조 선언
package echo

const dtoSrc = `package m
type Inner struct {
	Code int ` + "`json:\"code\"`" + `
}
type Base struct {
	ID int ` + "`json:\"id\"`" + `
}
type UserDto struct {
	Base
	Name   string ` + "`json:\"name\"`" + `
	Hidden string ` + "`json:\"-\"`" + `
	Nested Inner  ` + "`json:\"nested\"`" + `
}
var U UserDto
var L []UserDto
`
