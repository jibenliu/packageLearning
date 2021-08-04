package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"migrationVat/migrate/model"
	"strconv"
)

func main() {
	dsn := "invoice:invoice@tcp(10.5.10.117:3306)/dev_old?charset=utf8mb4&parseTime=True&loc=Local"
	oldDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("数据库链接失败 %v", err)
	}
	dsn = "invoice:invoice@tcp(10.5.10.117:3306)/mysoft_vat?charset=utf8mb4&parseTime=True&loc=Local"
	newDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("数据库链接失败 %v", err)
	}
	//newDB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//newDB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//migrateTaxEntity(oldDB, newDB)
	migrateTaxEntityMapping(oldDB, newDB)
}

func migrateZone(oldDB *gorm.DB, newDB *gorm.DB) {
	var oldZones []model.Taxsubject
	oldDB.Where("OrgType = ?", 1).Find(&oldZones)
	for _, oldZone := range oldZones {
		newZone := model.VatZone{}
		err := newDB.Where("name = ?", oldZone.Name).First(&newZone).Error
		if err == nil {
			continue
		}
		oldZoneParent := model.Taxsubject{}
		oldDB.Where("Id = ?", oldZone.ParentId).First(&oldZoneParent)
		newZoneParent := model.VatZone{}
		newDB.Where("name = ?", oldZoneParent.Name).First(&newZoneParent)
		newZone.Pid = int(newZoneParent.ID)
		newZone.Name = oldZone.Name
		if newZoneParent.ID != 0 {
			if newZoneParent.Pkey != "" {
				newZone.Pkey = newZoneParent.Pkey + "-" + strconv.Itoa(int(newZoneParent.ID))
			} else {
				newZone.Pkey = "0-" + strconv.Itoa(int(newZoneParent.ID))
			}
		} else {
			newZone.Pkey = "0"
		}

		newDB.Create(&newZone)
	}
}

func migrateTaxEntity(oldDB *gorm.DB, newDB *gorm.DB) {
	var oldTaxEntities []model.Taxsubject
	oldDB.Where("OrgType = ?", 2).Find(&oldTaxEntities)
	for _, oldTaxEntity := range oldTaxEntities {
		newTaxEntity := model.VatTaxEntity{}
		err := newDB.Where("tax_code = ?", oldTaxEntity.NSRSBH).First(&newTaxEntity).Error
		if err == nil {
			continue
		}
		oldTaxEntityZone := model.Taxsubject{}
		oldDB.Where("Id = ?", oldTaxEntity.ParentId).First(&oldTaxEntityZone)
		newTaxEntityZone := model.VatZone{}
		newDB.Where("name = ?", oldTaxEntityZone.Name).First(&newTaxEntityZone)
		newTaxEntity.Name = oldTaxEntity.Name
		newTaxEntity.ZoneId = int(newTaxEntityZone.ID)
		newTaxEntity.Address = oldTaxEntity.Address
		newTaxEntity.Mobile = oldTaxEntity.PhoneNum
		newTaxEntity.TaxCode = oldTaxEntity.NSRSBH
		newTaxEntity.HasAuth = 0
		newDB.Create(&newTaxEntity)

		var oldTaxEntityBanks []model.TaxSubjectBank
		var newTaxEntityBanks []model.VatTaxEntityBank
		err = oldDB.Where("TaxSubjectId = ?", oldTaxEntity.Id).Find(&oldTaxEntityBanks).Error
		if err == nil {
			for _, oldBank := range oldTaxEntityBanks {
				newTaxEntityBank := model.VatTaxEntityBank{}
				newTaxEntityBank.EntityId = newTaxEntity.Id
				newTaxEntityBank.BankName = oldBank.BankName
				newTaxEntityBank.BankAccount = oldBank.BankNum
				newTaxEntityBank.IsDefault = oldBank.IsDefault
				newTaxEntityBanks = append(newTaxEntityBanks, newTaxEntityBank)
			}
			newDB.Create(&newTaxEntityBanks)
		} else if err != gorm.ErrRecordNotFound { //防止数据异常
			newTaxEntityBank := model.VatTaxEntityBank{}
			newTaxEntityBank.EntityId = newTaxEntity.Id
			newTaxEntityBank.BankName = oldTaxEntity.Bank
			newTaxEntityBank.BankAccount = oldTaxEntity.BankNum
			newTaxEntityBank.IsDefault = 1
			newDB.Create(&newTaxEntityBank)
		}

		var newTaxEntityAddressInfo model.VatTaxEntityAddressInfo
		newTaxEntityAddressInfo.EntityId = newTaxEntity.Id
		newTaxEntityAddressInfo.Address = newTaxEntity.Address
		newTaxEntityAddressInfo.Mobile = newTaxEntity.Mobile
		newTaxEntityAddressInfo.IsDefault = 1
		newDB.Create(&newTaxEntityAddressInfo)
	}
}

func migrateTaxEntityMapping(oldDB *gorm.DB, newDB *gorm.DB) {
	type Result struct {
		TaxSubjectId string
		NSRSBH       string
	}
	var results []Result
	subQuery := oldDB.Model(&model.Project{}).Distinct("TaxSubjectId").Group("TaxSubjectId")
	oldDB.Model(&model.Taxsubject{}).Select("Id", "NSRSBH").Where("Id IN (?)", subQuery).Find(&results)
	var taxCodeArr []string
	for _, result := range results {
		taxCodeArr = append(taxCodeArr, result.NSRSBH)
	}
	type NewResult struct {
		Id      int
		TaxCode string
	}
	var newResults []NewResult
	newDB.Model(&model.VatTaxEntity{}).Select("id", "tax_code").Where("tax_code in (?)", taxCodeArr).Find(&newResults)
	var taxSubjectId2IdArr []map[string]int
	for _, newResult := range newResults {
		for _, result := range results {
			if result.NSRSBH == newResult.TaxCode {
				var temp = map[string]int{}
				temp[result.TaxSubjectId] = newResult.Id
				taxSubjectId2IdArr = append(taxSubjectId2IdArr, temp)
			}
		}
	}

}
