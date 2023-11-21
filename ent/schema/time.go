package schema

import (
	"time"

	"github.com/ugent-library/old-people-service/models"
)

var genBeginningOfTime = func() time.Time {
	return models.BeginningOfTime
}
