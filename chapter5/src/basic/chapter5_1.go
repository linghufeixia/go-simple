/**
 * Created with IntelliJ IDEA.
 * User: 令狐飞侠
 * Date: 2021-6-10
 * Description: 数据库的操作
 */
package basic

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//测试查询数据
func TestSelect()  {
	//打开数据库
	db,err:= OpenDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}
	//查询数据
	SelectData(db)

	db.Close()
}

//测试修改数据
func TestUpdate()  {
	//打开数据库
	db,err:= OpenDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	//修改数据
	result,err:= UpdateData(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	aID, _ := result.LastInsertId()
	fmt.Println(aID)

	//查询数据
	SelectData(db)

	db.Close()
}

//打开数据库
func OpenDatabase() (*sql.DB, error) {
	db, err:= sql.Open("mysql", "root:123456@tcp(localhost:3306)/db_test")
	return db, err
}

//查询数据
func SelectData(db *sql.DB)  {
	rows, err := db.Query("select id, user_name, user_pwd, user_phone from t_user")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var userName string
		var userPwd string
		var userPhone string
		rows.Scan(&id, &userName, &userPwd, &userPhone)
		fmt.Println(id)
		fmt.Println(userName)
		fmt.Println(userPwd)
		fmt.Println(userPhone)
	}
}

//修改数据
func UpdateData(db *sql.DB)  (sql.Result, error) {
	//修改
	results, err:= db.Exec("UPDATE t_user SET user_pwd=? WHERE user_name=?","654321","令狐冲" )
	return results,err
}
