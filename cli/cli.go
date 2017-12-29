package main

import (
	"github.com/LinJinghua/hrc/client"

	flag "github.com/spf13/pflag"
)

func cli()  {
	m := flag.StringP("method", "m", "sync", "Method for Reactive Client[sync | async]")
	flag.Parse()
	if len(*m) != 0 && *m == "async" {
		client.NewClient().GetServiceAsync()
	} else {
		client.NewClient().GetService()
	}
}

func main() {
	cli()
}
