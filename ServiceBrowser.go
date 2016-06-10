package goavahi

import (
	"github.com/guelfey/go.dbus"
)

type ServiceBrowserItem struct {
	Interface int32
	Protocol  int32
	Name      string
	Type      string
	Domain    string
	Flags     uint32
}

type ServiceBrowser struct {
	conn    *dbus.Conn
	obj     *dbus.Object
	addItem func(*ServiceBrowserItem)
	remItem func(*ServiceBrowserItem)
}

func (sb *ServiceBrowser) Start() {
	c := make(chan *dbus.Signal, 10)
	sb.obj.Call("org.freedesktop.Avahi.ServiceBrowser", 0)
	sb.conn.Signal(c)
	for v := range c {
		switch v.Name {
		case "org.freedesktop.Avahi.ServiceBrowser.ItemNew":
			if sb.addItem != nil {
				sb.addItem(&ServiceBrowserItem{
					v.Body[0].(int32),
					v.Body[1].(int32),
					v.Body[2].(string),
					v.Body[3].(string),
					v.Body[4].(string),
					v.Body[5].(uint32)})
			}
			break
		case "org.freedesktop.Avahi.ServiceBrowser.ItemRemove":
			if sb.remItem != nil {
				sb.remItem(&ServiceBrowserItem{
					v.Body[0].(int32),
					v.Body[1].(int32),
					v.Body[2].(string),
					v.Body[3].(string),
					v.Body[4].(string),
					v.Body[5].(uint32)})
			}
			break
		}
	}
	//"type='signal'"
}

func (sb *ServiceBrowser) SetAddItemCallback(fn func(*ServiceBrowserItem)) {
	sb.addItem = fn
}

func (sb *ServiceBrowser) SetRemoveItemCallback(fn func(*ServiceBrowserItem)) {
	sb.remItem = fn
}
