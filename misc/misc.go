package misc

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

// RandomString random string with prefix (default) and custom length {prefix}-XXXX9999
func RandomString(prefix *string, lenString int, lenNum int) string {
	const charStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const charNum = "0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	rndCharStr := make([]byte, lenString)
	rndCharNum := make([]byte, lenNum)
	for r := range rndCharStr {
		rndCharStr[r] = charStr[seededRand.Intn(len(charStr))]
	}
	for r := range rndCharNum {
		rndCharNum[r] = charNum[seededRand.Intn(len(charNum))]
	}
	if prefix == nil {
		return string(rndCharStr) + string(rndCharNum)
	}
	return *prefix + string(rndCharStr) + string(rndCharNum)
}

// FilenameToUUID rename exisiting filename to UUID
func FilenameToUUID(filename string) string {
	names := strings.Split(filename, ".")
	uuid := uuid.New().String()
	return fmt.Sprintf("%s.%s", uuid, names[len(names)-1])
}
