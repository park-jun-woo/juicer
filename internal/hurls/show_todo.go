//ff:func feature=hurl type=render control=sequence
//ff:what TODO 엔드포인트의 scan 스켈레톤 또는 기본 TODO 출력
package hurls

import (
	"fmt"
)

// showTODO prints skeleton info for a TODO endpoint, falling back to basic output.
func showTODO(ep *EndpointStatus, repoDir, testsDir string) {
	scanEP := scanEndpoint(repoDir, ep.ID)
	if scanEP != nil {
		printSkeleton(scanEP, testsDir)
		return
	}
	_, path := parseEndpointID(ep.ID)
	fmt.Printf("%s  TODO\n", ep.ID)
	fmt.Printf("  -> Write test to %s%s\n", testsDir, suggestFilename(path))
}
