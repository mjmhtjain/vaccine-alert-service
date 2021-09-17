package util

import (
	"embed"
	"time"
)

const (
	SLATimeout                           = 2 * time.Second
	URL_AppointmentSessionForWeek string = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict"
)

var EmbededFiles embed.FS
