//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what autoloadVarNames: import/require로 바인딩된 autoload 식별자 수집을 검증
package fastify

import "testing"

func TestAutoloadVarNames(t *testing.T) {
	src := []byte(`
import autoload from "@fastify/autoload";
const al = require("@fastify/autoload");
import other from "./other";
`)
	fi := mustParse(t, src)
	names := autoloadVarNames(fi)
	if !names["autoload"] {
		t.Error("expected import binding 'autoload'")
	}
	if !names["al"] {
		t.Error("expected require binding 'al'")
	}
	if names["other"] {
		t.Error("non-autoload import must not be included")
	}
}
