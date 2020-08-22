package dnsmsg

import (
	"testing"
)

func TestGetAuthorityServerName2ndNet(t *testing.T) {
	name := "vaddy.net"
	expect := "m.gtld-servers.net"
	result, err := getAuthorityServerName(name)
	if err != nil {
		t.Fail()
	}
	if result != expect {
		t.Fail()
	}
}
func TestGetAuthorityServerName3rdNet(t *testing.T) {
	name := "foo.vaddy.net"
	expect := "m.gtld-servers.net"
	result, err := getAuthorityServerName(name)
	if err != nil {
		t.Fail()
	}
	if result != expect {
		t.Fail()
	}
}

// not support 4th level domain
func TestGetAuthorityServerName4thNet(t *testing.T) {
	name := "foo.bar.vaddy.net"
	expect := ""
	result, err := getAuthorityServerName(name)
	if err == nil {
		t.Fail()
	}
	if result != expect {
		t.Fail()
	}
}

func TestGetAuthorityServerName2ndJp(t *testing.T) {
	name := "bitforest.jp"
	expect := "a.dns.jp"
	result, err := getAuthorityServerName(name)
	if err != nil {
		t.Fail()
	}
	if result != expect {
		t.Fail()
	}
}
func TestGetAuthorityServerName3rdJp(t *testing.T) {
	name := "foo.bitforest.jp"
	expect := "a.dns.jp"
	result, err := getAuthorityServerName(name)
	if err != nil {
		t.Fail()
	}
	if result != expect {
		t.Fail()
	}
}

// not support 4th level domain
func TestGetAuthorityServerName4thJp(t *testing.T) {
	name := "foo.bar.bitforest.jp"
	expect := ""
	result, err := getAuthorityServerName(name)
	if err == nil {
		t.Fail()
	}
	if result != expect {
		t.Fail()
	}
}

// no support domain
func TestGetAuthorityServerNameNoSupportDomain(t *testing.T) {
	name := "example.tokyo"
	expect := ""
	result, err := getAuthorityServerName(name)
	if err == nil {
		t.Fail()
	}
	if result != expect {
		t.Fail()
	}
}
