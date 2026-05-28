//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what Fastify 경로의 :param을 OpenAPI {param}으로 변환한다
package fastify

import "regexp"

var colonParamRe = regexp.MustCompile(`:([a-zA-Z_][a-zA-Z0-9_]*)(\([^)]*\))?`)

func fastifyPathToOpenAPI(path string) string {
	return colonParamRe.ReplaceAllString(path, "{$1}")
}
