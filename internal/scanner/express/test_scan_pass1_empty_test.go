//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestScanPass1_Empty 테스트
package express

import "testing"

func TestScanPass1_Empty(t *testing.T) {
	ctx := scanPass1(nil, t.TempDir())
	if ctx == nil || len(ctx.parsed) != 0 {
		t.Fatalf("expected empty ctx, got %+v", ctx)
	}
}
