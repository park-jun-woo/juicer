//ff:type feature=scan type=model topic=fastapi
//ff:what import 정보 구조체
package fastapi

// importInfo maps imported names to their source module path.
type importInfo struct {
	name   string // imported name (e.g., "UserCreate")
	module string // module path (e.g., ".models" or "app.models")
}
