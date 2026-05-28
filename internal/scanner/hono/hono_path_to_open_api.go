//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what Hono 경로의 :param을 OpenAPI {param}으로 변환한다
package hono

import "regexp"

var colonParamRe = regexp.MustCompile(`:([a-zA-Z_][a-zA-Z0-9_]*)(\([^)]*\))?`)

func honoPathToOpenAPI(path string) string {
	return colonParamRe.ReplaceAllString(path, "{$1}")
}
