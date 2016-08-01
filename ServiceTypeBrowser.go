package goavahi

import (
	"github.com/guelfey/go.dbus"
)

type ServiceTypeBrowserItem struct {
	Interface int32
	Protocol  int32
	Stype     string
	Domain    string
	Flags     uint32
}
type ServiceTypeBrowser struct {
	conn    *dbus.Conn
	obj     *dbus.Object
	addItem func(ServiceTypeBrowserItem)
	remItem func(ServiceTypeBrowserItem)
}

func (stb *ServiceTypeBrowser) Start() {
	c := make(chan *dbus.Signal, 10)
	stb.conn.Signal(c)
	stb.obj.Call("org.freedesktop.Avahi.ServiceTypeBrowser", 0)
	for v := range c {
		switch v.Name {
		case "org.freedesktop.Avahi.ServiceTypeBrowser.ItemNew":
			if stb.addItem != nil {
				stb.addItem(ServiceTypeBrowserItem{
					v.Body[0].(int32),
					v.Body[1].(int32),
					v.Body[2].(string),
					v.Body[3].(string),
					v.Body[4].(uint32)})
			}
			break
		case "org.freedesktop.Avahi.ServiceTypeBrowser.ItemRemove":
			if stb.remItem != nil {
				stb.remItem(ServiceTypeBrowserItem{
					v.Body[0].(int32),
					v.Body[1].(int32),
					v.Body[2].(string),
					v.Body[3].(string),
					v.Body[4].(uint32)})
			}
			break
		}
	}
}

func (stb *ServiceTypeBrowser) SetAddItemCallback(fn func(ServiceTypeBrowserItem)) {
	stb.addItem = fn
}

func (stb *ServiceTypeBrowser) SetRemoveItemCallback(fn func(ServiceTypeBrowserItem)) {
	stb.remItem = fn
}
