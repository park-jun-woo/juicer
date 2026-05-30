//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what firstRegisterMount 테스트 헬퍼
package fastify

import "testing"

func firstRegisterMount(t *testing.T, src string, inst map[string]bool) *pluginMount {
	t.Helper()
	fi, calls := regCalls(t, src)
	for _, c := range calls {
		if pm := extractRegisterCall(c, fi.Src, inst); pm != nil {
			return pm
		}
	}
	return nil
}
