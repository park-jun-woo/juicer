//ff:func feature=scan type=command control=sequence
//ff:what TestRunScan_Cov 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunScan_WithRootCov(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "out.yaml")
	runScan([]string{dir, "-o", outFile})
}

func TestRunScan_JSONCov(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "out.json")
	runScan([]string{dir, "--json", "-o", outFile})
}

func TestRunScan_OpenAPICov(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "out.yaml")
	runScan([]string{dir, "--openapi", "-o", outFile})
}

func TestRunScan_StdoutCov(t *testing.T) {
	dir := setupMinimalGoProject(t)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runScan([]string{})
}
