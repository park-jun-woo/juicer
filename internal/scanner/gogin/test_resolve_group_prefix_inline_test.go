//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestResolveGroupPrefix_InlineGroup 테스트
package gogin

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestResolveGroupPrefix_InlineGroup(t *testing.T) {
	dir, err := os.MkdirTemp("", "gogin-group-prefix-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// go.mod
	gomod := `module example.com/test

go 1.21

require github.com/gin-gonic/gin v1.10.0
`
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(gomod), 0644); err != nil {
		t.Fatal(err)
	}

	// main.go — calls api.RegisterRoutes with inline Group
	mainSrc := `package main

import (
	"github.com/gin-gonic/gin"
	"example.com/test/api"
)

func main() {
	r := gin.Default()
	api.RegisterRoutes(r.Group("/api/v1"), nil)
}
`
	if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte(mainSrc), 0644); err != nil {
		t.Fatal(err)
	}

	// api/handler.go — receives *gin.RouterGroup and registers routes
	apiDir := filepath.Join(dir, "api")
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		t.Fatal(err)
	}
	handlerSrc := `package api

import "github.com/gin-gonic/gin"

type Handler struct{}

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.POST("/users", func(c *gin.Context) {})
	rg.GET("/users/:id", func(c *gin.Context) {})
}
`
	if err := os.WriteFile(filepath.Join(apiDir, "handler.go"), []byte(handlerSrc), 0644); err != nil {
		t.Fatal(err)
	}

	// go mod tidy to resolve dependencies
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Skipf("go mod tidy failed (gin may not be cached): %v\n%s", err, out)
	}

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}

	// Expect 2 endpoints: POST /api/v1/users and GET /api/v1/users/{id}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d: %+v", len(result.Endpoints), result.Endpoints)
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		key := ep.Method + " " + ep.Path
		found[key] = true
	}

	if !found["GET /api/v1/users/{id}"] {
		t.Errorf("missing GET /api/v1/users/{id}, endpoints: %+v", result.Endpoints)
	}
	if !found["POST /api/v1/users"] {
		t.Errorf("missing POST /api/v1/users, endpoints: %+v", result.Endpoints)
	}
}
