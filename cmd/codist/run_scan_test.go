//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestRunScan_YAMLDefault YAML 기본 출력 분기 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunScan_YAMLDefault(t *testing.T) {
	dir := setupMinimalGoProject(t)
	// supafunc scanner requires supabase/functions to exist
	if err := os.MkdirAll(filepath.Join(dir, "supabase", "functions"), 0o755); err != nil {
		t.Fatal(err)
	}

	// default: framework auto-detected (gogin) -> YAML to stdout
	runScan([]string{dir})

	// each explicit framework branch in the switch
	frameworks := []string{
		"gogin", "fiber", "echo", "nestjs", "fastapi", "flask",
		"django", "fastify", "hono", "express", "spring",
		"quarkus", "dotnet", "supafunc", "actix", "laravel",
	}
	for _, fw := range frameworks {
		runScan([]string{"--framework", fw, dir})
	}

	// json output branch
	runScan([]string{"--framework", "gogin", "--json", dir})

	// openapi output branch
	runScan([]string{"--framework", "gogin", "--openapi", dir})

	// output-file branch
	out := filepath.Join(t.TempDir(), "spec.yaml")
	runScan([]string{"--framework", "gogin", "-o", out, dir})
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("expected output file written: %v", err)
	}
}
