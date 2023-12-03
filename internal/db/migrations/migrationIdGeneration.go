package migrations

import (
	"fmt"
	"time"
)

func GenerateMigrationID() string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s%s", timestamp, "01")
}
