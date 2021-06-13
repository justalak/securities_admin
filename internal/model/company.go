package model

import "time"

type Company struct {
	ID               int32 `gorm:"primaryKey"`
	Name             string
	ShortName        string
	Symbol           string `gorm:"index"`
	DateOfIssue      *time.Time
	DateOfListing    *time.Time
	Employees        int32
	ListingVolume    uint64
	IcbCode          int32
	ClassificationID int32
	Overview         string
}
