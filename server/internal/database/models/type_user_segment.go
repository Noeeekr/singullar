package models

import (
	"github.com/Noeeekr/borm"
)

type UserSegments string

const UserSegmentsName = "user_segments"

const (
	UNKNOWN_SEGMENT UserSegments = ""
	EF1             UserSegments = "ensino fundamental 1"
	EF2             UserSegments = "ensino fundamental 2"
	EM              UserSegments = "ensino medio"
)

var TypeUserSegment *borm.Enum = EnvironmentDatabase.RegisterEnum(UserSegmentsName, string(EF1), string(EF2), string(EM))
