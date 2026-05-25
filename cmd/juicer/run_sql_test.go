//ff:func feature=sql type=command control=sequence
//ff:what TestRunSQL_Cov 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQL_DefaultDirCov(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runSQL([]string{})
}

func TestRunSQL_WithDirCov(t *testing.T) {
	dir := t.TempDir()
	runSQL([]string{dir})
}

func TestRunSQL_JSONCov(t *testing.T) {
	dir := t.TempDir()
	outFile := filepath.Join(dir, "out.json")
	runSQL([]string{dir, "--json", "-o", outFile})
}

func TestRunSQL_YAMLToFileCov(t *testing.T) {
	dir := t.TempDir()
	outFile := filepath.Join(dir, "out.yaml")
	runSQL([]string{dir, "-o", outFile})
}

func TestRunSQL_SubcommandNextCov(t *testing.T) {
	dir := t.TempDir()
	repoDir := filepath.Join(dir, "repo")
	queriesDir := filepath.Join(dir, "queries")
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)
	sessionDir := filepath.Join(dir, ".huma")
	os.MkdirAll(sessionDir, 0o755)
	sessionJSON := `{"repo_dir":"` + repoDir + `","queries_dir":"` + queriesDir + `","methods":[]}`
	os.WriteFile(filepath.Join(sessionDir, "sql-session.json"), []byte(sessionJSON), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runSQL([]string{"next"})
}
