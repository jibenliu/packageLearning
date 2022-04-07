package config

const (
	NET_TRANS_PROTOCOL = "tcp"
	TRANS_TYPE         = "gob" //json
	HEADER_LEN         = 4     //header length
)

type CodecMode int

const (
	CODEC_GOB CodecMode = iota
	CODEC_JSON
)

const (
	Protocol_MsgVersion = 1
)
