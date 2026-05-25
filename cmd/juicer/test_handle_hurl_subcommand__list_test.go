//ff:func feature=hurl type=command control=sequence
//ff:what TestHandleHurlSubcommand_List 테스트
package main

import "testing"

func TestHandleHurlSubcommand_List(t *testing.T) {
	_, cleanup := setupHurlSession(t)
	defer cleanup()
	handleHurlSubcommand([]string{"list"})
}
