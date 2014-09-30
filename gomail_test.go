package gomail

import (
	"testing"
)

func TestImports(t *testing.T) {
	gomail := New()
	gomail.sendMail("subject", "body")
}
