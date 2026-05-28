//ff:type feature=scan type=model topic=spring
//ff:what 역할 추출용 정규표현식 패턴
package spring

import "regexp"

var hasRoleRegexp = regexp.MustCompile(`hasRole\(['"]([^'"]+)['"]\)`)
var hasAnyRoleRegexp = regexp.MustCompile(`hasAnyRole\(([^)]+)\)`)
var roleStringRegexp = regexp.MustCompile(`['"]([^'"]+)['"]`)
