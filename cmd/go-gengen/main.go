package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	def := env()
	env := env().ParseFlags()

	for k := range env {
		if env[k] == def[k] {
			fmt.Println("go-gengen: not set", k)
			return
		}
	}

	outs := string(CacheTemplate)
	for k, v := range def {
		outs = strings.ReplaceAll(outs, v, env[k])
	}

	file, err := os.OpenFile("gen_cache.go", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("go-gengen: internal error", err)
		return
	}
	defer file.Close()

	file.WriteString(outs)
}
