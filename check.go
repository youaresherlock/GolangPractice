	// 校验当前员工是否有权限导入数据
	ids := []uint{}
	for i, row := range sheet.Rows {
		if i < 3 {
			continue
		}
		id, err := strconv.Atoi(row.Cells[0].String())
		if err != nil || id <= 0 {
			return nil, gerr.BadRequest("错误原因: 文件第%d行数据有误，无法获取ID. error: %s", i+1, err.Error())
		}

		ids = append(ids, uint(id))
	}
	storesID := []uint{}
	db.Table((*types.MeterReadingTransfer)(nil).TableName()).
		Where("id in (?)", ids).
		Group("store_id").
		Pluck("store_id", &storesID)
	storesWithPermission, err := employee.GetEmployeeStoreIDs(db, eid, ids...)
	if err != nil {
		return nil, err
	}
	if len(storesWithPermission) != len(storesID) {
		return nil, errors.New("当前员工无操作此门店权限")
	}