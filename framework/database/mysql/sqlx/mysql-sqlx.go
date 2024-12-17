package sqlx

import (
	"fmt"
	"time"

	// 需要导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	UserId     int64     `db:"id"`          // 用户ID
	UserName   string    `db:"user_name"`   // 用户名
	Password   string    `db:"password"`    // 密码
	IdCard     string    `db:"id_card"`     // 身份证号
	Phone      string    `db:"phone"`       // 手机号
	Email      string    `db:"email"`       // 邮箱
	Sex        uint      `db:"sex"`         // 性别
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
}

var db *sqlx.DB

// 初始化调用一次，获取数据库连接
func init() {
	// 获取数据库连接
	conn, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db_user?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	db = conn
}

// TestInsert 插入数据
func TestInsert() {

	r, err := db.Exec("insert into user(user_name,password,id_card,phone,email,sex) values(?,?,?,?,?,?)", "李世民", "023432", "339900198206109000", "13421219090", "shiming@qq.com", 1)
	if err != nil {
		fmt.Println("insert exec failed, ", err)
		return
	}
	// 获取插入的ID
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert success: ", id)
	defer db.Close()
}

// TestUpdate 数据库更新
func TestUpdate() {
	r, err := db.Exec("update user set email = ? where id = ?", "tang@qq.com", 6)
	if err != nil {
		fmt.Println("update exec failed, ", err)
		return
	}
	row, err := r.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update success:", row)
	defer db.Close()
}

// TestDelete 数据库删除
func TestDelete() {
	r, err := db.Exec("delete from user where id = ?", 1)
	if err != nil {
		fmt.Println("delete exec failed, ", err)
		return
	}
	row, err := r.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("delete success:", row)
	defer db.Close()
}

// TestQuery 测试查询
func TestQuery() {
	var users []User
	err := db.Select(&users, "select id,user_name,password,id_card,phone,email,sex,create_time,update_time from user where phone = ?", "18993215532")
	if err != nil {
		fmt.Println("select exec failed, ", err)
		return
	}
	fmt.Printf("select result : %#v\n", users)
	defer db.Close()
}

// TestTransaction 数据库事务
func TestTransaction() {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}
	r, err := tx.Exec("insert into user(user_name,password,id_card,phone,email,sex) values(?,?,?,?,?,?)", "关羽", "435356", "330801196809095645", "18042485432", "guan@gmail.com", 1)
	if err != nil {
		fmt.Println("insert first exec failed, ", err)
		// 执行失败，回滚
		_ = tx.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("insert first exec failed, ", err)
		_ = tx.Rollback()
		return
	}
	fmt.Println("insert first success:", id)

	// 第二条插入失败
	r, err = tx.Exec("insert into user(user_name,password,id_card,phone,email,sex) values(?,?,?,?,?,?)", "李世民", "023432", "339900198206109000", "13666768848", "shiming@qq.com", 1)
	if err != nil {
		fmt.Println("insert second exec failed, ", err)
		// 执行失败，回滚
		_ = tx.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("insert second exec failed, ", err)
		_ = tx.Rollback()
		return
	}
	fmt.Println("insert second success:", id)
	// 提交事务
	_ = tx.Commit()
	defer db.Close()
}
