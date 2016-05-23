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

func (as *AvahiServer) GetState() (int32, error) {
	var i int32
	err := as.invoke("GetState").Store(&i)
	return i, err
}

func (as *AvahiServer) GetLocalServiceCookie() (uint32, error) {
	var c uint32
	err := as.invoke("GetLocalServiceCookie").Store(&c)
	return c, err
}

func (as *AvahiServer) GetAlternativeHostName(name string) (string, error) {
	var aname string
	err := as.invoke("GetAlternativeHostName", name).Store(&aname)
	return aname, err
}

func (as *AvahiServer) GetAlternativeServiceName(name string) (string, error) {
	var aname string
	err := as.invoke("GetAlternativeServiceName", name).Store(&aname)
	return aname, err
}

func (as *AvahiServer) GetNetworkInterfaceNameByIndex(i int) (string, error) {
	var name string
	err := as.invoke("GetNetworkInterfaceNameByIndex", i).Store(&name)
	return name, err
}

func (as *AvahiServer) GetNetworkInterfaceIndexByName(name string) (int, error) {
	var i int
	err := as.invoke("GetNetworkInterfaceIndexByName", name).Store(&i)
	return i, err
}

func (as *AvahiServer) ResolveHostName(_interface, protocol int32, name string, aprotocol int32, flags uint32) (error, int32, int32, string, int32, string, uint32) {
	var _if, proto, aproto int32
	var addr string
	err := as.invoke("ResolveHostNameResolveHostName", _interface, protocol, name, aprotocol, flags).Store(&_if, &proto, &name, &aproto, &addr, &flags)
	return err, _if, proto, name, proto, addr, flags
}

func (as *AvahiServer) ResolveAddress(_interface, protocol int32, address string, flags uint32) (error, int32, int32, int32, string, string, uint32) {
	var aproto int32
	var name string
	err := as.invoke("ResolveAddress", _interface, protocol, address, flags).Store(&_interface, &protocol, &aproto, &address, &name, &flags)
	return err, _interface, protocol, aproto, address, name, flags
}

func (as *AvahiServer) ResolveService(_interface, protocol int32, name, stype, domain string, aprotocol int32, flags uint32) (error, int32, int32, string, string, string, string, int32, string, uint16, [][]byte, uint32) {
	var host, address string
	var port uint16
	var txt [][]byte
	err := as.invoke("ResolveService", _interface, protocol, name, stype, domain, aprotocol, flags).Store(&_interface, &protocol, &name, &stype, &domain, &host, &aprotocol, &address, &port, &txt, &flags)
	return err, _interface, protocol, name, stype, domain, host, aprotocol, address, port, txt, flags
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
	err := obj.Call("org.freedesktop.Avahi.Server.ServiceBrowserNew", 0, _if, proto, stype, sdomain, flags).Store(&path)
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
