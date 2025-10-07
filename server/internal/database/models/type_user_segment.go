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

// var UserSegmentType *TypeInfo = &TypeInfo{
// 	Name: userSegmentTypeName,
// 	Queries: &TypeQueries{
// 		Create: transactions.NewRequest(fmt.Sprintf(`
// 			DO $$
// 			BEGIN
// 				IF NOT EXISTS (SELECT typname FROM pg_type WHERE typname = '%s') THEN
// 					CREATE TYPE %s AS ENUM (%s, %s, %s)
// 				END IF;
// 			END $$;
// 		`, userSegmentTypeName, userSegmentTypeName, EF1, EF2, EM)),
// 		Drop: transactions.NewRequest(fmt.Sprintf(`
// 			DROP TYPE IF EXIST %s;
// 		`, userSegmentTypeName)),
// 	},
// }
