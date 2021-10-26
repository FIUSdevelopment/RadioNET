package controllers

import (
	"database/sql"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
)

type CheckusrController struct {
	beego.Controller
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (t *CheckusrController) Post() {
	db, err := sql.Open("sqlite3", "database/users.db")
	checkErr(err)
	usr := t.GetString("usr");
	psw := t.GetString("psw");

	fmt.Println(usr)
	fmt.Println(psw)

	rows, err := db.Query("SELECT password FROM users WHERE username = usr")
	checkErr(err)

	if(rows == nil) {
		t.Data["json"] = true;
		t.ServeJSON();  
	} else {
		t.Data["json"] = false;
		t.ServeJSON();  
	}
}
