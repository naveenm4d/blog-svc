package adapters

import (
	"github.com/naveenm4d/blog-svc/proto"
)

type Handler interface {
	proto.BlogSvcServer
}
