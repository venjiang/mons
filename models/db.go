package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	. "github.com/venjiang/mons/core"
	"log"
)

func InitDb() (*gorp.DbMap, error) {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	// db, err := sql.Open("postgres", "/tmp/post_db.bin")
	driver := Config.MustValue("db", "driver_name")
	dataSource := Config.MustValue("db", "data_source")
	db, err := sql.Open(driver, dataSource)
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	// dbmap := &gorp.DbMap{Db: db, Dialect: dialect}
	dbmap := &gorp.DbMap{Db: db, Dialect: getDialectFromDriver(driver)}

	// user
	user := dbmap.AddTableWithName(User{}, "user").SetKeys(true, "Id")
	user.ColMap("Id").Rename("id")
	user.ColMap("Username").Rename("username").SetMaxSize(64).SetNotNull(true)
	user.ColMap("Password").Rename("password").SetMaxSize(128).SetNotNull(true)
	user.ColMap("Email").Rename("email").SetMaxSize(256).SetNotNull(true)
	user.ColMap("Intro").SetMaxSize(512)
	user.ColMap("Website").SetMaxSize(256)
	user.ColMap("CreatedTime").Rename("created_time")
	user.ColMap("CreatedAt").Rename("created_at")
	user.ColMap("UpdatedTime").Rename("updated_time")

	// post
	post := dbmap.AddTableWithName(Post{}, "post").SetKeys(true, "Id")
	post.ColMap("Title").SetMaxSize(128)
	post.ColMap("Excerpt").SetMaxSize(512)
	post.ColMap("Password").SetMaxSize(128)
	post.ColMap("UrlName").SetMaxSize(128)
	post.ColMap("Url").SetMaxSize(256)
	post.ColMap("Layout").SetMaxSize(64)
	post.ColMap("Ip").SetMaxSize(24)

	// comment
	comment := dbmap.AddTableWithName(Comment{}, "comment").SetKeys(true, "Id")
	comment.ColMap("Username").SetMaxSize(64)
	comment.ColMap("Website").SetMaxSize(256)
	comment.ColMap("Title").SetMaxSize(128)
	comment.ColMap("Content").SetMaxSize(512)
	comment.ColMap("Ip").SetMaxSize(24)

	// tag
	tag := dbmap.AddTableWithName(Tag{}, "tag").SetKeys(true, "Id")
	tag.ColMap("Name").SetMaxSize(64)

	// post tag
	posttag := dbmap.AddTableWithName(PostTag{}, "post_tag").SetKeys(true, "Id")
	_ = posttag

	// site
	site := dbmap.AddTableWithName(Site{}, "site").SetKeys(true, "Id")
	site.ColMap("Name").SetMaxSize(64)
	site.ColMap("Description").SetMaxSize(512)
	site.ColMap("Url").SetMaxSize(256)
	site.ColMap("Theme").SetMaxSize(64)
	site.ColMap("Layout").SetMaxSize(64)
	site.ColMap("AdminEmail").SetMaxSize(256)
	site.ColMap("DateFormat").SetMaxSize(64)
	site.ColMap("DatetimeFormat").SetMaxSize(64)
	site.ColMap("DatetimeShortFormat").SetMaxSize(64)
	site.ColMap("MailServerUrl").SetMaxSize(256)
	site.ColMap("MailServerLogin").SetMaxSize(128)
	site.ColMap("MailServerPassword").SetMaxSize(128)

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap, nil
}

// func (svc *Service) InitSchema() error {
// 	if err := svc.DbMap.CreateTablesIfNotExists(); err != nil {
// 		return err
// 	}
// 	return nil
// }
func getDialectFromDriver(driver string) gorp.Dialect {
	switch driver {
	case "mymysql":
		return gorp.MySQLDialect{"InnoDB", "UTF8"}
	case "mysql":
		return gorp.MySQLDialect{"InnoDB", "UTF8"}
	case "gomysql":
		return gorp.MySQLDialect{"InnoDB", "UTF8"}
	case "postgres":
		return gorp.PostgresDialect{}
	case "sqlite":
		return gorp.SqliteDialect{}
	case "sqlite3":
		return gorp.SqliteDialect{}
	}
	panic("GORP DIALECT is not set or is invalid.")
}
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func OpenDb() *gorp.DbMap {
	dbmap, err := InitDb()

	if err != nil {
		panic(err)
	}
	// defer dbmap.Db.Close()

	return dbmap
}
