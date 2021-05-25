package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mssql_exam/models"
	"strings"
	"time"
)

func main() {
	var server = "10.5.10.165"
	var user = "sa"
	var password = "95938"
	var database = "dotnet_erp60_test"
	//var port = 1433
	var version = "sql2016"

	connString := fmt.Sprintf("sqlserver://%s:%s@%s/%s?database=%s", user, password, server, version, database)
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	defer sqlDB.Close()

	//projectId := "04180488-8929-EB11-B565-00155D0A1C4E"
	//roomsSql := fmt.Sprintf("SELECT *, 0 AS _NAV_ORDER_F_ FROM [dbo].[s_Building] WHERE [ProjGUID] LIKE '\\%%s\\%' ORDER BY _NAV_ORDER_F_ OFFSET 0 ROWS FETCH NEXT 1000 ROWS ONLY", projectId)
	//产生查询语句的Statement
	//rows := db.Exec(`SELECT * FROM [s_Room] WHERE [BldGUID] = N'C12594EB-8929-EB11-B565-00155D0A1C4E' AND [RoomGUID] <> N'%s'`, "3EACBB14-8C29-EB11-B565-00155D0A1C4E")
	//defer rows.Close()

	_ = db.Callback().Create().Remove("gorm:created_at")
	_ = db.Callback().Update().Remove("gorm:updated_at")
	_ = db.Callback().Delete().Remove("gorm:deleted_at")
	_ = db.Callback().Query().Remove("gorm:deleted_at")
	rows, err := db.Model(&models.Room{}).
		Not("RoomGUID", "3EACBB14-8C29-EB11-B565-00155D0A1C4E").
		Where("BldGUID = ?", "C12594EB-8929-EB11-B565-00155D0A1C4E").
		Select("*, CAST(RoomGUID AS varchar(36)) AS RoomGUID").Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var trade models.Trade
	err = db.Debug().Model(&models.Trade{}).
		Where("RoomGUID = ?", "3EACBB14-8C29-EB11-B565-00155D0A1C4E").
		Select("*, " +
			" CAST(TradeGUID AS varchar(36)) AS TradeGUID," +
			" CAST(ModifiedGUID AS varchar(36)) AS ModifiedGUID," +
			" CAST(CreatedGUID AS varchar(36)) AS CreatedGUID," +
			" CAST(BUGUID AS varchar(36)) AS BUGUID," +
			" CAST(ZcContractGUID AS varchar(36)) AS ZcContractGUID," +
			" CAST(XDOrderGUID AS varchar(36)) AS XDOrderGUID," +
			" CAST(RGOrderGUID AS varchar(36)) AS RGOrderGUID," +
			" CAST(ContractGUID AS varchar(36)) AS ContractGUID," +
			" CAST(ProjGUID AS varchar(36)) AS ProjGUID" +
			"").Scan(&trade).Error
	if err != nil {
		fmt.Sprintf("error is %v", err)
		return
	}

	var tradeReDundancy models.TradeReDundancy
	err = db.Debug().Model(&models.TradeReDundancy{}).
		Where("TradeGUID = ?", trade.TradeGUID).
		Select("*, " +
			"CAST(TradeGUID AS varchar(36)) AS TradeGUID," +
			" CAST(ModifiedGUID AS varchar(36)) AS ModifiedGUID," +
			" CAST(CreatedGUID AS varchar(36)) AS CreatedGUID," +
			" CAST(TradeRedundancyGUID AS varchar(36)) AS TradeRedundancyGUID" +
			"").Scan(&tradeReDundancy).Error
	if err != nil {
		fmt.Sprintf("error is %v", err)
		return
	}

	var order models.Order
	err = db.Debug().Model(&models.Order{}).
		Where("TradeGUID = ?", trade.TradeGUID).
		Select("*, " +
			"CAST(TradeGUID AS varchar(36)) AS TradeGUID," +
			" CAST(ModifiedGUID AS varchar(36)) AS ModifiedGUID," +
			" CAST(CreatedGUID AS varchar(36)) AS CreatedGUID," +
			" CAST(BUGUID AS varchar(36)) AS BUGUID," +
			" CAST(JbrGUID AS varchar(36)) AS JbrGUID," +
			" CAST(ProjGUID AS varchar(36)) AS ProjGUID," +
			" CAST(OrderGUID AS varchar(36)) AS OrderGUID" +
			"").Scan(&order).Error
	if err != nil {
		fmt.Sprintf("error is %v", err)
		return
	}

	var getIns []models.GetIn
	err = db.Debug().Model(&models.GetIn{}).
		Where("SaleGUID = ?", trade.TradeGUID).
		Select("*, " +
			"CAST(SaleGUID AS varchar(36)) AS SaleGUID," +
			" CAST(ModifiedGUID AS varchar(36)) AS ModifiedGUID," +
			" CAST(ItemNameGUID AS varchar(36)) AS ItemNameGUID," +
			" CAST(ItemTypeGUID AS varchar(36)) AS ItemTypeGUID," +
			" CAST(ProjGUID AS varchar(36)) AS ProjGUID," +
			" CAST(VouchGUID AS varchar(36)) AS VouchGUID," +
			" CAST(CreatedGUID AS varchar(36)) AS CreatedGUID," +
			" CAST(GetinGUID AS varchar(36)) AS GetinGUID" +
			"").Scan(&getIns).Error
	if err != nil {
		fmt.Sprintf("error is %v", err)
		return
	}

	var vouchers []models.Voucher
	err = db.Debug().Model(&models.Voucher{}).
		Where("SaleGUID = ?", trade.TradeGUID).
		Select("*, " +
			"CAST(SaleGUID AS varchar(36)) AS SaleGUID," +
			" CAST(ModifiedGUID AS varchar(36)) AS ModifiedGUID," +
			" CAST(CreatedGUID AS varchar(36)) AS CreatedGUID," +
			" CAST(VouchGUID AS varchar(36)) AS VouchGUID," +
			" CAST(ProjGUID AS varchar(36)) AS ProjGUID," +
			" CAST(RoomGUID AS varchar(36)) AS RoomGUID," +
			" CAST(BUGUID AS varchar(36)) AS BUGUID" +
			"").Scan(&vouchers).Error
	if err != nil {
		fmt.Sprintf("error is %v", err)
		return
	}

	var fees []models.Fee
	err = db.Debug().Model(&models.Fee{}).
		Where("TradeGUID = ?", trade.TradeGUID).
		Select("*, " +
			"CAST(ItemNameGuid AS varchar(36)) AS ItemNameGuid," +
			" CAST(ItemTypeGuid AS varchar(36)) AS ItemTypeGuid," +
			" CAST(ProjGUID AS varchar(36)) AS ProjGUID," +
			" CAST(BUGUID AS varchar(36)) AS BUGUID" +
			"").Scan(&fees).Error
	if err != nil {
		fmt.Sprintf("error is %v", err)
		return
	}

	for rows.Next() {
		room := &models.Room{}
		err = db.ScanRows(rows, room)
		if err != nil {
			panic(err)
		}
		var newTrade models.Trade
		newTrade = trade
		newTrade.TradeGUID = strings.ToUpper(uuid.NewV4().String())
		newTrade.CreatedGUID = strings.ToUpper(uuid.NewV4().String())
		newTrade.ModifiedGUID = strings.ToUpper(uuid.NewV4().String())
		newTrade.RoomGUID = room.RoomGUID

		var newTradeReDundancy models.TradeReDundancy
		newTradeReDundancy = tradeReDundancy
		newTradeReDundancy.TradeRedundancyGUID = strings.ToUpper(uuid.NewV4().String())
		newTradeReDundancy.TradeGUID = newTrade.TradeGUID
		newTradeReDundancy.CreatedGUID = strings.ToUpper(uuid.NewV4().String())
		newTradeReDundancy.ModifiedGUID = strings.ToUpper(uuid.NewV4().String())

		var newOrder models.Order
		newOrder = order
		newOrder.OrderGUID = strings.ToUpper(uuid.NewV4().String())
		newOrder.TradeGUID = newTrade.TradeGUID
		newOrder.CreatedGUID = strings.ToUpper(uuid.NewV4().String())
		newOrder.ModifiedGUID = strings.ToUpper(uuid.NewV4().String())
		newOrder.RoomGUID = room.RoomGUID

		newTrade.RGOrderGUID = newOrder.OrderGUID
		newTrade.ZcOrderGUID = newOrder.OrderGUID

		voucherGUIDArr := map[int]string{}
		for index, voucher := range vouchers {
			var newVoucher models.Voucher
			newVoucher = voucher
			newVoucher.CreatedGUID = strings.ToUpper(uuid.NewV4().String())
			newVoucher.ModifiedGUID = strings.ToUpper(uuid.NewV4().String())
			newVoucher.VouchGUID = strings.ToUpper(uuid.NewV4().String())
			newVoucher.RoomGUID = room.RoomGUID
			newVoucher.SaleGUID = newTrade.TradeGUID
			voucherGUIDArr[index] = newVoucher.VouchGUID
			db.Create(&newVoucher)
		}

		for index, getIn := range getIns {
			var newGetIn models.GetIn
			newGetIn = getIn
			newGetIn.CreatedGUID = strings.ToUpper(uuid.NewV4().String())
			newGetIn.ModifiedGUID = strings.ToUpper(uuid.NewV4().String())
			newGetIn.GetinGUID = strings.ToUpper(uuid.NewV4().String())
			newGetIn.RoomGUID = room.RoomGUID
			newGetIn.SaleGUID = newTrade.TradeGUID
			newGetIn.VouchGUID = voucherGUIDArr[index]
			db.Create(&newGetIn)
		}

		for _, fee := range fees {
			var newFee models.Fee
			newFee = fee
			newFee.CreatedGUID = strings.ToUpper(uuid.NewV4().String())
			newFee.ModifiedGUID = strings.ToUpper(uuid.NewV4().String())
			newFee.FeeGUID = strings.ToUpper(uuid.NewV4().String())
			newFee.TradeGUID = newTrade.TradeGUID
			db.Create(&newFee)
		}

		db.Create(&newTrade)
		db.Create(&newOrder)
	}
}
