package cmd

/* https://github.com/KEINOS/Hello-Cobra */

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rootCmd(t *testing.T) {
	var (
		rootCmd = CreateRootCmd()
		argsTmp = []string{}
		buffTmp = new(bytes.Buffer)

		expect string
		actual string
	)

	rootCmd.SetOut(buffTmp)  // set output from os.Stdout -> buffTmp
	rootCmd.SetArgs(argsTmp) // set command args

	// Run `hello` command!
	if err := rootCmd.Execute(); err != nil {
		assert.FailNowf(t, "Failed to execute 'restoreCmd.Execute()'.", "Error msg: %v", err)
	}

	expect = ""
	actual = buffTmp.String() // resotre buffer
	assert.Equal(t, expect, actual,
		"Command 'build' with no parameters should produce an empty value.",
	)
}

func Test_rootCmdExecute(t *testing.T) {

	Execute()
}

func Test_rootCmdExecuteOtherConfig(t *testing.T) {

	cfgFile = ".mach"
	Execute()
}
