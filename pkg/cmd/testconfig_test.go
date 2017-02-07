package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/drud/drud-go/utils"
)

var (
	// DrudBin is the full path to the drud binary
	DrudBin = "ddev"

	// DevTestEnv is the name of the Dev DRUD environment to test
	DevTestEnv = "production"

	// DevTestApp is the name of the Dev DRUD app to test
	DevTestApp = "drudio"

	DevTestSites = [][]string{
		[]string{"drudio", DevTestEnv},
		[]string{"d8", DevTestEnv},
		[]string{"talentreef", DevTestEnv},
	}
)

func TestMain(m *testing.M) {
	if os.Getenv("DDEV_BINARY_FULLPATH") != "" {
		DrudBin = os.Getenv("DDEV_BINARY_FULLPATH")
	}

	err := os.Setenv("DRUD_NONINTERACTIVE", "true")
	if err != nil {
		fmt.Println("could not set noninteractive mode")
	}

	fmt.Println("Running tests.")
	os.Exit(m.Run())
}

func setActiveApp(appName string, deployName string) error {
	if appName == "" && deployName == "" {
		_, err := utils.RunCommand(DrudBin, []string{"config", "unset", "--activeapp", "--activedeploy"})
		return err
	}

	_, err := utils.RunCommand(DrudBin, []string{"config", "set", "--activeapp", appName, "--activedeploy", deployName})
	return err
}