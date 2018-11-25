package models

import (
	"time"
)

/*type Session struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"id"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,json=token,proto3" json:"token`
	TTL                  int64    `protobuf:"varint,3,opt,name=TTL,json=tTL,proto3" json:"ttl"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`

	User *User
}*/

func (s *Session) IsAlive() bool {
	return s.TTL > time.Now().Unix()
}