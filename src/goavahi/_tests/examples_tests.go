package main

import (
	"fmt"
	"time"

	"github.com/weltenseglr/goavahi"
)

func onServiceAdd(i *goavahi.ServiceBrowserItem) {
	fmt.Printf("Found service %s\n", i)
}

func onServiceRem(i *goavahi.ServiceBrowserItem) {
	fmt.Printf("service disconnected %s\n", i)
}

func main() {
	s, err := goavahi.NewSimple()

	fmt.Println("create service browser")
	err = s.BrowseServices("_foo._tcp", onServiceAdd, onServiceRem)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	txt := make(map[string]string, 2)
	txt["FOO"] = "BAR"
	txt["USR"] = "weltenseglr"

	err = s.AddService("test", "_foo._tcp", 9999, txt)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = s.AddServiceSubtype("subtest", "_foo._tcp", "_bar._sub._foo._tcp", 9999, txt)
	if err != nil {
		fmt.Println(err.Error())
	}
	s.EntryGroupCommit()

}
