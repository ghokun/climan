package cmd

import (
	"bytes"
	"text/template"

	"github.com/spf13/cobra"
)

const (
	logo = `
█▀▀ █░░ █ █▀▄▀█ ▄▀█ █▄░█  Version : {{.Version}}
█▄▄ █▄▄ █ █░▀░█ █▀█ █░▀█  Commit  : {{.Commit}}`
	versionTemplate = `{{.Version}}
`
)

var rootCmd = &cobra.Command{
	Use:   "climan",
	Short: "CLI tools version MANager for cloud native technologies",
}

func getLogo(version, commit string) string {
	tmpl, err := template.New("logo").Parse(logo)
	cobra.CheckErr(err)
	var buffer bytes.Buffer
	data := struct {
		Version string
		Commit  string
	}{
		Version: version,
		Commit:  commit,
	}
	err = tmpl.Execute(&buffer, data)
	cobra.CheckErr(err)
	return buffer.String()
}

func Execute(version, commit string) {
	rootCmd.Long = getLogo(version, commit)
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(versionTemplate)
	cobra.CheckErr(rootCmd.Execute())
}
