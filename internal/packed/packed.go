package packed

import (
	"apis/internal/logic/db"
	"apis/internal/service"
)

func init() {
	service.RegisterDB(db.New())
}
