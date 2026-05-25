//ff:func feature=hurl type=command control=sequence
//ff:what TestHandleHurlSubcommand_Reset 테스트
package main

import "testing"

func TestHandleHurlSubcommand_Reset(t *testing.T) {
	_, cleanup := setupHurlSession(t)
	defer cleanup()
	handleHurlSubcommand([]string{"reset"})
}
