package model

import (
	"time"
)

// Fact ...
type Fact struct {
	PeriodStart         time.Time
	PeriodEnd           time.Time
	PeriodKey           string
	IndicatorToMoID     int
	IndicatorToMoFactID int
	Value               int
	FactTime            time.Time
	IsPlan              int
	AuthUserID          int
	Comment             string
}
