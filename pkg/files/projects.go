package files

import (
	"os"
	"strconv"
	"sync"

	"github.com/google/uuid"
)

var (
	currentID int
	mutex     sync.Mutex // Mutex to protect currentID
)

func GenerateUUID() string {
	if os.Getenv("DETERMINISTIC_IDS") == "true" {
		mutex.Lock()
		defer mutex.Unlock()
		currentID++
		return strconv.Itoa(currentID)
	}
	return uuid.New().String()
}
