package types_test

import (
	"os"
	"testing"

	"github.com/elysium-station/black/app"
)

func TestMain(m *testing.M) {
	app.SetSDKConfig()
	os.Exit(m.Run())
}
