package tests

import (
	"testing"

	"github.com/pixel365/goepp"
)

func TestValidContactUpdateAddStatus(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/update/contact/valid_add_status.xml")
}

func TestValidContactUpdateRemStatus(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/update/contact/valid_rem_status.xml")
}

func TestValidContactUpdateChgEmail(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/update/contact/valid_chg_email.xml")
}

func TestValidContactUpdateChgAuthInfo(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/update/contact/valid_chg_authinfo.xml")
}

func TestValidContactUpdateChgDisclose(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/update/contact/valid_chg_disclose.xml")
}

func TestValidContactUpdateChgPostalInfo(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/update/contact/valid_chg_postalinfo.xml")
}

func TestValidContactUpdateAll(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/update/contact/valid_add_rem_chg.xml")
}

func TestInvalidContactUpdateNoObject(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/contact/invalid_no_object.xml")
}

func TestInvalidContactUpdateWrongNamespace(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/contact/invalid_wrong_namespace.xml")
}

func TestInvalidContactUpdateTwoObjects(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/contact/invalid_two_objects.xml")
}

func TestInvalidContactUpdateMissingID(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/contact/invalid_missing_id.xml")
}

func TestInvalidContactUpdateNoAddRemChg(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/contact/invalid_no_add_rem_chg.xml")
}

func TestInvalidContactUpdateAddEmpty(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/contact/invalid_add_empty.xml")
}

func TestInvalidContactUpdateAddStatusEmpty(t *testing.T) {
	mustFailParse(
		t,
		&goepp.CmdParser{},
		"testdata/update/contact/invalid_add_status_empty.xml",
	)
}

func TestInvalidContactUpdateRemEmpty(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/contact/invalid_rem_empty.xml")
}

func TestInvalidContactUpdateChgEmpty(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/update/contact/invalid_chg_empty.xml")
}

func TestInvalidContactUpdateChgEmailEmpty(t *testing.T) {
	mustFailParse(
		t,
		&goepp.CmdParser{},
		"testdata/update/contact/invalid_chg_email_empty.xml",
	)
}

func TestInvalidContactUpdateChgAuthInfoEmpty(t *testing.T) {
	mustFailParse(
		t,
		&goepp.CmdParser{},
		"testdata/update/contact/invalid_chg_authinfo_empty.xml",
	)
}

func TestInvalidContactUpdateChgAuthInfoPwAndNull(t *testing.T) {
	mustFailParse(
		t,
		&goepp.CmdParser{},
		"testdata/update/contact/invalid_chg_authinfo_pw_and_null.xml",
	)
}

func TestInvalidContactUpdateChgDiscloseEmpty(t *testing.T) {
	mustFailParse(
		t,
		&goepp.CmdParser{},
		"testdata/update/contact/invalid_chg_disclose_empty.xml",
	)
}

func TestInvalidContactUpdateChgPostalInfoBadType(t *testing.T) {
	mustFailParse(
		t,
		&goepp.CmdParser{},
		"testdata/update/contact/invalid_chg_postalinfo_bad_type.xml",
	)
}
