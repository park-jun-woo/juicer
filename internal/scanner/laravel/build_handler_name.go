//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what controller와 action으로 핸들러 문자열을 만든다
package laravel

// buildHandlerName creates handler string from controller and action.
func buildHandlerName(controller, action string) string {
	if controller == "" && action == "" {
		return "closure"
	}
	if controller == "" {
		return action
	}
	if action == "" {
		return controller
	}
	return controller + "@" + action
}
