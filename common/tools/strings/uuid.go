package strings

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func UUID() string {
	uid := uuid.NewV4()
	str := strings.ReplaceAll(uid.String(), "-", "")
	return str
}
