package update

import (
	"errors"
	"strings"

	"github.com/pixel365/goepp/command"
	"github.com/pixel365/goepp/command/create"
)

type HostData struct {
	Add    *HostAdd    `xml:"add,omitempty"`
	Remove *HostRemove `xml:"rem,omitempty"`
	Change *HostChange `xml:"chg,omitempty"`
	command.HostRef
}

type HostAdd struct {
	Statuses []HostStatus `xml:"status,omitempty"`
	Addrs    []HostAddr   `xml:"addr,omitempty"`
}

type HostRemove struct {
	Statuses []HostStatus `xml:"status,omitempty"`
	Addrs    []HostAddr   `xml:"addr,omitempty"`
}

type HostChange struct {
	Name string `xml:"name"`
}

type HostStatus struct {
	Value string `xml:"s,attr"`
}

type HostAddr struct {
	IP    create.HostAddrType `xml:"ip,attr,omitempty"`
	Value string              `xml:",chardata"`
}

func (h *HostData) Validate() error {
	if strings.TrimSpace(h.Name) == "" {
		return errors.New("host:name is required")
	}

	if h.Add == nil && h.Remove == nil && h.Change == nil {
		return errors.New("host:add, rem or chg is required")
	}

	if h.Add != nil {
		if err := h.Add.Validate(); err != nil {
			return err
		}
	}

	if h.Remove != nil {
		if err := h.Remove.Validate(); err != nil {
			return err
		}
	}

	if h.Change != nil {
		if err := h.Change.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (a *HostAdd) Validate() error {
	if len(a.Statuses) == 0 && len(a.Addrs) == 0 {
		return errors.New("add:status or addr is required")
	}

	if err := validateHostStatuses(a.Statuses); err != nil {
		return err
	}

	if err := validateHostAddrs(a.Addrs); err != nil {
		return err
	}

	return nil
}

func (r *HostRemove) Validate() error {
	if len(r.Statuses) == 0 && len(r.Addrs) == 0 {
		return errors.New("rem:status or addr is required")
	}

	if err := validateHostStatuses(r.Statuses); err != nil {
		return err
	}

	if err := validateHostAddrs(r.Addrs); err != nil {
		return err
	}

	return nil
}

func (c *HostChange) Validate() error {
	if strings.TrimSpace(c.Name) == "" {
		return errors.New("chg:name is required")
	}

	return nil
}

func validateHostStatuses(statuses []HostStatus) error {
	for _, status := range statuses {
		if strings.TrimSpace(status.Value) == "" {
			return errors.New("status value is required")
		}
	}

	return nil
}

func validateHostAddrs(addrs []HostAddr) error {
	for _, addr := range addrs {
		if strings.TrimSpace(addr.Value) == "" {
			return errors.New("addr value is required")
		}

		if addr.IP != "" && addr.IP != create.HostAddrTypeIPv4 &&
			addr.IP != create.HostAddrTypeIPv6 {
			return errors.New("addr ip must be v4 or v6")
		}
	}

	return nil
}
