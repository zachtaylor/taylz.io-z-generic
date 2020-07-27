package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
)

func main() {
	def := env()
	env := env().ParseFlags()
	fileName := "gen_cache.go"

	for k := range env {
		if env[k] == def[k] {
			fmt.Println("go-gengen: not set", k)
			return
		}
	}

	if _, err := os.Stat(fileName); err == nil {
		os.Remove(fileName)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("go-gengen: internal error", err)
		return
	}
	defer file.Close()

	stdlib := []string{"sync"}
	remote := []string{}
	if i := strings.Split(env["i"], ","); len(i) < 1 {
		if i := i[0]; i == "" {
		} else {
		}
	} else {
		for _, i := range i {
			if len(i) < 1 {
				continue
			}

			slashindex := strings.Index(i, "/")
			if slashindex < 0 {
				stdlib = append(stdlib, i)
			} else if ii := i[0:slashindex]; !strings.Contains(ii, ".") {
				stdlib = append(stdlib, i)
			} else {
				remote = append(remote, i)
			}
		}
	}

	sort.Strings(stdlib)
	sort.Strings(remote)

	tpl := template.Must(template.New("cache").Parse(CacheTemplate))
	data := CacheTemplateData{
		Package: env["p"],
		Key:     env["k"],
		Val:     env["v"],
		Stdlib:  stdlib,
		Remote:  remote,
	}
	if err := tpl.Execute(file, data); err != nil {
		fmt.Println("go-gengen: internal error", err)
	}
}

func sortImports(a, b string) bool {
	return false
}
