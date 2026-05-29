//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 미들웨어 슬라이스를 복사해 공유 변경을 방지한다
package laravel

func copyMiddleware(middleware []string) []string {
	mw := make([]string, len(middleware))
	copy(mw, middleware)
	return mw
}
