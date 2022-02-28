package models

type QueryApprovedCCInfo struct {
	ChaincodeId string `json:"chaincode_id" validate:"required,gt=0"`
	Org         string `json:"org" validate:"required,gt=0"`
	UserName    string `json:"user_name" validate:"required,gt=0"`
	ChannelId   string `json:"channel_id" validate:"required,gt=0"`
	Peer        string `json:"peer" validate:"required,gt=0"`
	Sequence    string `json:"sequence" validate:"required,gt=0"`
}

type Education struct {
	ObjectType string `json:"docType"`

	UserID   string `json:"user_id"`
	Msg      string `json:"msg"`
	ClassID  string `json:"class_id"`
	SchoolID string `json:"school_id"`

	/*
		Name       string `json:"Name"`     // 姓名
		Gender     string `json:"Gender"`   // 性别
		Nation     string `json:"Nation"`   // 民族
		EntityID   string `json:"EntityID"` // 身份证号
		Place      string `json:"Place"`    // 籍贯
		BirthDay   string `json:"BirthDay"` // 出生日期

		EnrollDate     string `json:"EnrollDate"`     // 入学日期
		GraduationDate string `json:"GraduationDate"` // 毕（结）业日期
		SchoolName     string `json:"SchoolName"`     // 学校名称
		Major          string `json:"Major"`          // 专业
		QuaType        string `json:"QuaType"`        // 学历类别
		Length         string `json:"Length"`         // 学制
		Mode           string `json:"Mode"`           // 学习形式
		Level          string `json:"Level"`          // 层次
		Graduation     string `json:"Graduation"`     // 毕（结）业
		CertNo         string `json:"CertNo"`         // 证书编号

		Course  string `json:"Course"`
		Teacher string `json:"Teacher"`

		Photo string `json:"Photo"` // 照片
	*/

	Historys []HistoryItem // 当前edu的历史记录
}

type HistoryItem struct {
	TxId      string
	Education Education
}

/*
type EducationInfo struct {
	ObjectType string `json:"docType" validate:"required,gt=0"`
	Name       string `json:"Name" validate:"required,gt=0"`     // 姓名
	Gender     string `json:"Gender" validate:"required,gt=0"`   // 性别
	Nation     string `json:"Nation" validate:"required,gt=0"`   // 民族
	EntityID   string `json:"EntityID" validate:"required,gt=0"` // 身份证号
	Place      string `json:"Place" validate:"required,gt=0"`    // 籍贯
	BirthDay   string `json:"BirthDay" validate:"required,gt=0"` // 出生日期

	EnrollDate     string `json:"EnrollDate" validate:"required,gt=0"`     // 入学日期
	GraduationDate string `json:"GraduationDate" validate:"required,gt=0"` // 毕（结）业日期
	SchoolName     string `json:"SchoolName" validate:"required,gt=0"`     // 学校名称
	Major          string `json:"Major" validate:"required,gt=0"`          // 专业
	QuaType        string `json:"QuaType" validate:"required,gt=0"`        // 学历类别
	Length         string `json:"Length" validate:"required,gt=0"`         // 学制
	Mode           string `json:"Mode" validate:"required,gt=0"`           // 学习形式
	Level          string `json:"Level" validate:"required,gt=0"`          // 层次
	Graduation     string `json:"Graduation" validate:"required,gt=0"`     // 毕（结）业
	CertNo         string `json:"CertNo" validate:"required,gt=0"`         // 证书编号

	Photo string `json:"Photo"` // 照片
}

type HistoryItem struct {
	TxId      string
	Education EducationInfo
}

*/
