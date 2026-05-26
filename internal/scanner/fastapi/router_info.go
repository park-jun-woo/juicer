//ff:type feature=scan type=model topic=fastapi
//ff:what 라우터 인스턴스 정보 구조체
package fastapi

// routerInfo tracks a FastAPI or APIRouter instance assignment.
type routerInfo struct {
	varName string // e.g., "app", "router", "v1_router"
	prefix  string // e.g., "/users" from APIRouter(prefix="/users")
	isFastAPI bool // true if FastAPI(), false if APIRouter()
}
