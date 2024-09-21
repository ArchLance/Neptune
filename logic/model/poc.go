package model

// 漏洞类型
const (
	// ArbitraryFileRead 任意文件读取
	ArbitraryFileRead = 1
	// ArbitraryFileUpload 文件上传
	ArbitraryFileUpload = 2
	// RemoteCommandExecute 远程命令执行
	RemoteCommandExecute = 3
	// InformationDisclosure 信息泄露
	InformationDisclosure = 4
	// JavaUnserialize JAVA反序列化
	JavaUnserialize = 5
	// PhpUnserialize PHP反序列化
	PhpUnserialize = 6
	// HorizontalPrivilegeEscalation 水平越权
	HorizontalPrivilegeEscalation = 7
	// VerticalPrivilegeEscalation 垂直越权
	VerticalPrivilegeEscalation = 8
	// SqlInjection SQL注入
	SqlInjection = 9
	// CrossSiteScripting 跨站脚本攻击(XSS)
	CrossSiteScripting = 10
	// ServerSideRequestForgery 服务端请求伪造(SSRF)
	ServerSideRequestForgery = 11
	// ServerSideTemplateInjection 服务端模版注入(SSTI)
	ServerSideTemplateInjection = 12
	// LogicVulnerability 逻辑缺陷
	LogicVulnerability = 13
	// ExternalEntityInjection 外部实体注入(XXE)
	ExternalEntityInjection = 14
	// PermissionBypass 权限绕过
	PermissionBypass = 15
	// FileInclusion 文件包含
	FileInclusion  = 16
	PLACEHOLDER_17 = 17
	PLACEHOLDER_18 = 18
	PLACEHOLDER_19 = 19
	PLACEHOLDER_20 = 20
)

type Poc struct {
	// 主键
	Id int `gorm:"type:int;primary_key;AUTO_INCREMENT"`
	// 漏洞编号
	VulnerabilityName string `gorm:"type:varchar(64)"`
	// poc名称
	PocName string `gorm:"type:varchar(64);not null;unique"`
	// 应用名称
	AppName string `gorm:"type:varchar(64)"`
	// 漏洞类型
	VulnerabilityType int `gorm:"type:int"`
	// 添加时间
	AddTime string `gorm:"type:varchar(64)"`
	// poc内容
	PocContent string `gorm:"type:text"`
}