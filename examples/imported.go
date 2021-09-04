package examples

import (
	"encoding/json"

	"github.com/google/uuid"
)

//go:generate go run ../main.go -type Imported
type Imported struct {
	id    uuid.UUID
	extra json.RawMessage
}
