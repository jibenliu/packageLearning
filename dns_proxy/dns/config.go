package dns

// address books
var (
	addressBookOfA = map[string][4]byte{
		"www.baidu1.com.": [4]byte{127, 0, 0, 8},
		"xiazemin.com.":   [4]byte{127, 0, 0, 1},
	}
	addressBookOfPTR = map[string]string{
		"127.0.0.8.in-addr.arpa.": "www.baidu.com.",
		"8.0.0.127.in-addr.arpa.": "xiazemin.com.",
		//[TypePTR] queryName: [8.0.0.127.in-addr.arpa.]
		//packing Answer: content: name is not in canonical format (it must end with a .)
	}
)
