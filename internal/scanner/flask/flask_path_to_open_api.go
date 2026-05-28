//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what Flask URL 규칙의 <type:name>을 OpenAPI {name} 형식으로 변환한다
package flask

import "regexp"

// flaskParamRe matches Flask URL rule variables: <name>, <int:name>, <float:price>, etc.
var flaskParamRe = regexp.MustCompile(`<(?:([a-zA-Z_]+):)?([a-zA-Z_][a-zA-Z0-9_]*)>`)

// flaskPathToOpenAPI converts Flask URL rules to OpenAPI path format.
// e.g., "/users/<int:user_id>" -> "/users/{user_id}"
func flaskPathToOpenAPI(path string) string {
	return flaskParamRe.ReplaceAllString(path, "{$2}")
}
