//ff:func feature=sql type=parse control=sequence
//ff:what 세션을 디스크에 직렬화하여 저장
package sqls

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveSession serializes and writes the session to disk.
//
func SaveSession(s *Session) error {
	if err := os.MkdirAll(sessionDirName, 0o755); err != nil {
		return fmt.Errorf("create session dir: %w", err)
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal session: %w", err)
	}
	if err := os.WriteFile(sessionPath(), data, 0o644); err != nil {
		return fmt.Errorf("write session: %w", err)
	}
	return nil
}

