package cmd

/* https://github.com/KEINOS/Hello-Cobra */

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_serveCmd(t *testing.T) {
	var (
		serveCmd = CreateServeCmd()
		argsTmp  = []string{}
		buffTmp  = new(bytes.Buffer)

		expect string
		actual string
	)

	serveCmd.SetOut(buffTmp)  // set output from os.Stdout -> buffTmp
	serveCmd.SetArgs(argsTmp) // set command args

	// Run `hello` command!
	if err := serveCmd.Execute(); err != nil {
		assert.FailNowf(t, "Failed to execute 'serveCmd.Execute()'.", "Error msg: %v", err)
	}

	expect = ""
	actual = buffTmp.String() // resotre buffer
	assert.Equal(t, expect, actual,
		"Command 'serve' with no parameters should produce an empty value.",
	)
}

func Test_serveCmd_Help(t *testing.T) {
	var (
		serveCmd = CreateServeCmd()
		argsTmp  = []string{"--help"}
		buffTmp  = new(bytes.Buffer)

		expect string
		actual string
	)

	serveCmd.SetOut(buffTmp)  // set output from os.Stdout -> buffTmp
	serveCmd.SetArgs(argsTmp) // set command args

	// Run `hello` command!
	if err := serveCmd.Execute(); err != nil {
		assert.FailNowf(t, "Failed to execute 'serveCmd.Execute()'.", "Error msg: %v", err)
	}

	expect = "Usage:"
	actual = buffTmp.String() // resotre buffer
	assert.Contains(t, actual, expect,
		"Command 'help' should show usage",
	)
}
