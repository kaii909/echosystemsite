package version

import (
	"fmt"
	"time"
)

var AppVersion = fmt.Sprintf("%d", time.Now().Unix())
