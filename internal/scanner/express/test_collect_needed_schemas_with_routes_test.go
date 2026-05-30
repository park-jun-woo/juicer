//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectNeededSchemas_WithRoutes 테스트
package express

import "testing"

func TestCollectNeededSchemas_WithRoutes(t *testing.T) {
	src := []byte(`
const app = express();
app.post('/users', validate(UserSchema), (req, res) => {});
`)
	fi := mustParse(t, src)
	ctx := &scanContext{
		parsed:     map[string]*fileInfo{"a.ts": fi},
		allRouters: map[string]map[string]bool{"a.ts": {"app": true}},
	}

	names := collectNeededSchemas(ctx)

	_ = names
}
