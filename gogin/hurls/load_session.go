//ff:func feature=hurl type=session control=sequence
//ff:what 디스크에서 세션 읽기 및 역직렬화
package hurls

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadSession reads and deserializes the hurl session from disk.
func LoadSession() (*Session, error) {
	data, err := os.ReadFile(sessionPath())
	if err != nil {
		return nil, fmt.Errorf("read session: %w", err)
	}
	var s Session
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("parse session: %w", err)
	}
	return &s, nil
}
