package checker

import "testing"

func TestInArray(t *testing.T) {
	var needle string = "vaddy.net"
	var haystack = []string{"vaddy.net", "hoge.vaddy.net"}

	if !in_array(needle, haystack) {
		t.Fail()
	}
}

func TestInArrayCaseInsensitive(t *testing.T) {
	var needle string = "VADDY.NET"
	var haystack = []string{"vaddy.net", "hoge.vaddy.net"}

	if !in_array(needle, haystack) {
		t.Fail()
	}
}

func TestInArrayCaseInsensitiveInHaystack(t *testing.T) {
	var needle string = "vaddy.net"
	var haystack = []string{"VADDY.NET", "HOGE.VADDY.NET"}

	if !in_array(needle, haystack) {
		t.Fail()
	}
}

func TestInArrayNotFound(t *testing.T) {
	var needle string = "foo"
	var haystack = []string{"vaddy.net", "hoge.vaddy.net"}

	if in_array(needle, haystack) {
		t.Fail()
	}
}

func TestCheckRecordNs(t *testing.T) {
	var domain string = "vaddy.net"
	var expect string = "ns-1151.awsdns-15.org. , ns-1908.awsdns-46.co.uk. , ns-457.awsdns-57.com. , ns-700.awsdns-23.net."

	err := CheckRecord("NS", domain, expect)
	if err != nil {
		t.Fail()
	}
}

func TestCheckRecordNsCaseInsensitiveCheck(t *testing.T) {
	var domain string = "vaddy.net"
	var expect string = "NS-1151.AWSDNS-15.ORG. , ns-1908.awsdns-46.co.uk. , ns-457.awsdns-57.com. , ns-700.awsdns-23.net."

	err := CheckRecord("NS", domain, expect)
	if err != nil {
		t.Fail()
	}
}

func TestCheckRecordNsJp(t *testing.T) {
	var domain string = "bitforest.jp"
	var expect string = "ns-722.awsdns-26.net. , ns-254.awsdns-31.com. , ns-1740.awsdns-25.co.uk. , ns-1319.awsdns-36.org."

	err := CheckRecord("NS", domain, expect)
	if err != nil {
		t.Fail()
	}
}

func TestCheckRecordNsJpNoLastPeriod(t *testing.T) {
	var domain string = "bitforest.jp"
	var expect string = "ns-722.awsdns-26.net , ns-254.awsdns-31.com , ns-1740.awsdns-25.co.uk , ns-1319.awsdns-36.org"

	err := CheckRecord("NS", domain, expect)
	if err != nil {
		t.Fail()
	}
}

func TestCheckRecordNsFail(t *testing.T) {
	var domain string = "vaddy.net"
	var expect string = "nstest.example.com , nstest2.example.com"

	err := CheckRecord("NS", domain, expect)
	if err == nil {
		t.Fail()
	}
}

func TestCheckRecordMx(t *testing.T) {
	var domain string = "vaddy.net"
	var expect string = "aspmx.l.google.com. , alt1.aspmx.l.google.com. , alt2.aspmx.l.google.com. , alt4.aspmx.l.google.com. , alt3.aspmx.l.google.com."

	err := CheckRecord("MX", domain, expect)
	if err != nil {
		t.Fail()
	}
}

func TestCheckRecordMxFail(t *testing.T) {
	var domain string = "vaddy.net"
	var expect string = "mxtest.example.com , mxtest2.example.com"

	err := CheckRecord("MX", domain, expect)
	if err == nil {
		t.Fail()
	}
}
