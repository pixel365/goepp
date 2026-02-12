package tests

import (
	"testing"

	"github.com/pixel365/goepp"
)

func TestValidOpReq(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/poll/valid_req.xml")
}

func TestValidOpAck(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/poll/valid_ack.xml")
}

func TestEmptyOp(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/poll/empty_op.xml")
}

func TestInvalidOp(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/poll/invalid_op.xml")
}
