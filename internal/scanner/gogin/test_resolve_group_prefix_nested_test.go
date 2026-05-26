//ff:func feature=scan type=test control=sequence
//ff:what TestResolveGroupPrefix_NestedGroup 테스트
package gogin

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestResolveGroupPrefix_NestedGroup(t *testing.T) {
	dir, err := os.MkdirTemp("", "gogin-nested-group-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	gomod := `module example.com/test

go 1.21

require github.com/gin-gonic/gin v1.10.0
`
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(gomod), 0644); err != nil {
		t.Fatal(err)
	}

	// main.go — parent router has prefix from Group, then inline Group adds more
	mainSrc := `package main

import (
	"github.com/gin-gonic/gin"
	"example.com/test/webhook"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	webhook.RegisterRoutes(api.Group("/v1/webhook"), nil)
}
`
	if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte(mainSrc), 0644); err != nil {
		t.Fatal(err)
	}

	webhookDir := filepath.Join(dir, "webhook")
	if err := os.MkdirAll(webhookDir, 0755); err != nil {
		t.Fatal(err)
	}
	handlerSrc := `package webhook

import "github.com/gin-gonic/gin"

type Handler struct{}

func RegisterRoutes(group *gin.RouterGroup, h *Handler) {
	group.POST("/sms", func(c *gin.Context) {})
}
`
	if err := os.WriteFile(filepath.Join(webhookDir, "handler.go"), []byte(handlerSrc), 0644); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Skipf("go mod tidy failed: %v\n%s", err, out)
	}

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}

	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d: %+v", len(result.Endpoints), result.Endpoints)
	}

	ep := result.Endpoints[0]
	expected := "POST /api/v1/webhook/sms"
	actual := ep.Method + " " + ep.Path
	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
