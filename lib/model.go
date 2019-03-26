package lib

import (
	"github.com/jinzhu/gorm"
	"math"
	"time"
)

type CAtModel struct {
	CreatedAt time.Time
}

type CuAtModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CudAtModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type OptionModel struct {
	Key   int
	Name  string
	Value string
}

type PaginationParam struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
}

type Pagination struct {
	Total     int
	TotalPage int
	Data      interface{}
	Offset    int
	Limit     int
	Page      int
	PrevPage  int
	NextPage  int
}

func Paginate(p *PaginationParam, dataSource interface{}) *Pagination {
	db := p.DB

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 25
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	} else {
		db = db.Order("id desc")
	}

	done := make(chan bool, 1)
	var pagination Pagination
	var count int
	var offset int

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	go totalCount(db, dataSource, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	db.Limit(p.Limit).Offset(offset).Find(dataSource)
	<-done
	pagination.Total = count
	pagination.Data = dataSource
	pagination.Page = p.Page

	pagination.Offset = offset
	pagination.Limit = p.Limit
	pagination.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		pagination.PrevPage = p.Page - 1
	} else {
		pagination.PrevPage = p.Page
	}

	if p.Page == pagination.TotalPage {
		pagination.NextPage = p.Page
	} else {
		pagination.NextPage = p.Page + 1
	}
	return &pagination
}

func totalCount(db *gorm.DB, countDataSource interface{}, done chan bool, count *int) {
	db.Model(countDataSource).Count(count)
	done <- true
}
