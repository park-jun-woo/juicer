//ff:func feature=scan type=test control=sequence topic=actix
//ff:what applySerdePart — default/skip/rename 토큰 파싱 분기를 검증
package actix

import "testing"

func TestApplySerdePart(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := &serdeAttr{}
		applySerdePart(r, "default")
		if !r.hasDefault {
			t.Error("expected hasDefault")
		}
	})
	t.Run("skip", func(t *testing.T) {
		r := &serdeAttr{}
		applySerdePart(r, "skip")
		if !r.skip {
			t.Error("expected skip")
		}
	})
	t.Run("skip_deserializing", func(t *testing.T) {
		r := &serdeAttr{}
		applySerdePart(r, "skip_deserializing")
		if !r.skip {
			t.Error("expected skip for skip_deserializing")
		}
	})
	t.Run("rename", func(t *testing.T) {
		r := &serdeAttr{}
		applySerdePart(r, `rename = "userName"`)
		if r.rename != "userName" {
			t.Errorf("rename = %q, want userName", r.rename)
		}
	})
	t.Run("rename_no_eq", func(t *testing.T) {
		r := &serdeAttr{}
		applySerdePart(r, "rename") // no '=' -> early return, no change
		if r.rename != "" {
			t.Errorf("expected empty rename, got %q", r.rename)
		}
	})
	t.Run("other", func(t *testing.T) {
		r := &serdeAttr{}
		applySerdePart(r, "flatten") // none of the conditions -> no change
		if r.hasDefault || r.skip || r.rename != "" {
			t.Errorf("expected no change, got %+v", r)
		}
	})
}
