package data

import "github.com/mikerybka/webmachine/pkg/types"

var Runtimes = []types.Runtime{
	{
		FileName:  "main.go",
		RunPrefix: []string{"go", "run"},
	},
	{
		FileName:  "main.py",
		RunPrefix: []string{"python3"},
	},
	{
		FileName:  "main.rb",
		RunPrefix: []string{"ruby"},
	},
	{
		FileName:  "main.js",
		RunPrefix: []string{"node"},
	},
}
