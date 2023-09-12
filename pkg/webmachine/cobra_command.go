package webmachine

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/mikerybka/util"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var CobraCommand *cobra.Command = &cobra.Command{
	Use:   "webmachine",
	Short: "webmachine is a tool for building web apps",
	Long:  `webmachine is a tool for building web apps.`,
}

func init() {
	CobraCommand.AddCommand(initCommand)
	CobraCommand.AddCommand(devCommand)
	CobraCommand.AddCommand(serveCommand)
	CobraCommand.AddCommand(deployCommand)
}

var initCommand *cobra.Command = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new webmachine project",
	Long:  `Initialize a new webmachine project.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdInit()
	},
}

var devCommand *cobra.Command = &cobra.Command{
	Use:   "dev",
	Short: "Run a development server",
	Long:  `Run a development server.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdDev()
	},
}

var serveCommand *cobra.Command = &cobra.Command{
	Use:   "serve",
	Short: "Start a production server",
	Long:  `Start a production server.`,
	Run: func(cmd *cobra.Command, args []string) {
		email := ""
		if len(args) > 1 {
			email = args[1]
		}
		cmdServe(email)
	},
}

var deployCommand *cobra.Command = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy to a production server",
	Long:  `Deploy to a production server.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdDeploy()
	},
}

func cmdInit() {
	flag.Parse()
	name := flag.Arg(1)
	var runtimesString string = flag.Arg(2)
	runtimes := []*Runtime{}
	for _, id := range strings.Split(runtimesString, ",") {
		rt, ok := Runtimes[id]
		if ok {
			runtimes = append(runtimes, rt)
		} else if id != "" {
			fmt.Println("Unknown runtime ID:", id)
		}
	}

	if name == "" {
		flag.Usage()
		return
	}

	err := Init(name, runtimes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cmdDev() {
	server := Server{
		Dir:     ".",
		DevMode: true,
		Logging: true,
	}
	browser.OpenURL("http://localhost:3000")
	err := http.ListenAndServe(":3000", &server)
	if err != nil {
		panic(err)
	}
}

func cmdDeploy() {
	cmd := exec.Command("rsync", "-avz", "--delete", "--exclude", ".git", ".", "/etc/web/")

	remoteHostname := flag.Arg(1)
	if remoteHostname != "" {
		cmd = exec.Command("rsync", "-avz", "--delete", "--exclude", ".git", ".", "root@"+remoteHostname+":/etc/web/")
	}

	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func cmdServe(email string) {
	server := Server{
		Dir:     "/etc/web",
		Logging: true,
	}
	err := util.ServeHTTPS(&server, email, "/etc/certs")
	if err != nil {
		panic(err)
	}
}
