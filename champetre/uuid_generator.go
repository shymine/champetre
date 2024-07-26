package champetre

import "github.com/google/uuid"

type UUIdGenerator struct {
}

func getUUId() string {
	return uuid.NewString()
}

func isUUIdDefault(uuid string) bool {
	return uuid == ""
}