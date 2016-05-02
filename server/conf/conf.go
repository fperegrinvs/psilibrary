package conf
import (
        "database/sql"
)


var Conn = "lstern:pandora123@tcp(psilibrary.cifcc2gdswrd.sa-east-1.rds.amazonaws.com:3306)/psilibrary?parseTime=true"
var Db = "mysql"
var Open = sql.Open
var PageSize = 20

