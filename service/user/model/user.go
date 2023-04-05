package model


User struct {
	Id         int64     `db:"id"`
	Number     string    `db:"number"`   // 学号
	Name       string    `db:"name"`     // 用户名称
	Password   string    `db:"password"` // 用户密码
	Gender     string    `db:"gender"`   // 男｜女｜未公开
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}