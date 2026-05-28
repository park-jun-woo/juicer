//ff:type feature=scan type=model topic=flask
//ff:what Blueprint 변수명 → prefix 매핑 타입
package flask

// blueprintPrefix maps blueprint variable name to its resolved prefix.
type blueprintPrefix map[string]string
