package goavahi

import (
	"errors"
	"strings"
)

type Simple struct {
	as *AvahiServer
	eg *EntryGroup
	sb *ServiceBrowser
}

func NewSimple() (*Simple, error) {
	as, err := Connect()
	if err != nil {
		return nil, err
	}
	return &Simple{as, nil, nil}, nil
}

/***************************************
			EntryGroup Wrappers
 ***************************************/

func (s *Simple) getEntryGroup() error {
	var err error
	if s.eg == nil {
		s.eg, err = s.as.EntryGroupNew()
	}
	return err
}

func (s *Simple) AddService(sname string, stype string, port uint16, txtRecords map[string]string) error {
	if err := s.getEntryGroup(); err != nil {
		return err
	}
	return s.eg.AddService(
		int32(-1), // avahi.IF_UNSPEC
		int32(-1), // avahi.PROTO_UNSPEC
		uint32(0), // flags
		sname,
		stype,
		"", // sdomain let avahi decide
		"", // shost: let avahi decide
		port,
		txtRecords)
}

func (s *Simple) AddServiceSubtype(sname string, stype string, substype string, port uint16, txtRecords map[string]string) error {
	if !strings.HasSuffix(substype, stype) {
		return errors.New("subtype must contain type! E.g. subtype '_my._sub._http._tcp' and type '_http._tcp'. Was :" + substype + " -- " + stype)
	}
	if err := s.getEntryGroup(); err != nil {
		return err
	}
	return s.eg.AddServiceSubtype(
		int32(-1), // avahi.IF_UNSPEC
		int32(-1), // avahi.PROTO_UNSPEC
		uint32(0), // flags
		sname,
		stype,
		"", // sdomain let avahi decide
		substype,
		"",   // shost: let avahi decide
		port, // port
		txtRecords)
}

func (s *Simple) EntryGroupCommit() error {
	return s.eg.Commit()
}

/***************************************
		ServiceBrowser Wrappers
 ***************************************/

func (s *Simple) getServiceBrowser(stype string) error {
	var err error
	if s.sb == nil {
		s.sb, err = s.as.ServiceBrowserNew(
			-1, // avahi.IF_UNSPEC
			-1, // avahi.PROTO_UNSPEC
			stype,
			"", // domain: let avahi decide
			0)  // no flags
	}
	return err
}

func (s *Simple) BrowseServices(stype string, onAdd func(*ServiceBrowserItem), onRemove func(*ServiceBrowserItem)) error {
	if err := s.getServiceBrowser(stype); err != nil {
		return err
	}
	s.sb.SetAddItemCallback(onAdd)
	s.sb.SetRemoveItemCallback(onRemove)
	go s.sb.Start()
	return nil
}
