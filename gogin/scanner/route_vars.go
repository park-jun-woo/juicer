package scanner

var ginMethods = map[string]bool{
	"GET": true, "POST": true, "PUT": true, "PATCH": true, "DELETE": true,
	"HEAD": true, "OPTIONS": true, "Any": true,
}

var ginRouterTypes = map[string]bool{
	"Engine":      true,
	"RouterGroup": true,
	"IRouter":     true,
	"IRoutes":     true,
}
