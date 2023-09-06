package webmachine

var Runtimes = map[string]*Runtime{
	"go": {
		InitScript: "go mod init {{.Name}}",
	},
}
