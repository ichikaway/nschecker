package dnsmsg

import (
	"regexp"
	"testing"
)

func TestGetAuthorityServerName2ndNet(t *testing.T) {
	name := "vaddy.net"
	expect := regexp.MustCompile(`[a-m].gtld-servers.net.`)
	result, err := getAuthorityServerName(name)
	if err != nil {
		t.Fail()
	}
	if !expect.MatchString(result[0]) {
		t.Error(result)
		t.Fail()
	}
}
func TestGetAuthorityServerName3rdNet(t *testing.T) {
	name := "foo.vaddy.net"
	expect := regexp.MustCompile(`[a-m].gtld-servers.net.`)
	result, err := getAuthorityServerName(name)
	if err != nil {
		t.Fail()
	}
	if !expect.MatchString(result[0]) {
		t.Error(result)
		t.Fail()
	}
}

func TestGetAuthorityServerName2ndJp(t *testing.T) {
	name := "bitforest.jp"
	expect := regexp.MustCompile(`[a-h].dns.jp.`)
	result, err := getAuthorityServerName(name)
	if err != nil {
		t.Fail()
	}
	if !expect.MatchString(result[0]) {
		t.Error(result)
		t.Fail()
	}
}
func TestGetAuthorityServerName3rdJp(t *testing.T) {
	name := "foo.bitforest.jp"
	expect := regexp.MustCompile(`[a-h].dns.jp.`)
	result, err := getAuthorityServerName(name)
	if err != nil {
		t.Fail()
	}
	if !expect.MatchString(result[0]) {
		t.Error(result)
		t.Fail()
	}
}
