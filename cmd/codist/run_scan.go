//ff:func feature=scan type=command control=selection
//ff:what scan 실행 — 프레임워크 감지 후 해당 스캐너로 엔드포인트를 추출한다
package main

import (
	"fmt"

	"github.com/park-jun-woo/codistill/internal/scanner"
	"github.com/park-jun-woo/codistill/internal/scanner/actix"
	"github.com/park-jun-woo/codistill/internal/scanner/django"
	"github.com/park-jun-woo/codistill/internal/scanner/dotnet"
	echoScanner "github.com/park-jun-woo/codistill/internal/scanner/echo"
	"github.com/park-jun-woo/codistill/internal/scanner/express"
	"github.com/park-jun-woo/codistill/internal/scanner/fastapi"
	"github.com/park-jun-woo/codistill/internal/scanner/fastify"
	"github.com/park-jun-woo/codistill/internal/scanner/fiber"
	"github.com/park-jun-woo/codistill/internal/scanner/flask"
	"github.com/park-jun-woo/codistill/internal/scanner/gogin"
	"github.com/park-jun-woo/codistill/internal/scanner/hono"
	"github.com/park-jun-woo/codistill/internal/scanner/laravel"
	"github.com/park-jun-woo/codistill/internal/scanner/nestjs"
	"github.com/park-jun-woo/codistill/internal/scanner/quarkus"
	"github.com/park-jun-woo/codistill/internal/scanner/spring"
	"github.com/park-jun-woo/codistill/internal/scanner/supafunc"
)

func runScan(root string, o scanOptions) error {
	fw := o.framework
	if fw == "" {
		fw = scanner.DetectFramework(root)
		if fw == "" {
			return fmt.Errorf("could not detect framework; specify --framework")
		}
	}

	var result *scanner.ScanResult
	var err error

	switch fw {
	case "gogin":
		result, err = gogin.Scan(root)
	case "fiber":
		result, err = fiber.Scan(root)
	case "echo":
		result, err = echoScanner.Scan(root)
	case "nestjs":
		result, err = nestjs.Scan(root)
	case "fastapi":
		result, err = fastapi.Scan(root)
	case "flask":
		result, err = flask.Scan(root)
	case "django":
		result, err = django.Scan(root)
	case "fastify":
		result, err = fastify.Scan(root)
	case "hono":
		result, err = hono.Scan(root)
	case "express":
		result, err = express.Scan(root)
	case "spring":
		result, err = spring.Scan(root)
	case "quarkus":
		result, err = quarkus.Scan(root)
	case "dotnet":
		result, err = dotnet.Scan(root)
	case "supafunc":
		result, err = supafunc.Scan(root)
	case "actix":
		result, err = actix.Scan(root)
	case "laravel":
		result, err = laravel.Scan(root)
	default:
		err = fmt.Errorf("unknown framework: %s", fw)
	}
	if err != nil {
		return err
	}

	// 출력 포맷과 무관하게 (method, path) 기준 중복 엔드포인트를 한 번만 제거한다.
	result.Endpoints = scanner.DeduplicateEndpoints(result.Endpoints)
	return writeScanResult(result, root, o)
}
