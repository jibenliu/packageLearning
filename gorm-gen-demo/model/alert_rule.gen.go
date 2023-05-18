// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAlertRule = "alert_rule"

// AlertRule mapped from table <alert_rule>
type AlertRule struct {
	ID               int64  `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	GroupID          int64  `gorm:"column:group_id;type:bigint(20);not null;comment:busi group id" json:"group_id"`
	Cate             string `gorm:"column:cate;type:varchar(128);not null" json:"cate"`
	DatasourceIds    string `gorm:"column:datasource_ids;type:varchar(255);not null;comment:datasource ids" json:"datasource_ids"`
	Cluster          string `gorm:"column:cluster;type:varchar(128);not null" json:"cluster"`
	Name             string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Note             string `gorm:"column:note;type:varchar(1024);not null" json:"note"`
	Prod             string `gorm:"column:prod;type:varchar(255);not null" json:"prod"`
	Algorithm        string `gorm:"column:algorithm;type:varchar(255);not null" json:"algorithm"`
	AlgoParams       string `gorm:"column:algo_params;type:varchar(255)" json:"algo_params"`
	Delay            int64  `gorm:"column:delay;type:int(11);not null" json:"delay"`
	Severity         int64  `gorm:"column:severity;type:tinyint(1);not null;comment:1:Emergency 2:Warning 3:Notice" json:"severity"`
	Disabled         int64  `gorm:"column:disabled;type:tinyint(1);not null;comment:0:enabled 1:disabled" json:"disabled"`
	PromForDuration  int64  `gorm:"column:prom_for_duration;type:int(11);not null;comment:prometheus for, unit:s" json:"prom_for_duration"`
	RuleConfig       string `gorm:"column:rule_config;type:text;not null;comment:rule_config" json:"rule_config"`
	PromQl           string `gorm:"column:prom_ql;type:text;not null;comment:promql" json:"prom_ql"`
	PromEvalInterval int64  `gorm:"column:prom_eval_interval;type:int(11);not null;comment:evaluate interval" json:"prom_eval_interval"`
	EnableStime      string `gorm:"column:enable_stime;type:varchar(255);not null;default:00:00" json:"enable_stime"`
	EnableEtime      string `gorm:"column:enable_etime;type:varchar(255);not null;default:23:59" json:"enable_etime"`
	EnableDaysOfWeek string `gorm:"column:enable_days_of_week;type:varchar(255);not null;comment:split by space: 0 1 2 3 4 5 6" json:"enable_days_of_week"`
	EnableInBg       int64  `gorm:"column:enable_in_bg;type:tinyint(1);not null;comment:1: only this bg 0: global" json:"enable_in_bg"`
	NotifyRecovered  int64  `gorm:"column:notify_recovered;type:tinyint(1);not null;comment:whether notify when recovery" json:"notify_recovered"`
	NotifyChannels   string `gorm:"column:notify_channels;type:varchar(255);not null;comment:split by space: sms voice email dingtalk wecom" json:"notify_channels"`
	NotifyGroups     string `gorm:"column:notify_groups;type:varchar(255);not null;comment:split by space: 233 43" json:"notify_groups"`
	NotifyRepeatStep int64  `gorm:"column:notify_repeat_step;type:int(11);not null;comment:unit: min" json:"notify_repeat_step"`
	NotifyMaxNumber  int64  `gorm:"column:notify_max_number;type:int(11);not null" json:"notify_max_number"`
	RecoverDuration  int64  `gorm:"column:recover_duration;type:int(11);not null;comment:unit: s" json:"recover_duration"`
	Callbacks        string `gorm:"column:callbacks;type:varchar(255);not null;comment:split by space: http://a.com/api/x http://a.com/api/y" json:"callbacks"`
	RunbookURL       string `gorm:"column:runbook_url;type:varchar(255)" json:"runbook_url"`
	AppendTags       string `gorm:"column:append_tags;type:varchar(255);not null;comment:split by space: service=n9e mod=api" json:"append_tags"`
	Annotations      string `gorm:"column:annotations;type:text;not null;comment:annotations" json:"annotations"`
	CreateAt         int64  `gorm:"column:create_at;type:bigint(20);not null" json:"create_at"`
	CreateBy         string `gorm:"column:create_by;type:varchar(64);not null" json:"create_by"`
	UpdateAt         int64  `gorm:"column:update_at;type:bigint(20);not null" json:"update_at"`
	UpdateBy         string `gorm:"column:update_by;type:varchar(64);not null" json:"update_by"`
}

// TableName AlertRule's table name
func (*AlertRule) TableName() string {
	return TableNameAlertRule
}
