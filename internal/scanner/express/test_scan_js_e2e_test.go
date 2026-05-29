//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what E2E: 순수 JS(CommonJS require/module.exports) 프로젝트의 크로스 파일 마운트 prefix 합성 검증
package express

import "testing"

func TestScan_JS_CommonJSMountPrefix(t *testing.T) {
	dir := t.TempDir()

	// leaf route file (CommonJS)
	authRoute := `
const express = require('express');
const router = express.Router();
router.post('/register', validate, authController.register);
router.post('/login', authController.login);
module.exports = router;
`
	writeFile(t, dir, "routes/v1/auth.route.js", authRoute)

	// v1 index mounts the auth route under /auth (array-of-objects mount)
	v1Index := `
const express = require('express');
const authRoute = require('./auth.route');
const router = express.Router();
const routes = [{ path: '/auth', route: authRoute }];
routes.forEach((route) => { router.use(route.path, route.route); });
module.exports = router;
`
	writeFile(t, dir, "routes/v1/index.js", v1Index)

	// app mounts the v1 router under /v1
	appSrc := `
const express = require('express');
const routes = require('./routes/v1');
const app = express();
app.use('/v1', routes);
app.get('/health', healthCheck);
`
	writeFile(t, dir, "app.js", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	expected := []string{
		"POST /v1/auth/register",
		"POST /v1/auth/login",
		"GET /health",
	}
	for _, e := range expected {
		if !found[e] {
			t.Errorf("missing endpoint %s, got %v", e, found)
		}
	}
}
