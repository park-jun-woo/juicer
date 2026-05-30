//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestLastSegmentSingular 테스트
package laravel

import "testing"

func TestLastSegmentSingular(t *testing.T) {
	if got := lastSegmentSingular("users"); got != "user" {
		t.Fatalf("got %q", got)
	}
	if got := lastSegmentSingular("{product}/reviews"); got != "review" {
		t.Fatalf("got %q", got)
	}
}
