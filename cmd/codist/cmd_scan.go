//ff:func feature=scan type=command control=sequence
//ff:what scan 서브커맨드 빌더 — 플래그 등록 후 RunE에서 runScan에 위임한다
package main

import (
	"github.com/spf13/cobra"
)

func newScanCmd() *cobra.Command {
	var o scanOptions
	cmd := &cobra.Command{
		Use:   "scan [project-root]",
		Short: "Scan source code and extract endpoints",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := "."
			if len(args) > 0 {
				root = args[0]
			}
			return runScan(root, o)
		},
	}
	cmd.Flags().BoolVar(&o.jsonOut, "json", false, "output JSON")
	cmd.Flags().BoolVar(&o.openapi, "openapi", false, "output OpenAPI 3.0 YAML")
	cmd.Flags().StringVar(&o.baseFile, "base", "", "base OpenAPI spec to merge with")
	cmd.Flags().StringVarP(&o.outFile, "output", "o", "", "output file path")
	cmd.Flags().StringVar(&o.framework, "framework", "", "framework to scan (gogin, fiber, echo, nestjs, fastify, hono, fastapi, flask, django, express, spring, quarkus, dotnet, supafunc, actix, laravel)")
	return cmd
}
