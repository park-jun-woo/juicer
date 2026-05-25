//ff:func feature=hurl type=command control=sequence
//ff:what TestHandleHurlSubcommand_Next 테스트
package main

import "testing"

func TestHandleHurlSubcommand_Next(t *testing.T) {
	_, cleanup := setupHurlSession(t)
	defer cleanup()
	// next with empty session => "All tests complete!"
	handleHurlSubcommand([]string{"next"})
}
