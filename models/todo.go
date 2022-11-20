package models

import "time"

type Todo struct {
	ID        int64     `xorm:"bigint not null unique 'id' autoincr pk" json:"id"`
	Name      string    `xorm:"varchar(255) not null 'name'" json:"name"`
	Email     string    `xorm:"varchar(255) not null 'email'" json:"email"`
	CreatedAt time.Time `xorm:"'created_at'" json:"created_at"`
	UpdatedAt time.Time `xorm:"'updated_at'" json:"-"`
}
