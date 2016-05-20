package goavahi

import (
	"fmt"

	"github.com/guelfey/go.dbus"
)

type Avahi_dbus struct {
	connection *dbus.Conn
	obj        *dbus.Object
}

func GetServer(conn *dbus.Conn) (*AvahiServer, error) {
	r := AvahiServer{conn, conn.Object("org.freedesktop.Avahi", "/")}
	return &r, nil
}

func Connect() (*AvahiServer, error) {
	dconn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	r, err := GetServer(dconn)
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

func renderRecord(r map[string]string) [][]byte {
	var txt [][]byte
	for k, v := range r {
		txt = append(txt, []byte(k+"="+v))
	}
	return txt
}
