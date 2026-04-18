package smoke

import (
	"testing"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func TestGatewaySmoke(t *testing.T) {
	env := fixtures.New(t)

	runSetupSmoke(t, env)
	runAuthSmoke(t, env)
	runHotelsSmoke(t, env)
	runRoomsSmoke(t, env)
	runBookingsSmoke(t, env)
	runCleanupSmoke(t, env)
}
