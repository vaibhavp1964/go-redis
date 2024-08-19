package constants

type RESPConst string

const (
	CRLF          RESPConst = "\r\n"
	SIMPLE_STRING RESPConst = "+"
	BULK_STRING   RESPConst = "$"
	ARRAY         RESPConst = "*"
	CRLF_LEADING  RESPConst = "\r"
)
