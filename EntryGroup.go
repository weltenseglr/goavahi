package goavahi

import (
	"github.com/guelfey/go.dbus"
)

type EntryGroup struct {
	obj *dbus.Object
}

func (a *EntryGroup) AddService(_if int32, proto int32, flags uint32, sname string, stype string, sdomain string, shost string, port uint16, txtRecords map[string]string) error {
	txt := renderRecord(txtRecords)
	dc := a.obj.Call("org.freedesktop.Avahi.EntryGroup.AddService", 0,
		_if,
		proto,
		flags,
		sname,
		stype,
		sdomain,
		shost,
		port,
		txt)
	return dc.Err
}

func (a *EntryGroup) AddServiceSubtype(_if int32, proto int32, flags uint32, sname string, stype string, sdomain string, substype string, shost string, port uint16, txtRecords map[string]string) error {
	txt := renderRecord(txtRecords)
	dc := a.obj.Call("org.freedesktop.Avahi.EntryGroup.AddServiceSubtype", 0,
		_if,
		proto,
		flags,
		sname,
		stype,
		sdomain,
		substype,
		shost,
		port,
		txt)
	return dc.Err
}

func (e *EntryGroup) Commit() error {
	call := e.obj.Call("org.freedesktop.Avahi.EntryGroup.Commit", 0)
	return call.Err
}
