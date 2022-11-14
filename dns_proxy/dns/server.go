package dns

/*
dig命令主要用来从 DNS 域名服务器查询主机地址信息。

查找www.baidu.com的ip （A记录）：
命令：dig @127.0.0.1 www.baidu.com
根据ip查找对应域名 （PTR记录）：
命令：dig @127.0.0.1 -x 220.181.38.150

*/

import (
	"fmt"
	"net"

	"golang.org/x/net/dns/dnsmessage"
)

func Serve() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 53})
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Listing ...")
	for {
		buf := make([]byte, 512)
		_, addr, _ := conn.ReadFromUDP(buf)
		var msg dnsmessage.Message
		if err := msg.Unpack(buf); err != nil {
			fmt.Println(err)
			continue
		}
		go ServerDNS(addr, conn, msg)
	}
}

// ServerDNS serve
func ServerDNS(addr *net.UDPAddr, conn *net.UDPConn, msg dnsmessage.Message) {
	// query info
	if len(msg.Questions) < 1 {
		return
	}
	question := msg.Questions[0]
	var (
		queryTypeStr = question.Type.String()
		queryNameStr = question.Name.String()
		queryType    = question.Type
		queryName, _ = dnsmessage.NewName(queryNameStr)
	)
	fmt.Printf("[%s] queryName: [%s]\n", queryTypeStr, queryNameStr)
	// find record
	var resource dnsmessage.Resource
	switch queryType {
	case dnsmessage.TypeA, dnsmessage.TypeAAAA:
		if rst, ok := addressBookOfA[queryNameStr]; ok {
			resource = NewAResource(queryName, rst)
		} else {
			fmt.Printf("not fount A record queryName: [%s] \n", queryNameStr)
			Response(addr, conn, msg)
			return
		}
	case dnsmessage.TypePTR:
		if rst, ok := addressBookOfPTR[queryName.String()]; ok {
			resource = NewPTRResource(queryName, rst)
		} else {
			fmt.Printf("not fount PTR record queryName: [%s] \n", queryNameStr)
			Response(addr, conn, msg)
			return
		}
	default:
		fmt.Printf("not support dns queryType: [%s] \n", queryTypeStr)
		return
	}
	// send response
	msg.Response = true
	msg.Answers = append(msg.Answers, resource)
	Response(addr, conn, msg)
}

// Response return
func Response(addr *net.UDPAddr, conn *net.UDPConn, msg dnsmessage.Message) {
	packed, err := msg.Pack()
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := conn.WriteToUDP(packed, addr); err != nil {
		fmt.Println(err)
	}
}

// NewAResource A record
func NewAResource(query dnsmessage.Name, a [4]byte) dnsmessage.Resource {
	return dnsmessage.Resource{
		Header: dnsmessage.ResourceHeader{
			Name:  query,
			Class: dnsmessage.ClassINET,
			TTL:   600,
		},
		Body: &dnsmessage.AResource{
			A: a,
		},
	}
}

// NewPTRResource PTR record
func NewPTRResource(query dnsmessage.Name, ptr string) dnsmessage.Resource {
	name, _ := dnsmessage.NewName(ptr)
	return dnsmessage.Resource{
		Header: dnsmessage.ResourceHeader{
			Name:  query,
			Class: dnsmessage.ClassINET,
		},
		Body: &dnsmessage.PTRResource{
			PTR: name,
		},
	}
}
