package main

import (
	"bufio"
	"encoding/gob"
	"flag"
	"fmt"
	"net"
	"os"
)

var fPath string
var h bool

func usage() {
	fmt.Fprintf(os.Stderr, "hubs_dump.go -f [RAW_DATA_FILE]")
}

func init() {
	flag.BoolVar(&h, "h", false, "Print this help")
	flag.StringVar(&fPath, "f", "filepath", "Path of target raw data file.")
	flag.Usage = usage
}

type SeedNode struct {
	IP   net.IP
	Port int64
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}

	fd, fdErr := os.OpenFile(fPath, os.O_RDONLY, 0644)
	if fdErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to OpenFile([%s])\n%+v", fPath, fdErr)
		os.Exit(1)
	}
	defer fd.Close()

	br := bufio.NewReader(fd)

	dec := gob.NewDecoder(br)
	var d []SeedNode
	decErr := dec.Decode(&d)
	if decErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to Decode raw data.\n%+v", decErr)
		os.Exit(1)
	}

	for _, seedNode := range d {
		fmt.Printf("%s:%d\n", seedNode.IP, seedNode.Port)
	}
}
