package model

import (
	"time"

	"github.com/lib/pq"
)

type ProjectPool struct {
	ID            	   	int							`json:"id" gorm:"primaryKey;autoIncrement"`
	Title								string					`json:"title"`
	SubTitle						string					`json:"sub_title"`
	Description					string					`json:"description"`
	Source							string					`json:"source"`
	StartDate						time.Time				`json:"start_date"`
	EndDate							time.Time				`json:"end_date"`
	ProjectList					pq.Int64Array		`gorm:"type:integer[]" json:"project_list"`
	UpdatedAt        		time.Time				`json:"updated_at"`
	CreatedAt         	time.Time				`json:"created_at"`
	DeletedAt         	*time.Time			`json:"deleted_at"`
}