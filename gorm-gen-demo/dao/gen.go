// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q               = new(Query)
	AlertAggrView   *alertAggrView
	AlertCurEvent   *alertCurEvent
	AlertHisEvent   *alertHisEvent
	AlertMute       *alertMute
	AlertRule       *alertRule
	AlertSubscribe  *alertSubscribe
	AlertingEngine  *alertingEngine
	AppServer       *appServer
	Board           *board
	BoardPayload    *boardPayload
	BuiltinCate     *builtinCate
	BusiGroup       *busiGroup
	BusiGroupMember *busiGroupMember
	Chart           *chart
	ChartGroup      *chartGroup
	ChartShare      *chartShare
	Config          *config
	Dashboard       *dashboard
	Datasource      *datasource
	MetricView      *metricView
	NotifyTpl       *notifyTpl
	RecordingRule   *recordingRule
	Role            *role
	RoleOperation   *roleOperation
	SsoConfig       *ssoConfig
	Target          *target
	TaskRecord      *taskRecord
	TaskTpl         *taskTpl
	TaskTplHost     *taskTplHost
	User            *user
	UserGroup       *userGroup
	UserGroupMember *userGroupMember
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AlertAggrView = &Q.AlertAggrView
	AlertCurEvent = &Q.AlertCurEvent
	AlertHisEvent = &Q.AlertHisEvent
	AlertMute = &Q.AlertMute
	AlertRule = &Q.AlertRule
	AlertSubscribe = &Q.AlertSubscribe
	AlertingEngine = &Q.AlertingEngine
	AppServer = &Q.AppServer
	Board = &Q.Board
	BoardPayload = &Q.BoardPayload
	BuiltinCate = &Q.BuiltinCate
	BusiGroup = &Q.BusiGroup
	BusiGroupMember = &Q.BusiGroupMember
	Chart = &Q.Chart
	ChartGroup = &Q.ChartGroup
	ChartShare = &Q.ChartShare
	Config = &Q.Config
	Dashboard = &Q.Dashboard
	Datasource = &Q.Datasource
	MetricView = &Q.MetricView
	NotifyTpl = &Q.NotifyTpl
	RecordingRule = &Q.RecordingRule
	Role = &Q.Role
	RoleOperation = &Q.RoleOperation
	SsoConfig = &Q.SsoConfig
	Target = &Q.Target
	TaskRecord = &Q.TaskRecord
	TaskTpl = &Q.TaskTpl
	TaskTplHost = &Q.TaskTplHost
	User = &Q.User
	UserGroup = &Q.UserGroup
	UserGroupMember = &Q.UserGroupMember
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:              db,
		AlertAggrView:   newAlertAggrView(db, opts...),
		AlertCurEvent:   newAlertCurEvent(db, opts...),
		AlertHisEvent:   newAlertHisEvent(db, opts...),
		AlertMute:       newAlertMute(db, opts...),
		AlertRule:       newAlertRule(db, opts...),
		AlertSubscribe:  newAlertSubscribe(db, opts...),
		AlertingEngine:  newAlertingEngine(db, opts...),
		AppServer:       newAppServer(db, opts...),
		Board:           newBoard(db, opts...),
		BoardPayload:    newBoardPayload(db, opts...),
		BuiltinCate:     newBuiltinCate(db, opts...),
		BusiGroup:       newBusiGroup(db, opts...),
		BusiGroupMember: newBusiGroupMember(db, opts...),
		Chart:           newChart(db, opts...),
		ChartGroup:      newChartGroup(db, opts...),
		ChartShare:      newChartShare(db, opts...),
		Config:          newConfig(db, opts...),
		Dashboard:       newDashboard(db, opts...),
		Datasource:      newDatasource(db, opts...),
		MetricView:      newMetricView(db, opts...),
		NotifyTpl:       newNotifyTpl(db, opts...),
		RecordingRule:   newRecordingRule(db, opts...),
		Role:            newRole(db, opts...),
		RoleOperation:   newRoleOperation(db, opts...),
		SsoConfig:       newSsoConfig(db, opts...),
		Target:          newTarget(db, opts...),
		TaskRecord:      newTaskRecord(db, opts...),
		TaskTpl:         newTaskTpl(db, opts...),
		TaskTplHost:     newTaskTplHost(db, opts...),
		User:            newUser(db, opts...),
		UserGroup:       newUserGroup(db, opts...),
		UserGroupMember: newUserGroupMember(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AlertAggrView   alertAggrView
	AlertCurEvent   alertCurEvent
	AlertHisEvent   alertHisEvent
	AlertMute       alertMute
	AlertRule       alertRule
	AlertSubscribe  alertSubscribe
	AlertingEngine  alertingEngine
	AppServer       appServer
	Board           board
	BoardPayload    boardPayload
	BuiltinCate     builtinCate
	BusiGroup       busiGroup
	BusiGroupMember busiGroupMember
	Chart           chart
	ChartGroup      chartGroup
	ChartShare      chartShare
	Config          config
	Dashboard       dashboard
	Datasource      datasource
	MetricView      metricView
	NotifyTpl       notifyTpl
	RecordingRule   recordingRule
	Role            role
	RoleOperation   roleOperation
	SsoConfig       ssoConfig
	Target          target
	TaskRecord      taskRecord
	TaskTpl         taskTpl
	TaskTplHost     taskTplHost
	User            user
	UserGroup       userGroup
	UserGroupMember userGroupMember
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		AlertAggrView:   q.AlertAggrView.clone(db),
		AlertCurEvent:   q.AlertCurEvent.clone(db),
		AlertHisEvent:   q.AlertHisEvent.clone(db),
		AlertMute:       q.AlertMute.clone(db),
		AlertRule:       q.AlertRule.clone(db),
		AlertSubscribe:  q.AlertSubscribe.clone(db),
		AlertingEngine:  q.AlertingEngine.clone(db),
		AppServer:       q.AppServer.clone(db),
		Board:           q.Board.clone(db),
		BoardPayload:    q.BoardPayload.clone(db),
		BuiltinCate:     q.BuiltinCate.clone(db),
		BusiGroup:       q.BusiGroup.clone(db),
		BusiGroupMember: q.BusiGroupMember.clone(db),
		Chart:           q.Chart.clone(db),
		ChartGroup:      q.ChartGroup.clone(db),
		ChartShare:      q.ChartShare.clone(db),
		Config:          q.Config.clone(db),
		Dashboard:       q.Dashboard.clone(db),
		Datasource:      q.Datasource.clone(db),
		MetricView:      q.MetricView.clone(db),
		NotifyTpl:       q.NotifyTpl.clone(db),
		RecordingRule:   q.RecordingRule.clone(db),
		Role:            q.Role.clone(db),
		RoleOperation:   q.RoleOperation.clone(db),
		SsoConfig:       q.SsoConfig.clone(db),
		Target:          q.Target.clone(db),
		TaskRecord:      q.TaskRecord.clone(db),
		TaskTpl:         q.TaskTpl.clone(db),
		TaskTplHost:     q.TaskTplHost.clone(db),
		User:            q.User.clone(db),
		UserGroup:       q.UserGroup.clone(db),
		UserGroupMember: q.UserGroupMember.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		AlertAggrView:   q.AlertAggrView.replaceDB(db),
		AlertCurEvent:   q.AlertCurEvent.replaceDB(db),
		AlertHisEvent:   q.AlertHisEvent.replaceDB(db),
		AlertMute:       q.AlertMute.replaceDB(db),
		AlertRule:       q.AlertRule.replaceDB(db),
		AlertSubscribe:  q.AlertSubscribe.replaceDB(db),
		AlertingEngine:  q.AlertingEngine.replaceDB(db),
		AppServer:       q.AppServer.replaceDB(db),
		Board:           q.Board.replaceDB(db),
		BoardPayload:    q.BoardPayload.replaceDB(db),
		BuiltinCate:     q.BuiltinCate.replaceDB(db),
		BusiGroup:       q.BusiGroup.replaceDB(db),
		BusiGroupMember: q.BusiGroupMember.replaceDB(db),
		Chart:           q.Chart.replaceDB(db),
		ChartGroup:      q.ChartGroup.replaceDB(db),
		ChartShare:      q.ChartShare.replaceDB(db),
		Config:          q.Config.replaceDB(db),
		Dashboard:       q.Dashboard.replaceDB(db),
		Datasource:      q.Datasource.replaceDB(db),
		MetricView:      q.MetricView.replaceDB(db),
		NotifyTpl:       q.NotifyTpl.replaceDB(db),
		RecordingRule:   q.RecordingRule.replaceDB(db),
		Role:            q.Role.replaceDB(db),
		RoleOperation:   q.RoleOperation.replaceDB(db),
		SsoConfig:       q.SsoConfig.replaceDB(db),
		Target:          q.Target.replaceDB(db),
		TaskRecord:      q.TaskRecord.replaceDB(db),
		TaskTpl:         q.TaskTpl.replaceDB(db),
		TaskTplHost:     q.TaskTplHost.replaceDB(db),
		User:            q.User.replaceDB(db),
		UserGroup:       q.UserGroup.replaceDB(db),
		UserGroupMember: q.UserGroupMember.replaceDB(db),
	}
}

type queryCtx struct {
	AlertAggrView   IAlertAggrViewDo
	AlertCurEvent   IAlertCurEventDo
	AlertHisEvent   IAlertHisEventDo
	AlertMute       IAlertMuteDo
	AlertRule       IAlertRuleDo
	AlertSubscribe  IAlertSubscribeDo
	AlertingEngine  IAlertingEngineDo
	AppServer       IAppServerDo
	Board           IBoardDo
	BoardPayload    IBoardPayloadDo
	BuiltinCate     IBuiltinCateDo
	BusiGroup       IBusiGroupDo
	BusiGroupMember IBusiGroupMemberDo
	Chart           IChartDo
	ChartGroup      IChartGroupDo
	ChartShare      IChartShareDo
	Config          IConfigDo
	Dashboard       IDashboardDo
	Datasource      IDatasourceDo
	MetricView      IMetricViewDo
	NotifyTpl       INotifyTplDo
	RecordingRule   IRecordingRuleDo
	Role            IRoleDo
	RoleOperation   IRoleOperationDo
	SsoConfig       ISsoConfigDo
	Target          ITargetDo
	TaskRecord      ITaskRecordDo
	TaskTpl         ITaskTplDo
	TaskTplHost     ITaskTplHostDo
	User            IUserDo
	UserGroup       IUserGroupDo
	UserGroupMember IUserGroupMemberDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AlertAggrView:   q.AlertAggrView.WithContext(ctx),
		AlertCurEvent:   q.AlertCurEvent.WithContext(ctx),
		AlertHisEvent:   q.AlertHisEvent.WithContext(ctx),
		AlertMute:       q.AlertMute.WithContext(ctx),
		AlertRule:       q.AlertRule.WithContext(ctx),
		AlertSubscribe:  q.AlertSubscribe.WithContext(ctx),
		AlertingEngine:  q.AlertingEngine.WithContext(ctx),
		AppServer:       q.AppServer.WithContext(ctx),
		Board:           q.Board.WithContext(ctx),
		BoardPayload:    q.BoardPayload.WithContext(ctx),
		BuiltinCate:     q.BuiltinCate.WithContext(ctx),
		BusiGroup:       q.BusiGroup.WithContext(ctx),
		BusiGroupMember: q.BusiGroupMember.WithContext(ctx),
		Chart:           q.Chart.WithContext(ctx),
		ChartGroup:      q.ChartGroup.WithContext(ctx),
		ChartShare:      q.ChartShare.WithContext(ctx),
		Config:          q.Config.WithContext(ctx),
		Dashboard:       q.Dashboard.WithContext(ctx),
		Datasource:      q.Datasource.WithContext(ctx),
		MetricView:      q.MetricView.WithContext(ctx),
		NotifyTpl:       q.NotifyTpl.WithContext(ctx),
		RecordingRule:   q.RecordingRule.WithContext(ctx),
		Role:            q.Role.WithContext(ctx),
		RoleOperation:   q.RoleOperation.WithContext(ctx),
		SsoConfig:       q.SsoConfig.WithContext(ctx),
		Target:          q.Target.WithContext(ctx),
		TaskRecord:      q.TaskRecord.WithContext(ctx),
		TaskTpl:         q.TaskTpl.WithContext(ctx),
		TaskTplHost:     q.TaskTplHost.WithContext(ctx),
		User:            q.User.WithContext(ctx),
		UserGroup:       q.UserGroup.WithContext(ctx),
		UserGroupMember: q.UserGroupMember.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
