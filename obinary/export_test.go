package obinary

import (
	"github.com/kumsumit/orientgo/obinary/rw"
)

func ReadErrorResponse(r *rw.Reader) (serverException error) {
	return readErrorResponse(r, CurrentProtoVersion)
}
