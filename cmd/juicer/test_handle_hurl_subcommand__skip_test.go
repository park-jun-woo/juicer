//ff:func feature=hurl type=command control=sequence
//ff:what TestHandleHurlSubcommand_Skip 테스트
package main

import "testing"

func TestHandleHurlSubcommand_Skip(t *testing.T) {
	_, cleanup := setupHurlSession(t)
	defer cleanup()
	handleHurlSubcommand([]string{"skip"})
}
