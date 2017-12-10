package user

/**
 * 用户表
 */
type User struct {
	Id         int    `orm:"type(int);pk;auto" form:"-"`
	Username   string `orm:"type(varchar);unique;size(32)" form:"username" valid:"required;MinSize:6;MaxSize:32"`
	Password   string `orm:"type(char);size(32)" form:"password" valid:"MinSize(6);MaxSize(16)"`
	Mobile     string `orm:"type(varchar);size(15)" form:"mobile" valid:"Phone;MinSize(11);Max(15)"`
	Nickname   string `orm:"type(varchar);size(50)" form:"nickname" valid:"Max(50)"`
	UserType   int    `orm:"type(tinyint)" form:"-"`
	State      int    `orm:"type(tinyint);size(1);default(1)" form:"-"`
	IsDelete   int    `orm:"type(tinyint);size(1);default(0)" form:"-"`
	CreateTime string `orm:"type(datetime)" form:"-"`
	UpdateTime string `orm:"type(timestamp);auto_now_add" form:"-"`
}

/**
 *
 */
type UserCompanyInfo struct {
	Id          int    `orm:"type(int);pk;auto" form:"-"`
	CompanyName string `orm:"type(varchar);size(100)" form:"companyName" valid:"required;MaxSize:100"`
	IsDelete    int    `orm:"type(tinyint);size(1);default(0)" form:"-"`
	CreateTime  string `orm:"type(datetime)" form:"-"`
	UpdateTime  string `orm:"type(timestamp);auto_now_add" form:"-"`
}

type UserCustomer struct {
	Id          int    `orm:"type(int);pk;auto" form:"-"`
	UserId      int    `orm:"type(int);" form:"-"`
	CompanyId   int    `orm:"type(int);" form:"-"`
	OpenId      string `orm:"type(char);size(32)" form:"-"`
	CustName    string `orm:"type(varchar);size(50)" form:"nickname" valid:"Max(50)"`
	CustLevel   string `orm:"type(tinyint)" form:"-"`
	CustNum     string `orm:"type(char);size(32)"`
	WechatNum   string `orm:"type(varchar);size(50)" form:"wechatNum" valid:"Max(50)"`
	Mobile      string `orm:"type(varchar);size(15)" form:"mobile" valid:"Phone;MinSize(11);Max(15)"`
	AreaCode    int    `orm:"type(int)" form:"areaCode" valid:"numeric"`
	Address     string `orm:"type(varchar);size(100)" form:"address" valid:"Max(100)"`
	DeveloperId int    `orm:"type(int)" form:"developerId" valid:"numeric"`
	DepositBank string `orm:"type(varchar);size(50)" form:"depositBank" valid:"Max(50)"`
	BankAccount string `orm:"type(varchar);size(20)" form:"bankAccount" valid:"Max(20)"`
	State       int    `orm:"type(tinyint);size(1);default(1)" form:"-"`
	IsDelete    int    `orm:"type(tinyint);size(1);default(0)" form:"-"`
	CreateTime  string `orm:"type(datetime)" form:"-"`
	UpdateTime  string `orm:"type(timestamp);auto_now_add" form:"-"`
}

type UserStaffInfo struct {
	Id         int    `orm:"type(int);pk;auto" form:"-"`
	CompanyId  int    `orm:"type(int);" form:"-"`
	StaffName  string `orm:"type(varchar);size(50)" form:"staffName" valid:"Max(50)"`
	StaffNum   string `orm:"type(varchar);size(32)" form:"staffNum" valid:"Max(50)"`
	Mobile     string `orm:"type(varchar);size(15)" form:"mobile" valid:"Phone;MinSize(11);Max(15)"`
	IsDelete   int    `orm:"type(tinyint);size(1);default(0)" form:"-"`
	CreateTime string `orm:"type(datetime)" form:"-"`
	UpdateTime string `orm:"type(timestamp);auto_now_add" form:"-"`
}

type UserToken struct {
	Id         int    `orm:"type(int);pk;auto" form:"-"`
	UserId     int    `orm:"type(int);" form:"-"`
	TokenType  int    `orm:"type(tinyint)" form:"-"`
	Token      string `orm:"type(char);size(32)"`
	IsDelete   int    `orm:"type(tinyint);size(1);default(0)" form:"-"`
	CreateTime string `orm:"type(datetime)" form:"-"`
	UpdateTime string `orm:"type(timestamp);auto_now_add" form:"-"`
}
