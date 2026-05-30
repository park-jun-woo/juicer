//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestCollectResources 테스트
package quarkus

import "testing"

func TestCollectResources(t *testing.T) {
	fi := qFileInfo(t, sampleResource)
	got := collectResources([]*fileInfo{fi})
	if len(got) != 1 {
		t.Fatalf("got %d", len(got))
	}
}
