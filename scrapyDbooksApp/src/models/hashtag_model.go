package models

import (
    "database/sql"
    "entities"
    "time"
)

type HashtagModel struct {
    Db *sql.DB
}

func (hashtagModel HashtagModel)insertsql string(){
    _.err := hashtagModel.Db.Exec(sql)
    if err != nil {
    }
}
