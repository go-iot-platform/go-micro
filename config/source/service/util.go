package service

import (
	"time"

	"github.com/panovateam/go-micro/config/source"
	proto "github.com/panovateam/go-micro/config/source/service/proto"
)

func toChangeSet(c *proto.ChangeSet) *source.ChangeSet {
	return &source.ChangeSet{
		Data:      []byte(c.Data),
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: time.Unix(c.Timestamp, 0),
		Source:    c.Source,
	}
}
