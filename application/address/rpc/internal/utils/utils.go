package utils

import (
	"fmt"
)

func CacheKey(id int64) string {
	return fmt.Sprintf("%s#%d", CacheAddressKey, id)
}
