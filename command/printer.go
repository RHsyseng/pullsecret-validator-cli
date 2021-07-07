package command

import (
	"fmt"
	v "github.com/RHsyseng/lib-ps-validator"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorWhite  = "\033[37m"
)

func writeOutputTable(result v.WebData, filters []string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{string(ColorGreen) + "Valid Auths Entries" + string(ColorWhite), string(ColorRed) + "Expired Auths Entries" + string(ColorWhite), string(ColorYellow) + "Connection Issues" + string(ColorWhite)})

	t.AppendRow([]interface{}{string(ColorWhite) + fmt.Sprintf("%v", result.ResultOK), string(ColorWhite) + fmt.Sprintf("%v", result.ResultKO), string(ColorWhite) + fmt.Sprintf("%v", result.ResultCon)})

	t.SetStyle(table.StyleLight)
	t.Render()
}
