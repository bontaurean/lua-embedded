package main

import (
	"fmt"
	"lua-embedded/handlers"

	lua "github.com/yuin/gopher-lua"
)

func main() {
	l := lua.NewState()
	defer l.Close()

	l.PreloadModule("handlers", handlers.Configure)

	if err := l.DoString(`
		local m = require("handlers")
		cpuCount = m.load("cpus")
		cpuCount = m.process(cpuCount)
		cpuCount = m.process(cpuCount)
		cpuCount = m.process(cpuCount)
		m.store("cpus", cpuCount)

		assert(m.load("cpus")==3)
	`); err != nil {
		fmt.Println(err)
	}
}
