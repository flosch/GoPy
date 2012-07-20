package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/flosch/GoPy/src/vm"

	"os"
	"runtime/pprof"
)

const (
	VersionMajor = 0
	VersionMinor = 1
)

var Version = fmt.Sprintf("%d.%d", VersionMajor, VersionMinor)

func main() {
	file, err := os.Create("perf.az")
	if err != nil {
		panic(err)
	}
	if err := pprof.StartCPUProfile(file); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	fmt.Printf("GoPy %s - (C) 2012 Florian Schlachter, Berlin\n\n", Version)

	debug := flag.Bool("debug", false, "Enables debug mode")
	filename := flag.String("filename", "", "Compiled python file (.pyc)")
	flag.Parse()

	if *filename == "" {
		log.Println("Please provide a filename.")
		flag.Usage()
		return
	}

	log.Printf("[Settings filename=%s]\n\n", *filename)

	vm, err := vm.NewVM(*filename, *debug)
	if err != nil {
		log.Fatal("Error during creating VM: ", err.Error())
	}

	log.Printf("VM created [filename=%s,name=%s]\n", *vm.Filename(), *vm.Name())

	if err := vm.Run(); err != nil {
		log.Fatal("Error during execution: ", err.Error())
	}
}
