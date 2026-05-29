//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 파일 AST에서 aliased_import 매핑(로컬명→원본명)을 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// collectImportAliases walks all aliased_import nodes in a file and builds
// a local-name -> original-name map so register_blueprint variable names that
// reference an aliased import can be resolved back to the canonical Blueprint name.
func collectImportAliases(root *sitter.Node, src []byte) importAlias {
	aliases := make(importAlias)
	nodes := findAllByType(root, "aliased_import")
	for _, n := range nodes {
		local, orig := parseAliasedImport(n, src)
		if local != "" && orig != "" {
			aliases[local] = orig
		}
	}
	return aliases
}
