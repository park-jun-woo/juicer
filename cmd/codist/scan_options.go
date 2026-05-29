//ff:type feature=scan type=model
//ff:what scan 명령의 파싱된 플래그 값 묶음
package main

type scanOptions struct {
	jsonOut   bool
	openapi   bool
	baseFile  string
	outFile   string
	framework string
}
