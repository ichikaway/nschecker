package dnsmsg

type A struct {
	Int32Field int32
	ByteField  byte
}

/*
func TestConnect(t *testing.T) {
	//Send([]byte("Hello from Client"))
	//Send([]byte{0xe3, 0x81, 0x82})
	//a := A{Int32Field: 0x123456, ByteField: 0xFF}
	//a := A{Int32Field: 0xe38182, ByteField: 0xFF}

	//a := try("www.vaddy.net")
	a := try("a.net")
	fmt.Printf("%v\n",a)
	b := bytes.NewBuffer([]byte(""))

	// 構造体をバイナリにする
	binary.Write(b, binary.BigEndian, &a)
	//binary.Write(b, binary.LittleEndian, &a)
	fmt.Printf("%x\n",b.Bytes())
	Send(b.Bytes())
}
*/
/*
func TestConnectJp(t *testing.T) {
	id := uint16(60000)
	h := NewHeader(id)
	header := NewHeaderPayload(h)
	q := NewQuestion("bitforest.jp", TypeNS)
	question := NewQuestionPayload(q)
	req := append(header, question...)
	buf := Send("a.dns.jp:53", req)

	var m dnsmessage.Message
	m.Unpack(buf)
	//fmt.Print(m)
	rb := m.Authorities[0].Body.(*dnsmessage.NSResource)
	//rb := m.Answers[0].Body.(*dnsmessage.NSResource)
	fmt.Println(rb.NS.String())
}
*/
/*
func TestConnectNet(t *testing.T) {
	id := uint16(60000)
	h := NewHeader(id)
	header := NewHeaderPayload(h)
	q := NewQuestion("vaddy.net", TypeNS)
	question := NewQuestionPayload(q)
	req := append(header, question...)
	buf := Send("m.gtld-servers.net:53", req)
	//buf := Send("ns-700.awsdns-23.net:53", req)

	var m dnsmessage.Message
	m.Unpack(buf)
	//a := m.Answers[0]
	//var rb dnsmessage.NSResource = a
	//rb := m.Authorities[0].Body.(*dnsmessage.NSResource)
	if len(m.Authorities) == 0 {
		//error
	}

	var ret []string
	for _, authrotiy := range m.Authorities {
		rr := authrotiy.Body.(*dnsmessage.NSResource)
		ret = append(ret, rr.NS.String())
	}

	fmt.Println(ret)

	//rb := m.Answers[0].Body.(*dnsmessage.NSResource)
	//fmt.Println(rb.NS.String())

	//fmt.Printf("%x", m.Answers[0].Body.(*dnsmessage.NSResource))
	//fmt.Printf("%s", m.Answers[0].Body)
	//fmt.Printf("%+v", m.Answers[0].Body)
	//fmt.Printf("%#v", m.Answers[0].Body)
	//fmt.Print(m.Answers[0].Body)
	//fmt.Print(m.Answers[1].Body)
	//fmt.Print(m.Answers[2].Body)
	//fmt.Print(m.Answers[3].Body)
}
*/
