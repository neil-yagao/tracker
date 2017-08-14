package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Migration_20170810_151002 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Migration_20170810_151002{}
	m.Created = "20170810_151002"
	migration.Register("Migration_20170810_151002", m)
}

// Run the migrations
func (m *Migration_20170810_151002) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`
		CREATE TABLE IF NOT EXISTS 'movement_video' (
		    'id' bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
		    'mapped_movement_id' bigint NOT NULL,
		    'location' varchar(1024) NOT NULL DEFAULT '' ,
		    'uploaded_by_id' bigint NOT NULL,
		    'uploaded_at date NOT NULL
		) ENGINE=InnoDB`)

}

// Reverse the migrations
func (m *Migration_20170810_151002) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("Drop Table 'movement_video'")
}
