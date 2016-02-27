package conf
import (
	"database/sql"
)


var Conn = "root:pandora@tcp(127.0.0.1:3306)/psilibrary?parseTime=true"
var Db = "mysql"
var Open = sql.Open

