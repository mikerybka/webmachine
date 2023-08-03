package data

import "github.com/mikerybka/webmachine/pkg/types"

var Runtimes = []types.Runtime{
	{
		FileName:  "main.go",
		CmdPrefix: []string{"go", "run"},
	},
	{
		FileName:  "main.py",
		CmdPrefix: []string{"python3"},
	},
	{
		FileName:  "main.rb",
		CmdPrefix: []string{"ruby"},
	},
	{
		FileName:  "main.js",
		CmdPrefix: []string{"node"},
	},
}
