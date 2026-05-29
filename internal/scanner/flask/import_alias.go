//ff:type feature=scan type=model topic=flask
//ff:what import 로컬명 → 원본명 매핑 타입
package flask

// importAlias maps a locally bound import name to its original (canonical) name.
// e.g., from .auth import auth as auth_blueprint -> {"auth_blueprint": "auth"}
type importAlias map[string]string
