package main

import (
	"fmt"
	"strings"

	"github.com/nfx/go-htmltable"

	_ "embed"
)

//go:embed perms.html
var data string

func main() {
	type Ticker struct {
		Permission  string `header:"Permission"`
		Value       string `header:"Value"`
		Description string `header:"Description"`
		CType       string `header:"Channel Type"`
	}

	out, err := htmltable.NewSliceFromString[Ticker](data)
	if err != nil {
		panic(err)
	}
	fmt.Println("package perms")
	fmt.Println("var Dict = map[int64]DictItem{")
	seen := []string{}
	for _, o := range out {
		a := strings.Split(o.Permission, " ")
		fmt.Println(a[0] + ": {")
		fmt.Println("\t\"" + a[0] + "\",")
		seen = append(seen, a[0])
		if len(a) > 1 {
			if a[1] == "*" {
				fmt.Println("\tPERM_WARN_2FA,")
			} else {
				fmt.Println("\tPERM_WARN_TIMEOUT,")
			}
		} else {
			fmt.Println("\tPERM_WARN_NONE,")
		}

		v := strings.ReplaceAll(o.Value, "(", " ")
		v = strings.ReplaceAll(v, ")", "")
		b := strings.Split(v, " ")

		fmt.Println("\t" + b[len(b)-1] + ",")

		fmt.Println("\t\"" + strings.ReplaceAll(o.Description, "\n", `\n`) + "\",")

		if len(o.CType) == 0 {
			fmt.Println("\tPERM_UNTYPED,")
		} else {
			fmt.Println("\t" + strings.ReplaceAll(o.CType, ",", "|") + ",")
		}
		fmt.Println("},")
	}
	fmt.Println("}")

	fmt.Println("const (")
	fmt.Println("\t"+seen[0], "= 1 << iota")
	fmt.Println("\t" + strings.Join(seen[1:], "\n\t"))
	fmt.Println(")")

	fmt.Println("const ALL = ", strings.Join(seen, "|\n"))
}
