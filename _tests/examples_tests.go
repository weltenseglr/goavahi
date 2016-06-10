package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/weltenseglr/goavahi"
)

func onServiceAdd(i *goavahi.ServiceBrowserItem) {
	fmt.Printf("Found service %s\n", i)
}

func onServiceRem(i *goavahi.ServiceBrowserItem) {
	fmt.Printf("service disconnected %s\n", i)
}

func check(msg string, i interface{}, r error) {

}

func main() {
	s, err := goavahi.NewSimple()

	vs, err := s.GetVersionString()
	if err == nil {
		fmt.Println(vs)
	} else {
		fmt.Println(err.Error())
	}

	v, err := s.GetAPIVersion()
	if err == nil {
		fmt.Printf("API version: %d\n", v)
	} else {
		fmt.Println(err.Error())
	}

	str, err := s.GetHostName()
	if err == nil {
		fmt.Printf("Hostname: %s\n", str)
	} else {
		fmt.Println(err.Error())
	}

	str, err = s.GetHostNameFqdn()
	if err == nil {
		fmt.Printf("HostnameFqdn: %s\n", str)
	} else {
		fmt.Println(err.Error())
	}

	str, err = s.GetDomainName()
	if err == nil {
		fmt.Printf("Domain: %s\n", str)
	} else {
		fmt.Println(err.Error())
	}

	b, err := s.IsNSSSupportAvailable()
	if err == nil {
		fmt.Println("IsNSSSupportAvailable " + strconv.FormatBool(b))
	} else {
		fmt.Println(err.Error())
	}

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

	// wait a little
	time.Sleep(time.Second * 10)
}
