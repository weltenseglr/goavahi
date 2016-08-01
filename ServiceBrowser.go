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
	path    dbus.ObjectPath
	addItem func(*ServiceBrowserItem)
	remItem func(*ServiceBrowserItem)
}

func (as *AvahiServer) ServiceBrowserNew(_if int32, proto int32, stype string, sdomain string, flags uint32) (*ServiceBrowser, error) {
	var path dbus.ObjectPath
	obj := as.conn.Object("org.freedesktop.Avahi", "/")
	err := obj.Call("org.freedesktop.Avahi.Server.ServiceBrowserNew", 0, _if, proto, stype, sdomain, flags).Store(&path)
	if err != nil {
		return nil, err
	}
	obj = as.conn.Object("org.freedesktop.Avahi", path)
	sb := ServiceBrowser{path, nil, nil}
	return &sb, nil
}

func (sb *ServiceBrowser) SetAddItemCallback(fn func(ServiceBrowserItem)) {
	Server.AddHandler(sb.path, "org.freedesktop.Avahi.ServiceBrowser.ItemNew", func(v *dbus.Signal) {
		fn(ServiceBrowserItem{
			v.Body[0].(int32),
			v.Body[1].(int32),
			v.Body[2].(string),
			v.Body[3].(string),
			v.Body[4].(string),
			v.Body[5].(uint32)})
	})
}

func (sb *ServiceBrowser) SetRemoveItemCallback(fn func(ServiceBrowserItem)) {
	Server.AddHandler(sb.path, "org.freedesktop.Avahi.ServiceBrowser.ItemRemove", func(v *dbus.Signal) {
		fn(ServiceBrowserItem{
			v.Body[0].(int32),
			v.Body[1].(int32),
			v.Body[2].(string),
			v.Body[3].(string),
			v.Body[4].(string),
			v.Body[5].(uint32)})
	})
}
