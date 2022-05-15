package services

import "time"

func CalculateExpiryDate(validity int32) time.Time {
	t := time.Now()
	return t.Add(time.Hour * 24 * time.Duration(validity))
}
