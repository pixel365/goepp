package tests

import (
	"testing"

	"github.com/pixel365/goepp"
)

func TestValidHostUpdateAddAddr(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/update/host/valid_add_addr.xml")
}

func TestValidHostUpdateRemStatus(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/update/host/valid_rem_status.xml")
}

func TestValidHostUpdateChgName(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/update/host/valid_chg_name.xml")
}

func TestValidHostUpdateAll(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/update/host/valid_add_rem_chg.xml")
}

func TestInvalidHostUpdateNoObject(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_no_object.xml")
}

func TestInvalidHostUpdateWrongNamespace(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_wrong_namespace.xml")
}

func TestInvalidHostUpdateTwoObjects(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_two_objects.xml")
}

func TestInvalidHostUpdateMissingName(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_missing_name.xml")
}

func TestInvalidHostUpdateNoAddRemChg(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_no_add_rem_chg.xml")
}

func TestInvalidHostUpdateAddEmpty(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_add_empty.xml")
}

func TestInvalidHostUpdateAddAddrEmpty(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_add_addr_empty.xml")
}

func TestInvalidHostUpdateAddAddrBadIP(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_add_addr_bad_ip.xml")
}

func TestInvalidHostUpdateAddStatusEmpty(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_add_status_empty.xml")
}

func TestInvalidHostUpdateRemEmpty(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_rem_empty.xml")
}

func TestInvalidHostUpdateChgEmpty(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_chg_empty.xml")
}

func TestInvalidHostUpdateChgNameEmpty(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/host/invalid_chg_name_empty.xml")
}
