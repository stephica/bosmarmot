package tendermint

import (
	"testing"

	"os"
	"github.com/tendermint/tendermint/config"
	"code.monax.io/platform/bosmarmot/logging/lifecycle"
)

const testDir = "./scratch"

func TestLaunchGenesisValidator(t *testing.T) {
	os.RemoveAll(testDir)
	os.MkdirAll(testDir, 0777)
	os.Chdir(testDir)
	conf := config.DefaultConfig()
	logger, _ := lifecycle.NewStdErrLogger()
	LaunchGenesisValidator(conf, logger)
}
