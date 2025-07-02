package utils

import (
	"strconv"
)

func ParseUserID(userID string) (uint, error) {
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
