package model

import (
	"time"

	"github.com/lib/pq"
)

type ProjectPool struct {
	Title								string					`json:"title"`
	SubTitle						string					`json:"sub_title"`
	Description					string					`json:"description"`
	Source							string					`json:"source"`
	StartDate						time.Time				`json:"start_date"`
	EndDate							time.Time				`json:"end_date"`
	ProjectList					pq.Int64Array		`gorm:"type:integer[]" json:"project_list"`
}