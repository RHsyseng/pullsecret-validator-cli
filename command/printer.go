package command

import (
	"encoding/json"
	"fmt"
	v "github.com/RHsyseng/lib-ps-validator"
	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/yaml.v3"
	_ "gopkg.in/yaml.v3"
	"os"
	"strings"
)

const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorWhite  = "\033[37m"
)

type Output struct {
	Auths struct {
		Valid           []string `yaml:"valid" json:"valid"`
		Expired         []string `yaml:"expired" json:"expired"`
		ConnectionIssue []string `yaml:"connection_issue" json:"connection_issue"`
	} `yaml:"auths" json:"auths"`
}

func writeOutputTable(result v.WebData) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{string(ColorGreen) + "Valid Auths Entries" + string(ColorWhite), string(ColorRed) + "Expired Auths Entries" + string(ColorWhite), string(ColorYellow) + "Connection Issues" + string(ColorWhite)})

	t.AppendRow([]interface{}{string(ColorWhite) + fmt.Sprintf("%v", result.ResultOK), string(ColorWhite) + fmt.Sprintf("%v", result.ResultKO), string(ColorWhite) + fmt.Sprintf("%v", result.ResultCon)})

	t.SetStyle(table.StyleLight)
	t.Render()
}

func writeOutputJson(result v.WebData) {
	ok := strings.Fields(fmt.Sprintf("%v", result.ResultOK))
	ko := strings.Fields(fmt.Sprintf("%v", result.ResultKO))
	con := strings.Fields(fmt.Sprintf("%v", result.ResultCon))

	out := Output{
		Auths: struct {
			Valid           []string `yaml:"valid" json:"valid"`
			Expired         []string `yaml:"expired" json:"expired"`
			ConnectionIssue []string `yaml:"connection_issue" json:"connection_issue"`
		}{
			Valid:           ok,
			Expired:         ko,
			ConnectionIssue: con,
		},
	}

	o, _ := json.Marshal(out)
	fmt.Println(string(o))
}

func writeOutputYaml(result v.WebData) {
	ok := strings.Fields(fmt.Sprintf("%v", result.ResultOK))
	ko := strings.Fields(fmt.Sprintf("%v", result.ResultKO))
	con := strings.Fields(fmt.Sprintf("%v", result.ResultCon))

	out := Output{
		Auths: struct {
			Valid           []string `yaml:"valid" json:"valid"`
			Expired         []string `yaml:"expired" json:"expired"`
			ConnectionIssue []string `yaml:"connection_issue" json:"connection_issue"`
		}{
			Valid:           ok,
			Expired:         ko,
			ConnectionIssue: con,
		},
	}

	o, _ := yaml.Marshal(out)
	fmt.Println(string(o))
}
