package model

type Project struct {
	ProjectId    string `gorm:"column:ProjectId;NOT NULL" json:"ProjectId"` // 主键
	ProjectName  string `gorm:"column:ProjectName" json:"ProjectName"`      // 项目名称
	ProjectCode  string `gorm:"column:ProjectCode" json:"ProjectCode"`      // 项目编码
	TaxSubjectId string `gorm:"column:TaxSubjectId" json:"TaxSubjectId"`    // 项目所属税务主体
	Remark       string `gorm:"column:Remark" json:"Remark"`                // 备注
	LandPrice    string `gorm:"column:LandPrice" json:"LandPrice"`
	DeAmount     string `gorm:"column:DeAmount" json:"DeAmount"`
	AreaCovered  string `gorm:"column:AreaCovered" json:"AreaCovered"`
	IsDevProject string `gorm:"column:IsDevProject" json:"IsDevProject"`
	DefaultSKR   string `gorm:"column:DefaultSKR" json:"DefaultSKR"`
	DefaultFHR   string `gorm:"column:DefaultFHR" json:"DefaultFHR"`
}

func (m *Project) TableName() string {
	return "ygz_s_project"
}

type ProjectRelation struct {
	ProjectRelationId string `gorm:"column:ProjectRelationId;NOT NULL" json:"ProjectRelationId"` // 主键
	ProjectId         string `gorm:"column:ProjectId" json:"ProjectId"`                          // 项目Id
	ProjectInfoId     string `gorm:"column:ProjectInfoId" json:"ProjectInfoId"`                  // 异构系统项目Id
}

func (m *ProjectRelation) TableName() string {
	return "ygz_s_projectrelation"
}

type InterfaceProjectInfo struct {
	ProjectInfoId string `gorm:"column:ProjectInfoId;NOT NULL" json:"ProjectInfoId"` // 主键
	ProjectName   string `gorm:"column:ProjectName" json:"ProjectName"`              // 项目名称
	ProjectCode   string `gorm:"column:ProjectCode" json:"ProjectCode"`              // 项目编码
	CompanyId     string `gorm:"column:CompanyId" json:"CompanyId"`                  // 所属公司
	CompanyName   string `gorm:"column:CompanyName" json:"CompanyName"`              // 所属公司名称
	SourceSystem  string `gorm:"column:SourceSystem" json:"SourceSystem"`            // 数据来源系统
	SourceId      string `gorm:"column:SourceId" json:"SourceId"`                    // 数据来源Id
	ParentCode    string `gorm:"column:ParentCode" json:"ParentCode"`                // 父级编码
	ParentGUID    string `gorm:"column:ParentGUID" json:"ParentGUID"`                // 父级GUID
}

func (m *InterfaceProjectInfo) TableName() string {
	return "ygz_interface_projectinfo"
}
