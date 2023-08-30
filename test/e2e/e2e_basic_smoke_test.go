package test_test

import (
	"testing"

	"github.com/defenseunicorns/uds-capability-artifactory/test/e2e/types"
	"github.com/defenseunicorns/uds-capability-artifactory/test/e2e/utils"
	teststructure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/require"
)

// TestAllServicesRunning waits until all services report that they are ready.
func TestAllServicesRunning(t *testing.T) { //nolint:funlen
	// BOILERPLATE, EXPECTED TO BE PRESENT AT THE BEGINNING OF EVERY TEST FUNCTION

	t.Parallel()
	platform := types.NewTestPlatform(t)
	defer platform.Teardown()
	utils.SetupTestPlatform(t, platform)
	// The repo has now been downloaded to /root/app and the software factory package deployment has been initiated.
	teststructure.RunTestStage(platform.T, "TEST", func() {
		// END BOILERPLATE

		// TEST CODE STARTS HERE.

		// Just make sure we can hit the cluster
		output, err := platform.RunSSHCommandAsSudo(`kubectl get nodes`)
		require.NoError(t, err, output)

		// Wait for the Artifactory StatefulSet to exist.
		output, err = platform.RunSSHCommandAsSudo(`timeout 1200 bash -c "while ! kubectl get statefulset artifactory -n artifactory; do sleep 5; done"`)
		require.NoError(t, err, output)

		// Wait for the Artifactory StatefulSet to report that it is ready
		output, err = platform.RunSSHCommandAsSudo(`kubectl rollout status statefulset/artifactory -n artifactory --watch --timeout=1200s`)
		require.NoError(t, err, output)

		// Setup DNS records for cluster services
		output, err = platform.RunSSHCommandAsSudo(`cd ~/app && utils/metallb/dns.sh && utils/metallb/hosts-write.sh`)
		require.NoError(t, err, output)

		// Ensure that Artifactory is available outside of the cluster.
		output, err = platform.RunSSHCommandAsSudo(`timeout 1200 bash -c "while ! curl -L -s --fail --show-error https://artifactory.bigbang.dev/artifactory/api/system/ping > /dev/null; do sleep 5; done"`)
		require.NoError(t, err, output)
	})
}
