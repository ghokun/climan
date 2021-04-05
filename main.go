package main

import "github.com/ghokun/climan/cmd"

var (
	version = "dev"
	commit  = "none"
)

func main() {
	cmd.Execute(version, commit)
}
