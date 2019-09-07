// 水电账单
package meters

import (
	"fmt"
	"strconv"
	"time"

	"git.funxdata.com/funxcloud/commons/date"
	"git.funxdata.com/funxcloud/commons/excel"
	"git.funxdata.com/funxcloud/commons/gerr"
	"git.funxdata.com/funxcloud/commons/gtype/imagex"
	"git.funxdata.com/funxcloud/commons/gtype/money"
	"git.funxdata.com/funxcloud/commons/mathx"
	"git.funxdata.com/funxcloud/commons/pages"
	"git.funxdata.com/funxcloud/funxcloud/core/global"
	"git.funxdata.com/funxcloud/funxcloud/pkgs/as/models/orderv3"
	"git.funxdata.com/funxcloud/funxcloud/pkgs/eam/models/devices"
	"git.funxdata.com/funxcloud/funxcloud/pkgs/eam/models/room"
	"git.funxdata.com/funxcloud/funxcloud/pkgs/pms/models/resident"
	"git.funxdata.com/funxcloud/funxcloud/types"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

// ReadingOrderListItem 读数账单列表项
type ReadingOrderListItem struct {
	*types.MeterReadingTransfer
	StoreName         string                        `gorm:"column:store_name" json:"store_name"`
	ResidentName      string                        `gorm:"column:resident_name" json:"resident_name"`
	RoomNumber        string                        `gorm:"column:room_number" json:"room_number"`
	LastTime          time.Time                     `gorm:"column:last_time" json:"last_time"`
	LastReading       float64                       `gorm:"column:last_reading" json:"last_reading"`
	Difference        float64                       `gorm:"column:difference" json:"difference"`
	Fee               money.Money                   `gorm:"column:fee" json:"fee"`
	Status            string                        `gorm:"column:status" json:"status"`
	ThisOrderStatus   types.MeterReadingOrderStatus `gorm:"column:"this_order_status"`
	LastStatus        types.MeterReadingStatus      `gorm:"column:last_status"`
	ThisStatus        types.MeterReadingStatus      `gorm:"column:this_status"`
	LastReadingSource string                        `gorm:"column:last_reading_source" json:"last_reading_source"`
	ThisReadingSource string                        `gorm:"column:this_reading_source" json:"this_reading_source"`
	LastImage         imagex.Image                  `gorm:"column:last_image" json:"last_image"`
	ThisImage         imagex.Image                  `gorm:"column:this_image" json:"this_image"`

	RoomGasPrice         money.Money `gorm:"column:gas_price"`
	RoomColdWaterPrice   money.Money `gorm:"column:cold_water_price"`
	RoomHotWaterPrice    money.Money `gorm:"column:hot_water_price"`
	RoomElectricityPrice money.Money `gorm:"column:elec_price"`
}

// ReadingOrderList 读数账单列表
type ReadingOrderList struct {
	pages.Page
	Items []*ReadingOrderListItem `json:"items"`
}

// ReadingOrderListOption 水电读数账单列表提交参数信息
type ReadingOrderListOption struct {
	pages.Page
	StoreID     uint                          `json:"store_id" form:"store_id"`
	StoreIDs    []uint                        `json:"store_ids[]" form:"store_ids[]"`
	MType       types.DeviceType              `json:"type" form:"type"`
	OrderStatus types.MeterReadingOrderStatus `json:"order_status" form:"status"`
	Year        int                           `json:"year" form:"year"`
	Month       time.Month                    `json:"month" form:"month"`
	RoomNumber  string                        `json:"number" form:"number"`
}

// ListOrders 获取账单列表信息
func ListOrders(db *gorm.DB, logger *logrus.Entry, opt *ReadingOrderListOption) (*ReadingOrderList, error) {
	out := &ReadingOrderList{Items: []*ReadingOrderListItem{}}
	query := db.Table((*types.MeterReadingTransfer)(nil).TableName()).
		Joins("left join boss_store on boss_store.id = boss_meter_reading_transfer.store_id").
		Joins("left join boss_room_union on boss_room_union.id = boss_meter_reading_transfer.room_id").
		Joins("left join boss_resident on boss_resident.id = boss_meter_reading_transfer.resident_id").
		Where("boss_meter_reading_transfer.deleted_at is NULL")

	if opt.StoreID != 0 {
		query = query.Where("boss_meter_reading_transfer.store_id = ?", opt.StoreID)
	}
	if len(opt.StoreIDs) > 0 {
		query = query.Where("boss_meter_reading_transfer.store_id in (?)", opt.StoreIDs)
	}
	if opt.MType != "" {
		query = query.Where("boss_meter_reading_transfer.type = ?", opt.MType)
	}
	if opt.RoomNumber != "" {
		query = query.Where("boss_meter_reading_transfer.number like ?", "%"+opt.RoomNumber+"%")
	}
	if opt.OrderStatus == types.MeterReadingOrderStatusHasorder {
		query = query.Where("boss_meter_reading_transfer.order_status in (?)", []types.MeterReadingOrderStatus{
			types.MeterReadingOrderStatusHasorder,
		})
	} else if opt.OrderStatus == types.MeterReadingOrderStatusNoorder {
		query = query.Where("boss_meter_reading_transfer.order_status in (?)", []types.MeterReadingOrderStatus{
			types.MeterReadingOrderStatusNoorder,
			types.MeterReadingOrderStatusNoreading,
		})
	}
	if opt.Year != 0 && opt.Month != 0 {
		query = query.Where("boss_meter_reading_transfer.year = ? and boss_meter_reading_transfer.month = ?", opt.Year, opt.Month)
	}
	query = query.Where("boss_meter_reading_transfer.order_status <> ?", types.MeterReadingOrderStatusInitReading).
		Where("boss_meter_reading_transfer.is_valid <> ?", types.MeterReadingInValid).
		Select("boss_meter_reading_transfer.*").
		Order("boss_meter_reading_transfer.this_time desc, boss_meter_reading_transfer.id desc")
	if err := query.Count(&out.Count).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	if out.Count == 0 {
		return out, nil
	}

	// *types.MeterReadingTransfer
	// StoreName         string       `json:"store_name"`
	// ResidentName      string       `json:"resident_name"`
	// RoomNumber        string       `json:"room_number"`
	// LastTime          time.Time    `json:"last_time"`
	// LastReading       float64      `json:"last_reading"`
	// Difference        float64      `json:"difference"`
	// Fee               money.Money  `json:"fee"`
	// Status            string       `json:"status"`
	// LastReadingSource string       `json:"last_reading_source"`
	// ThisReadingSource string       `json:"this_reading_source"`
	// LastImage         imagex.Image `json:"last_image"`
	// ThisImage         imagex.Image `json:"this_image"`

	if err := query.
		Joins("left join boss_meter_reading_transfer as b on b.room_id = boss_meter_reading_transfer.room_id and" +
			" b.type = boss_meter_reading_transfer.type and" +
			" b.is_valid = 'HAS_VALID'").
		Order("b.this_time desc").
		Group("boss_meter_reading_transfer.id").
		Select("boss_meter_reading_transfer.*" +
			",boss_store.name as store_name" +
			",boss_resident.name as resident_name, boss_room_union.number as room_number" +
			",boss_room_union.gas_price as gas_price" +
			",boss_room_union.hot_water_price as hot_water_price" +
			",boss_room_union.cold_water_price as cold_water_price" +
			",b.this_time as last_time" +
			",b.this_reading as last_reading" +
			",b.status as last_status" +
			",boss_meter_reading_transfer.status as this_status" +
			",boss_meter_reading_transfer.order_status as this_order_status",
		).Find(&out.Items).Error; err != nil {
		return nil, err
	}

	outStatusMap := map[types.MeterReadingOrderStatus]string{
		types.MeterReadingOrderStatusNoreading:    "录入中",
		types.MeterReadingOrderStatusNoresident:   "不需生成账单",
		types.MeterReadingOrderStatusNoorder:      "未生成账单",
		types.MeterReadingOrderStatusHasorder:     "已生成",
		types.MeterReadingOrderStatusZeroOrder:    "不需生成账单",
		types.MeterReadingOrderStatusInitReading:  "不需生成账单",
		types.MeterReadingOrderStatusErrorReading: "错误读数",
	}

	sourceMap := map[types.MeterReadingStatus]string{
		types.MeterReadingStatusInit:      "初始读数",
		types.MeterReadingStatusChangeOld: "换表新读数",
		types.MeterReadingStatusChangeNew: "换表旧读数",
		types.MeterReadingStatusFullStart: "爆表旧读数",
		types.MeterReadingStatusFullEnd:   "爆表新读数",
		types.MeterReadingStatusNewRent:   "新入住读数",
		types.MeterReadingStatusNormal:    "日常读数",
	}

	for _, item := range out.Items {
		price := money.Money(0)
		switch item.Type {
		case types.DeviceTypeHotWater:
			price = item.RoomHotWaterPrice
		case types.DeviceTypeColdWater:
			price = item.RoomColdWaterPrice
		case types.DeviceTypeElectric:
			price = item.RoomElectricityPrice
		case types.DeviceTypeGas:
			price = item.RoomGasPrice
		}
		item.Difference = item.ThisReading - item.LastReading
		item.Fee = price.Mul(item.ThisReading - item.LastReading)
		item.LastReadingSource = sourceMap[item.LastStatus]
		item.ThisReadingSource = sourceMap[item.ThisStatus]
		item.Status = outStatusMap[item.ThisOrderStatus]
	}

	return out, nil

	// if err:=query.Select(
	// 	"boss_meter_reading_transfer.*,boss_store.name as store_name,boss_resident.name as resident_name,boss_room_union.number as room_number,"+
	// 	"boss_room_union.gas_price as room_gas_price,b.this_time as b_this_time,b.this_reading as b_this_reading").
	// 	Joins("left join boss_meter_reading_transfer as b on b.room_id = boss_meter_reading_transfer.room_id"+
	// 	"and b.type=boss_meter_reading_transfer.type and b.is_valid='HAS_VALID'").
	// 	Order("b.this_time desc").
	// 	Group("boss_meter_reading_transfer.id").
	// 	Find(&out.Items).
	// 	Error; err!=nil{
	// 	return nil,err
	// }

	// queryOpt := &ReadingQueryOption{
	// 	StoreID:    opt.StoreID,
	// 	StoreIDs:   opt.StoreIDs,
	// 	Mtype:      opt.MType,
	// 	RoomNumber: opt.RoomNumber,
	// }
	// if opt.OrderStatus == types.MeterReadingOrderStatusHasorder {
	// 	queryOpt.OrderStatus = []types.MeterReadingOrderStatus{
	// 		types.MeterReadingOrderStatusHasorder,
	// 	}
	// } else if opt.OrderStatus == types.MeterReadingOrderStatusNoorder {
	// 	queryOpt.OrderStatus = []types.MeterReadingOrderStatus{
	// 		types.MeterReadingOrderStatusNoorder,
	// 		types.MeterReadingOrderStatusNoreading,
	// 	}
	// }

	// out := &ReadingOrderList{Items: []*ReadingOrderListItem{}}
	// query := queryOpt.Query(db)
	// if opt.Year != 0 && opt.Month != 0 {
	// 	query = query.Where("year = ? and month = ?", opt.Year, int(opt.Month))
	// }
	// query = query.Where("order_status <> ?", types.MeterReadingOrderStatusInitReading).
	// 	Where("is_valid <> ?", types.MeterReadingInValid).
	// 	Select("boss_meter_reading_transfer.*").
	// 	Order("this_time desc, id desc")
	// if err := query.Count(&out.Count).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
	// 	return nil, err
	// }
	// if out.Count == 0 {
	// 	return out, nil
	// }

	// nowReading := []*types.MeterReadingTransfer{}
	// err := query.Offset(opt.Offset).Limit(opt.Limit).Scan(&nowReading).Error
	// if err != nil {
	// 	return nil, err
	// }

	// outStatusMap := map[types.MeterReadingOrderStatus]string{
	// 	types.MeterReadingOrderStatusNoreading:    "录入中",
	// 	types.MeterReadingOrderStatusNoresident:   "不需生成账单",
	// 	types.MeterReadingOrderStatusNoorder:      "未生成账单",
	// 	types.MeterReadingOrderStatusHasorder:     "已生成",
	// 	types.MeterReadingOrderStatusZeroOrder:    "不需生成账单",
	// 	types.MeterReadingOrderStatusInitReading:  "不需生成账单",
	// 	types.MeterReadingOrderStatusErrorReading: "错误读数",
	// }

	// sourceMap := map[types.MeterReadingStatus]string{
	// 	types.MeterReadingStatusInit:      "初始读数",
	// 	types.MeterReadingStatusChangeOld: "换表新读数",
	// 	types.MeterReadingStatusChangeNew: "换表旧读数",
	// 	types.MeterReadingStatusFullStart: "爆表旧读数",
	// 	types.MeterReadingStatusFullEnd:   "爆表新读数",
	// 	types.MeterReadingStatusNewRent:   "新入住读数",
	// 	types.MeterReadingStatusNormal:    "日常读数",
	// }

	// for _, item := range nowReading {
	// opt := &ReadingQueryOption{
	// 	RoomID: item.RoomID,
	// 	Mtype:  item.Type,
	// 	IsValid: []types.MeterReadingValid{
	// 		types.MeterReadingHasValid,
	// 	},
	// }
	// lastReading, err := opt.GetValidLastReading(db, item)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	feeOption := &FeeOption{
	// 		Mtype:          item.Type,
	// 		LastReading:    lastReading.ThisReading,
	// 		CurrentReading: item.ThisReading,
	// 		RoomID:         item.RoomID,
	// 	}

	// 	var (
	// 		testReading       types.MeterReadingTransfer
	// 		fee               money.Money
	// 		diff              float64
	// 		outStatus         = outStatusMap[item.OrderStatus]
	// 		lastReadingSource = sourceMap[lastReading.Status]
	// 		thisReadingSource = sourceMap[item.Status]
	// 	)
	// 	if *lastReading == testReading {
	// 		fee = 0
	// 		diff = 0
	// 	} else {
	// result, err := ReadingFee(db, feeOption)
	// 		fee = result
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		diff = item.ThisReading - lastReading.ThisReading
	// 	}
	// 	s, err := store.GetStoreByID(db, item.StoreID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	room, err := room.Get(db, item.RoomID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	listItem := &ReadingOrderListItem{
	// 		MeterReadingTransfer: item,
	// 		StoreName:            s.Name,
	// 		// ResidentName:         r.Name,
	// 		RoomNumber:        room.Number,
	// 		LastTime:          lastReading.ThisTime,
	// 		LastReading:       lastReading.ThisReading,
	// 		Difference:        diff,
	// 		Fee:               fee,
	// 		Status:            outStatus,
	// 		LastReadingSource: lastReadingSource,
	// 		ThisReadingSource: thisReadingSource,
	// 		LastImage:         lastReading.Image,
	// 		ThisImage:         item.Image,
	// 	}
	// 	if item.ResidentID != nil && *item.ResidentID != 0 {
	// 		r, err := resident.GetResidentByID(db, *item.ResidentID)
	// 		if err != nil {
	// 			// return nil, err
	// 		} else {
	// 			listItem.ResidentName = r.Name
	// 		}
	// 	}
	// 	out.Items = append(out.Items, listItem)
	// }
	// return out, nil
}

//ExportReadingOrderTemplate 导出账单项
type ExportReadingOrderTemplate struct {
	StoreName    string      `xlsx:"0"`
	ResidentName string      `xlsx:"1"`
	RoomNumber   string      `xlsx:"2"`
	Type         string      `xlsx:"3"`
	LastTime     time.Time   `xlsx:"4"`
	LastReading  string      `xlsx:"5"`
	ThisTime     time.Time   `xlsx:"6"`
	ThisReading  string      `xlsx:"7"`
	Difference   float64     `xlsx:"8"`
	Weight       float64     `xlsx:"9"`
	Fee          money.Money `xlsx:"10"`
	Status       string      `xlsx:"11"`
}

// ExportListReadingOrder 导出水电账单列表
func ExportListReadingOrder(db *gorm.DB, logrus *logrus.Entry, opt *ReadingOrderListOption) (*xlsx.File, error) {
	opt.Offset = 0
	opt.Limit = 500
	rows := []interface{}{}

	for {
		orderList, err := ListOrders(db, logrus, opt)
		if err != nil {
			return nil, err
		} else if orderList.Count > 1500 {
			return nil, gerr.BadRequest("需要导出的记录过多，请添加筛选条件后重新导出")
		}
		for _, item := range orderList.Items {
			outItem := &ExportReadingOrderTemplate{
				StoreName:    item.StoreName,
				ResidentName: item.ResidentName,
				RoomNumber:   item.RoomNumber,
				Type:         types.DeviceTypeMap[item.Type],
				LastTime:     item.LastTime,
				LastReading:  strconv.FormatFloat(item.LastReading, 'f', 2, 64) + "(" + item.LastReadingSource + ")",
				ThisTime:     item.ThisTime,
				ThisReading:  strconv.FormatFloat(item.ThisReading, 'f', 2, 64) + "(" + item.ThisReadingSource + ")",
				Difference:   item.Difference,
				Weight:       item.Weight,
				Fee:          item.Fee,
				Status:       item.Status,
			}
			rows = append(rows, outItem)
		}
		orderList, err = ListOrders(db, logrus, opt)
		if err != nil {
			return nil, err
		}

		if opt.Offset+opt.Limit >= 5000 {
			logrus.Errorf("export too much records. %+v", orderList.Count)
			break
		} else if orderList.Count > opt.Offset+opt.Limit {
			opt.Offset += opt.Limit
		} else {
			break
		}
	}

	f, err := excel.NewWriter()
	if err != nil {
		return nil, err
	}
	f.SetTitleStr("水电账单列表-导出模板")
	headers := []*excel.Col{
		&excel.Col{
			Header: "门店名称", Width: 25,
		},
		&excel.Col{
			Header: "姓名", Width: 12,
		},
		&excel.Col{
			Header: "房间号", Width: 12,
		},
		&excel.Col{
			Header: "类型", Width: 12,
		},
		&excel.Col{
			Header: "上次抄表时间", Width: 25,
		},
		&excel.Col{
			Header: "上次读数", Width: 20,
		},
		&excel.Col{
			Header: "本次抄表时间", Width: 25,
		},
		&excel.Col{
			Header: "本次读数", Width: 20,
		},
		&excel.Col{
			Header: "差值", Width: 15,
		},
		&excel.Col{
			Header: "权重", Width: 12,
		},
		&excel.Col{
			Header: "费用", Width: 10,
		},
		&excel.Col{
			Header: "状态", Width: 12,
		},
	}
	f.SetData(rows, headers...)
	f.Save()
	return f.File, nil
}

// FeeOption 费用选择
type FeeOption struct {
	Mtype          types.DeviceType
	LastReading    float64
	CurrentReading float64
	RoomID         uint
}

// ReadingFee 计算设备读数费用
func ReadingFee(db *gorm.DB, opt *FeeOption) (money.Money, error) {
	price, err := devices.GetMeterUnitPrice(db, opt.RoomID, opt.Mtype)
	if err != nil {
		return 0, err
	}
	return price.Mul(opt.CurrentReading - opt.LastReading), nil
}

// getNormalPayCircleYearMonth 月度账单的支付周期
func getNormalPayCircleYearMonth(t time.Time) (int, time.Month) {
	if t.Day() >= 20 {
		nextmonth := date.NextMonthDay(t)
		return nextmonth.Year(), nextmonth.Month()
	}
	return t.Year(), t.Month()
}

// getNormalFeeCircleYearMonth 月度账单的费用周期，如果是10号之前，归为上个月的费用
func getNormalFeeCircleYearMonth(t time.Time) (int, time.Month) {
	if t.Day() <= 10 {
		lastmonth := date.MonthFirstDay(t).Add(time.Hour * (-48))
		return lastmonth.Year(), lastmonth.Month()
	}
	return t.Year(), t.Month()
}

// GenerateOrderOption 生成水电读数账单选择
type GenerateOrderOption struct {
	StoreID  uint             `json:"store_id" form:"store_id"`
	StoreIDs []uint           `json:"store_ids[]" form:"store_ids[]"`
	Mtype    types.DeviceType `json:"type" form:"type"`
	Year     int              `json:"year" form:"year" binding:"required"`
	Month    time.Month       `json:"month" form:"month" binding:"required"`
}

// OrderErrMsgItem 读数生成错误结果
type OrderErrMsgItem struct {
	ID          uint    `json:"id"`
	Error       string  `json:"error"`
	RoomNumber  string  `json:"room_number"`
	ThisReading float64 `json:"this_reading"`
	Weight      float64 `json:"weight"`
}

// OrderErrMsg 读数生成错误结果
type OrderErrMsg struct {
	Items []*OrderErrMsgItem `json:"items"`
	Count int
	Succ  int
	Fail  int
}

// CreateNormalReadingOrder 生成月度水电读数账单
func CreateNormalReadingOrder(db *gorm.DB, log *logrus.Entry, in *GenerateOrderOption) (*OrderErrMsg, error) {
	normalReadings, err := listNormalWaitValidReadings(db, in.StoreID, in.Mtype)
	if err != nil {
		log.Errorf("CreateNormalReadingOrder: can not listNormalWaitValidReadings, %s", err)
		return nil, err
	}

	ret := &OrderErrMsg{
		Items: []*OrderErrMsgItem{},
	}
	now := time.Now()

	for _, item := range normalReadings {
		ret.Count++
		err := doCreateNormalReadingOrder(db, log, item, now)
		if err != nil {
			ret.Fail++
			ret.Items = append(ret.Items, &OrderErrMsgItem{
				ID:          item.ID,
				Error:       item.Err,
				ThisReading: item.ThisReading,
				Weight:      item.Weight,
			})
		} else {
			ret.Succ++
		}
	}
	return ret, nil
}

// doCreateNormalReadingOrder 生成一条读数记录的账单并更新读数记录
func doCreateNormalReadingOrder(db *gorm.DB, log *logrus.Entry, r *types.MeterReadingTransfer, now time.Time) error {
	if r.ResidentID == nil || *r.ResidentID == 0 {
		updates := map[string]interface{}{
			"order_status": types.MeterReadingOrderStatusNoresident,
			"is_valid":     types.MeterReadingHasValid,
		}
		err := db.Table((*types.MeterReadingTransfer)(nil).TableName()).
			Where("id = ?", r.ID).
			Updates(updates).
			Error
		return err
	}

	fee, err := generateNormalMeterFee(db, log, r, now)
	if err != nil {
		r.OrderStatus = types.MeterReadingOrderStatusErrorReading
		r.Err = err.Error()
	} else if fee == nil {
		r.IsValid = types.MeterReadingHasValid
		r.OrderStatus = types.MeterReadingOrderStatusZeroOrder
	} else {
		r.IsValid = types.MeterReadingHasValid
		r.OrderStatus = types.MeterReadingOrderStatusHasorder
	}

	tx := db.Begin()
	if fee != nil {
		if err := tx.Create(fee).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	err = db.Table((*types.MeterReadingTransfer)(nil).TableName()).
		Where("id = ?", r.ID).
		Updates(r).
		Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// generateNormalMeterFee 生成日常设备账单 输入房间id和设备类型
func generateNormalMeterFee(db *gorm.DB, l *logrus.Entry, item *types.MeterReadingTransfer, now time.Time) (*types.Order, error) {
	log := l.WithField("room_id", item.RoomID)
	lastr, err := GetLastReading(db, item.RoomID, item.Type)
	if err != nil {
		log.Errorf("generateNormalMeterFee: 无法获取上次读数, %s", err)
		return nil, err
	}
	item.ThisReading = mathx.Round(item.ThisReading, 2)
	lastr.ThisReading = mathx.Round(lastr.ThisReading, 2)

	if item.ThisReading < lastr.ThisReading {
		return nil, gerr.BadRequest("%.2f小于上次读数%.2f", item.ThisReading, lastr.ThisReading)
	} else if item.ThisReading == lastr.ThisReading {
		// 零元账单
		return nil, nil
	}

	up, err := devices.GetMeterUnitPrice(db, item.RoomID, item.Type)
	if err != nil {
		log.Errorf("got meter price failed, %s", err)
		return nil, err
	}
	res, err := resident.Get(db, *item.ResidentID)
	if err != nil {
		log.Errorf("got resident failed, %s", err)
		return nil, err
	}
	odType, err := item.Type.ParseOrderType()
	if err != nil {
		log.Errorf("parse order type failed, %s", err)
		return nil, err
	}

	var (
		my, mm    = getNormalFeeCircleYearMonth(now)
		py, pm    = getNormalPayCircleYearMonth(now)
		mergeTime = fmt.Sprintf("%02d-%02d", my, mm)
		payCircle = fmt.Sprintf("%02d-%02d", py, pm)
		money     = colcMeterFee(item.ThisReading, lastr.ThisReading, item.Weight, up)
	)
	if money.IsZero() {
		// 不应该的零元账单
		log.Errorf("generateNormalMeterFee: wrong zero order.")
		return nil, nil
	}
	od := &types.Order{
		SimpleFee: types.SimpleFee{
			ResidentID:   res.ID,
			Type:         odType,
			Money:        money,
			BeginTime:    lastr.ThisTime,
			EndTime:      item.ThisTime,
			BeginReading: lastr.ThisReading,
			EndReading:   item.ThisReading,
			PayCircle:    payCircle,
		},
		CustomerID: res.CustomerID,
		StoreID:    item.StoreID,
		Number:     orderv3.GenerateOrderNumber(),
		ShouldPay:  money,
		HasPaid:    0,
		RoomID:     item.RoomID,
		Year:       my,
		Month:      mm,
		MergeTime:  mergeTime,
		GroupType:  types.OrderGroupTypeTemp,
		AddWay:     types.OrderAddWayRenewals,
		Status:     types.OrderStatusGenerate,
	}
	return od, nil
}

// generateChangeMeterFee 生成换表账单 输入房间id和设备类型
func generateChangeMeterFee(db *gorm.DB, item *ChangeMeterOption, now time.Time) (*types.Order, error) {
	r, err := room.Get(db, item.RoomID)
	if err != nil {
		return nil, err
	}
	if r.ResidentID == 0 {
		return nil, nil
	}

	lastr, err := GetLastReading(db, item.RoomID, item.MeterType)
	if err != nil {
		return nil, err
	}
	odType, err := item.MeterType.ParseOrderType()
	if err != nil {
		return nil, err
	}

	res, err := resident.Get(db, r.ResidentID)
	if err != nil {
		return nil, err
	}
	up, err := devices.GetMeterUnitPrice(db, item.RoomID, item.MeterType)
	if err != nil {
		return nil, err
	}

	var (
		year, month          = now.Year(), now.Month()
		mergeTime, payCircle = now.Format(global.MergeTimeFormat), now.Format(global.MergeTimeFormat)
		money                = colcMeterFee(item.OldReading.Reading, lastr.ThisReading, item.OldReading.Weight, up)
	)

	od := &types.Order{
		SimpleFee: types.SimpleFee{
			ResidentID:   res.ID,
			Type:         odType,
			Money:        money,
			BeginTime:    lastr.ThisTime,
			EndTime:      item.OldReading.ThisTime,
			BeginReading: lastr.ThisReading,
			EndReading:   item.OldReading.Reading,
			PayCircle:    payCircle,
		},
		CustomerID: res.CustomerID,
		StoreID:    r.StoreID,
		Number:     orderv3.GenerateOrderNumber(),
		ShouldPay:  money,
		HasPaid:    0,
		RoomID:     item.RoomID,
		Year:       year,
		Month:      month,
		MergeTime:  mergeTime,
		GroupType:  types.OrderGroupTypeTemp,
		AddWay:     types.OrderAddWayAdd,
		Status:     types.OrderStatusPending,
	}

	return od, nil
}
