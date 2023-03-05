package protocol

type SP_CL2LS_REQ_LOGIN struct {
	ID            string `size:"33"`
	Password      string `size:"33"`
	ClientVerA    int32
	ClientVerB    int32
	ClientVerC    int32
	LoginType     int32
	Cookie_TEGid  [64]uint8
	Cookie_authid [255]uint8
}
