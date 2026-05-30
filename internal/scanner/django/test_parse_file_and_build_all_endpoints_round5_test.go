//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseFile_And_BuildAllEndpoints_Round5 테스트
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_And_BuildAllEndpoints_Round5(t *testing.T) {
	dir := t.TempDir()
	urls := "from django.urls import path\nfrom .views import PingView\nurlpatterns = [path('ping/', PingView.as_view())]\n"
	views := "from rest_framework.views import APIView\nfrom rest_framework.response import Response\nclass PingView(APIView):\n    def get(self, request):\n        return Response()\n"
	if err := os.WriteFile(filepath.Join(dir, "urls.py"), []byte(urls), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "views.py"), []byte(views), 0o644); err != nil {
		t.Fatal(err)
	}
	fiUrls, err := parseFile(dir, filepath.Join(dir, "urls.py"))
	if err != nil {
		t.Fatal(err)
	}
	if fiUrls.root == nil {
		t.Fatal("parseFile root nil")
	}
	fiViews, err := parseFile(dir, filepath.Join(dir, "views.py"))
	if err != nil {
		t.Fatal(err)
	}
	eps := buildAllEndpoints([]fileInfo{*fiUrls, *fiViews})
	if len(eps) == 0 {
		t.Fatalf("expected endpoints, got %d", len(eps))
	}
}
