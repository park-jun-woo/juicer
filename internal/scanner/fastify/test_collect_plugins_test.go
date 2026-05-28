//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what fastify.register() 플러그인 수집 테스트
package fastify

import "testing"

func TestCollectPlugins(t *testing.T) {
	src := []byte(`
import Fastify from "fastify";
import userRoutes from "./routes/users";
const app = Fastify();
app.register(userRoutes, { prefix: "/api/users" });
app.register(require("./routes/posts"), { prefix: "/api/posts" });
`)
	fi := mustParse(t, src)
	instances := collectInstances(fi)
	mounts := collectPlugins(fi, instances)
	if len(mounts) != 2 {
		t.Fatalf("expected 2 plugin mounts, got %d", len(mounts))
	}
	if mounts[0].PluginRef != "userRoutes" {
		t.Errorf("mount[0].PluginRef: want userRoutes, got %s", mounts[0].PluginRef)
	}
	if mounts[0].Prefix != "/api/users" {
		t.Errorf("mount[0].Prefix: want /api/users, got %s", mounts[0].Prefix)
	}
	if mounts[1].PluginRef != "./routes/posts" {
		t.Errorf("mount[1].PluginRef: want ./routes/posts, got %s", mounts[1].PluginRef)
	}
	if mounts[1].Prefix != "/api/posts" {
		t.Errorf("mount[1].Prefix: want /api/posts, got %s", mounts[1].Prefix)
	}
}
