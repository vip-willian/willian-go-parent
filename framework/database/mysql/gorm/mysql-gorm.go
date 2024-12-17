package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type User struct {
	UserId     int64      `gorm:"column:id;primaryKey"` // 用户ID
	UserName   string     `gorm:"column:user_name"`     // 用户名
	Password   string     `gorm:"column:password"`      // 密码
	IdCard     string     `gorm:"column:id_card"`       // 身份证号
	Phone      string     `gorm:"column:phone"`         // 手机号
	Email      string     `gorm:"column:email"`         // 邮箱
	Sex        uint       `gorm:"column:sex"`           // 性别
	CreateTime *time.Time `gorm:"column:create_time"`   // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time"`   // 更新时间
}

var db *gorm.DB

// 初始化调用一次，获取数据库连接
func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_user?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁止使用复数
			TablePrefix:   "",   // 表名前缀
		}})
	if err != nil {
		log.Println("mysql connect failed", err)
		return
	}
	db = database
}

// TestInsert 插入数据
func TestInsert() {

	now := time.Now()
	user := &User{
		UserName:   "令狐冲",
		Password:   "54321",
		IdCard:     "339900198409214455",
		Phone:      "13849565535",
		Email:      "linhu@gmail.com",
		Sex:        1,
		CreateTime: &now,
		UpdateTime: &now,
	}
	result := db.Create(user)
	if result.Error != nil {
		log.Println("insert data failed", result.Error)
		return
	}
	log.Printf("insert data success rowAffected:%d,userId:%d", result.RowsAffected, user.UserId)
}

// TestBatchInsert 批量插入数据
func TestBatchInsert() {

	now := time.Now()
	users := []*User{
		{
			UserName:   "张三丰",
			Password:   "abcdef",
			IdCard:     "331010183402105694",
			Phone:      "15677884422",
			Email:      "sanfeng@sina.com",
			Sex:        1,
			CreateTime: &now,
			UpdateTime: &now,
		},
		{
			UserName:   "妲己",
			Password:   "0001",
			IdCard:     "339901199201214455",
			Phone:      "18993215532",
			Email:      "daji@qq.com",
			Sex:        0,
			CreateTime: &now,
			UpdateTime: &now,
		},
		{
			UserName:   "哪吒",
			Password:   "www.nz",
			IdCard:     "310021202401110987",
			Phone:      "13645086629",
			Email:      "nz@gmail.com",
			Sex:        1,
			CreateTime: &now,
			UpdateTime: &now,
		},
	}
	// 指定插入的批次大小
	result := db.CreateInBatches(users, 100)
	if result.Error != nil {
		log.Println("insert data failed", result.Error)
		return
	}
	log.Printf("insert data success rowAffected:%d", result.RowsAffected)
}

// TestQuery 测试查询
func TestQuery() {
	// https://gorm.io/zh_CN/docs/query.html
	var user User
	var users []User

	// 按ID查询
	db.First(&user, 1)
	log.Printf("按主键ID查询 : %#v\n", user)
	// 按ID IN查询
	db.Find(&users, []int{2, 3, 4})
	log.Printf("按主键IN查询: %#v\n", users)

	// 按条件查询
	db.Where("name = ?", "哪吒").First(&user)
	log.Printf("按名字查询结果 result: %#v\n", user)

	db.Where("phone in ?", []string{"15677884422", "15677884423"}).Find(&users)
	log.Printf("按手机号IN查询结果 result: %#v\n", user)
}

// TestUpdate 测试更新
func TestUpdate() {
	var user = User{UserId: 1}
	// 更新单列
	db.Model(&user).Update("password", "664564")
	// 更新多列
	db.Model(&user).Updates(User{Phone: "15698320912", UserName: "吴彦祖"})
	// 按手机号更新
	db.Model(&User{}).Where("phone = ?", "18993215532").Update("email", "daji@souhu.com")
}

// TestDelete 测试删除
func TestDelete() {
	// 删除按主键ID
	db.Delete(&User{}, 1)
	// 删除按主键集合
	db.Delete(&User{}, []int{1, 2, 3})
	// 删除按其他自定义条件
	db.Where("user_name LIKE ?", "%张%").Delete(&User{})
}

// TestAutoTransaction 测试事务
func TestAutoTransaction() {
	now := time.Now()
	user1 := &User{
		UserName:   "zhu",
		Password:   "657465746",
		IdCard:     "314084199006114421",
		Phone:      "15857208890",
		Email:      "kong@gmail.com",
		Sex:        1,
		CreateTime: &now,
		UpdateTime: &now,
	}
	user2 := &User{
		// 用户2手机号重复
		UserName:   "唐僧",
		Password:   "m",
		IdCard:     "334455192008064425",
		Phone:      "13666768848",
		Email:      "ts@google.com",
		Sex:        1,
		CreateTime: &now,
		UpdateTime: &now,
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		// 使用tx进行插入操作
		if err := tx.Create(user1).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		if err := tx.Create(user2).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		log.Println("transaction execute failed", err)
		return
	}
	log.Println("transaction execute success")
}

// TestManualTransaction 测试手动事务
func TestManualTransaction() {
	now := time.Now()
	user1 := &User{
		UserName:   "zhu",
		Password:   "657465746",
		IdCard:     "314084199006114421",
		Phone:      "15857208890",
		Email:      "kong@gmail.com",
		Sex:        1,
		CreateTime: &now,
		UpdateTime: &now,
	}
	user2 := &User{
		// 用户2手机号重复
		UserName:   "唐僧",
		Password:   "m",
		IdCard:     "334455192008064425",
		Phone:      "13666768848",
		Email:      "ts@google.com",
		Sex:        1,
		CreateTime: &now,
		UpdateTime: &now,
	}
	tx := db.Begin()
	// 使用tx进行插入操作
	if err := tx.Create(user1).Error; err != nil {
		// 返回任何错误都会回滚事务
		log.Println("transaction execute failed", err)
		tx.Rollback()
		return
	}
	if err := tx.Create(user2).Error; err != nil {
		// 返回任何错误都会回滚事务
		log.Println("transaction execute failed", err)
		tx.Rollback()
		return
	}
	log.Println("transaction execute success")
	tx.Commit()
}
