package goavahi

import (
	"fmt"
	"log"

	"github.com/guelfey/go.dbus"
)

type DomainBrowser struct{}
type ServiceTypeBrowser struct{}
type ServiceResolver struct{}
type HostNameResolver struct{}
type AddressResolver struct{}
type RecordBrowser struct{}

type Avahi_dbus struct {
	connection *dbus.Conn
	obj        *dbus.Object
}

func Server(conn *dbus.Conn) (*AvahiServer, error) {
	r := AvahiServer{conn, conn.Object("org.freedesktop.Avahi", "/")}
	return &r, nil
}

func Connect() (*AvahiServer, error) {
	dconn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	r, err := Server(dconn)
	return r, err
}

func Dbus_Test() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	ch := make(chan *dbus.Call, 10)
	conn.BusObject().Go("org.freedesktop.DBus.ListNames", 0, ch)
	select {
	case call := <-ch:
		if call.Err != nil {
			panic(err)
		}
		list := call.Body[0].([]string)
		for _, v := range list {
			fmt.Println(v)
		}
	}
}

func (a *Avahi_dbus) BrowseAll() error {
	log.Println("Avahi Browser called.")
	/*
		out i interface
		out i protocol
		out s name
		out s type
		out s domain
		out s host
		out i aprotocol
		out s address
		out q port
		out aay txt
		out u flags
	*/

	dconn, err := dbus.SystemBus()
	if err != nil {
		return err
	}
	var path dbus.ObjectPath
	obj := dconn.Object("org.freedesktop.Avahi", "/")
	err = obj.Call("org.freedesktop.Avahi.Server.ServiceBrowserNew", 0,
		int32(-1), // avahi.IF_UNSPEC
		int32(-1),
		"_http._tcp",
		"",
		uint32(0)).Store(&path)
	if err != nil {
		return err
	}
	serviceBrowserObject := dconn.Object("org.freedesktop.Avahi", path)

	log.Printf("Data: %s\nObject: %s", path, serviceBrowserObject)
	return nil
}

func renderRecord(r map[string]string) [][]byte {
	var txt [][]byte
	for k, v := range r {
		txt = append(txt, []byte(k+"="+v))
	}
	return txt
}
