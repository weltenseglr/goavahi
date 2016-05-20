package goavahi

import (
	"github.com/guelfey/go.dbus"
)

type AvahiServer struct {
	conn *dbus.Conn
	obj  *dbus.Object
}

func (as *AvahiServer) GetVersionString() {

}

func (as *AvahiServer) GetAPIVersion() {

}

func (as *AvahiServer) GetHostName() {

}

func (as *AvahiServer) SetHostName() {

}

func (as *AvahiServer) GetHostNameFqdn() {

}

func (as *AvahiServer) GetDomainName() {

}

func (as *AvahiServer) IsNSSSupportAvailable() {

}

func (as *AvahiServer) GetState() {

}

func (as *AvahiServer) GetLocalServiceCookie() {

}

func (as *AvahiServer) GetAlternativeHostName() {

}

func (as *AvahiServer) GetAlternativeServiceName() {

}

func (as *AvahiServer) GetNetworkInterfaceNameByIndex() {

}

func (as *AvahiServer) GetNetworkInterfaceIndexByName() {

}

func (as *AvahiServer) ResolveHostName() {

}

func (as *AvahiServer) ResolveAddress() {

}

func (as *AvahiServer) ResolveService() {

}

func (as *AvahiServer) EntryGroupNew() (*EntryGroup, error) {

	var path dbus.ObjectPath
	err := as.obj.Call("org.freedesktop.Avahi.Server.EntryGroupNew", 0).Store(&path)
	if err != nil {
		return nil, err
	}
	obj := as.conn.Object("org.freedesktop.Avahi", path)
	return &EntryGroup{obj}, nil
}

func (as *AvahiServer) DomainBrowserNew() {

}

func (as *AvahiServer) ServiceTypeBrowserNew() {

}

func (as *AvahiServer) ServiceBrowserNew(_if int32, proto int32, stype string, sdomain string, flags uint32) (*ServiceBrowser, error) {
	var path dbus.ObjectPath
	obj := as.conn.Object("org.freedesktop.Avahi", "/")
	err := obj.Call("org.freedesktop.Avahi.Server.ServiceBrowserNew", 0,
		_if,
		proto,
		stype,
		sdomain,
		flags).Store(&path)
	if err != nil {
		return nil, err
	}
	obj = as.conn.Object("org.freedesktop.Avahi", path)
	return &ServiceBrowser{as.conn, obj, nil, nil}, nil
}

func (as *AvahiServer) ServiceResolverNew() {

}

func (as *AvahiServer) HostNameResolverNew() {

}

func (as *AvahiServer) AddressResolverNew() {

}

func (as *AvahiServer) RecordBrowserNew() {

}
