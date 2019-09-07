package notifyctl

import (
	"fmt"
	"time"

	"git.funxdata.com/brisk/brisk/gerr"
	"git.funxdata.com/brisk/commons/format"
	"git.funxdata.com/funxcloud/funxcloud/core/fx"
	"git.funxdata.com/funxcloud/funxcloud/core/md"
	"git.funxdata.com/funxcloud/funxcloud/core/sysconfig"
	"git.funxdata.com/funxcloud/funxcloud/pkgs/oa/models/notify"
	"git.funxdata.com/funxcloud/funxcloud/types"
	"github.com/funxdata/wechat/template"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// notifyBusinessDailyData 经营日报表推送数据
func notifyBusinessDailyData(openid string) {
	tplid := sysconfig.Get().WechatEmployeeNotifyTpl.TaskRemind
	data := map[string]*template.DataItem{
		"first":    &template.DataItem{Value: "经营分析日报表"},
		"keyword1": &template.DataItem{Value: fmt.Sprintf("x月x日-报表")},
		"keyword2": &template.DataItem{Value: ""},
		"keyword3": &template.DataItem{Value: time.Now().Format(format.TimeLayout)},
		"keyword4": &template.DataItem{Value: ""},
		"remark":   &template.DataItem{Value: ""},
	}

	msg := &template.Message{
		ToUser:     openid,
		TemplateID: tplid,
		Color:      "black",
		Data:       data,
	}
	err := notify.SendEmployeeTpl(msg, "pages/analysis/business-date/index")
	if err != nil {
		logrus.Errorf("send notify failed: %v", err)
	}
}

// notifyOperationDailyData 运营日报表推送数据
func notifyOperationDailyData(openid string) {
	tplid := sysconfig.Get().WechatEmployeeNotifyTpl.TaskRemind
	data := map[string]*template.DataItem{
		"first":    &template.DataItem{Value: "运营分析日报表"},
		"keyword1": &template.DataItem{Value: fmt.Sprintf("x月x日-报表")},
		"keyword2": &template.DataItem{Value: ""},
		"keyword3": &template.DataItem{Value: time.Now().Format(format.TimeLayout)},
		"keyword4": &template.DataItem{Value: ""},
		"remark":   &template.DataItem{Value: ""},
	}

	msg := &template.Message{
		ToUser:     openid,
		TemplateID: tplid,
		Color:      "black",
		Data:       data,
	}
	err := notify.SendEmployeeTpl(msg, "pages/analysis/operate-date/index")
	if err != nil {
		logrus.Errorf("send notify failed: %v", err)
	}
}

// notifyBusinessMonthlyData 经营月报表推送数据
func notifyBusinessMonthlyData(openid string) {
	tplid := sysconfig.Get().WechatEmployeeNotifyTpl.TaskRemind
	data := map[string]*template.DataItem{
		"first":    &template.DataItem{Value: "经营分析月报表"},
		"keyword1": &template.DataItem{Value: fmt.Sprintf("x月x日-报表")},
		"keyword2": &template.DataItem{Value: ""},
		"keyword3": &template.DataItem{Value: time.Now().Format(format.TimeLayout)},
		"keyword4": &template.DataItem{Value: ""},
		"remark":   &template.DataItem{Value: ""},
	}

	msg := &template.Message{
		ToUser:     openid,
		TemplateID: tplid,
		Color:      "black",
		Data:       data,
	}
	err := notify.SendEmployeeTpl(msg, "pages/analysis/business-month/index")
	if err != nil {
		logrus.Errorf("send notify failed: %v", err)
	}
}

// notifyOperationMonthlyData 运营月报表推送数据
func notifyOperationMonthlyData(openid string) {
	tplid := sysconfig.Get().WechatEmployeeNotifyTpl.TaskRemind
	data := map[string]*template.DataItem{
		"first":    &template.DataItem{Value: "运营分析月报表"},
		"keyword1": &template.DataItem{Value: fmt.Sprintf("x月x日-报表")},
		"keyword2": &template.DataItem{Value: ""},
		"keyword3": &template.DataItem{Value: time.Now().Format(format.TimeLayout)},
		"keyword4": &template.DataItem{Value: ""},
		"remark":   &template.DataItem{Value: ""},
	}

	msg := &template.Message{
		ToUser:     openid,
		TemplateID: tplid,
		Color:      "black",
		Data:       data,
	}
	err := notify.SendEmployeeTpl(msg, "pages/analysis/operate-month/index")
	if err != nil {
		logrus.Errorf("send notify failed: %v", err)
	}
}

// Push 推送给已经订阅的用户
func Push(ctx *gin.Context) {
	var (
		businessDailyOpenIDs  []string
		businessMonthOpenIDs  []string
		operationDailyOpenIDs []string
		operationMonthOpenIDs []string
	)

	// 经营分析日报表
	sql := `boss_res_employee_store.position_id IN 
	(SELECT boss_rel_position_message.receiver_id FROM boss_rel_position_message WHERE
		boss_rel_position_message.receiver_type = ? AND boss_rel_position_message.message_id IN (
		(SELECT boss_message_bullet_topic.id from boss_message_bullet_topic WHERE push_wechat = ? AND 
			NAME= ?)) AND boss_employee.employee_mp_openid IS NOT NULL`

	err := md.BizDB(ctx).Table((*types.Employee)(nil).TableName()).
		Joins("LEFT JOIN boss_res_employee_store ON boss_employee.id = boss_res_employee_store.employee_id AND boss_res_employee_store.deleted_at IS NULL").
		Where(sql, "POSITION", true, "BUSINESS_DAILY_REPORT_NOTIFY").
		Pluck("boss_employee.employee_mp_openid", &businessDailyOpenIDs).
		Error
	if err != nil {
		fx.BadRequest(ctx, err.Error)
	}

	for _, openId := 

	// {
	// 	// 经营分析日报表
	// 	// employee.EmployeeMpOpenid
	// 	// notifyData(employee.EmployeeMpOpenid)
	// }
	// {
	// 	// 经营分析日报表
	// }
	// {
	// 	// 经营分析日报表
	// }
	// {
	// 	// 经营分析日报表
	// }
}

// ListBulletTopic 消息通知列表
func ListBulletTopic(ctx *gin.Context) {
	logger := md.Logger(ctx)
	l, err := notify.List(md.BizDB(ctx))
	if err != nil {
		logger.Errorf("got message bullet topic failed : %v", err)
		fx.Err(ctx, gerr.InternalError("查询消息通知列表失败"))
		return
	}
	fx.Succ(ctx, l)
}

// DetailBulletTopic 消息通知设置详情
func DetailBulletTopic(ctx *gin.Context) {
	id := fx.ID(ctx)
	out, err := notify.DetailBulletTopic(md.BizDB(ctx), id)
	if err != nil {
		fx.BadRequest(ctx, err.Error())
		return
	}
	fx.Succ(ctx, out)
}

// UpdateBulletTopic 消息推送修改
func UpdateBulletTopic(ctx *gin.Context) {
	in := &notify.UpdateOption{}
	if err := ctx.ShouldBindJSON(in); err != nil {
		fx.Err(ctx, err)
		return
	}
	in.MessageTopicID = fx.ID(ctx)
	ret, err := notify.UpdateBulletTopic(md.BizDB(ctx), in)
	if err != nil {
		fx.BadRequest(ctx, err.Error())
		return
	}
	fx.Succ(ctx, ret)
}

// ListBulletShows 职位要推送的消息列表
func ListBulletShows(ctx *gin.Context) {
	employeeID := md.EmployeeID(ctx)
	if employeeID < 1 {
		fx.BadRequest(ctx, "未找到员工信息")
		return
	}
	ret, err := notify.ListBulletShows(md.BizDB(ctx), employeeID)
	if err != nil {
		fx.BadRequest(ctx, err.Error())
		return
	}
	fx.Succ(ctx, ret)
}
