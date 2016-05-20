package goavahi

import (
	"fmt"

	"github.com/guelfey/go.dbus"
)

type AvahiServer struct {
	conn *dbus.Conn
	obj  *dbus.Object
}

func (as *AvahiServer) invoke(method string, args ...interface{}) *dbus.Call {
	method = "org.freedesktop.Avahi.Server." + method
	if args == nil {
		return as.obj.Call(method, 0)
	}
	return as.obj.Call(method, 0, args)

}

func (as *AvahiServer) GetVersionString() (string, error) {
	var vs string
	err := as.invoke("GetVersionString").Store(&vs)
	return vs, err
}

func (as *AvahiServer) GetAPIVersion() (uint32, error) {
	var v uint32
	err := as.invoke("GetAPIVersion").Store(&v)
	return v, err
}

func (as *AvahiServer) GetHostName() (string, error) {
	var hostname string
	err := as.invoke("GetHostName").Store(&hostname)
	return hostname, err
}

func (as *AvahiServer) SetHostName(hostname string) error {
	return as.invoke("SetHostName", hostname).Err
}

func (as *AvahiServer) GetHostNameFqdn() (string, error) {
	var fqdn string
	err := as.invoke("GetHostNameFqdn").Store(&fqdn)
	return fqdn, err
}

func (as *AvahiServer) GetDomainName() (string, error) {
	var domain string
	err := as.invoke("GetDomainName").Store(&domain)
	return domain, err
}

func (as *AvahiServer) IsNSSSupportAvailable() (bool, error) {
	var b bool
	err := as.invoke("IsNSSSupportAvailable").Store(&b)
	fmt.Println(b)
	return b, err
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
