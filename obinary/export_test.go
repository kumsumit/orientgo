package obinary

import (
	"orient/obinary/rw"
)

func ReadErrorResponse(r *rw.Reader) (serverException error) {
	return readErrorResponse(r, CurrentProtoVersion)
}
