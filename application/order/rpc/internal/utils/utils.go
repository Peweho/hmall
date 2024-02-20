package utils

import "fmt"

func CacheKey(id int) string {
	return fmt.Sprintf("%s#%d", CacheOrderKey, id)
}
