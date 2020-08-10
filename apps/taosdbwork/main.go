package main

import (
	"database/sql"
	"fmt"
	"time"
)

// MEHistory 10 分钟一次
type MEHistory struct {
	StartDate    *time.Time
	EndDate      *time.Time
	StartValueV1 float64
	StartValueV2 float64
	StartValueV3 float64
	StartValueV4 float64
	StartValueV5 float64
	EndValueV1   float64
	EndValueV2   float64
	EndValueV3   float64
	EndValueV4   float64
	EndValueV5   float64
	UsedValueV1  float64
	UsedValueV2  float64
	UsedValueV3  float64
	UsedValueV4  float64
	UsedValueV5  float64
	BeforeMoney  float64
	AfterMoney   float64
	LastCrash    float64
	V1PrePrice   float64
	V2PrePrice   float64
	V3PrePrice   float64
	V4PrePrice   float64
	V5PrePrice   float64
	UsedCrash    float64 //
}

// MEDevice 电表
type MEDevice struct {
	Number        string  //
	PublicDevice  bool    //
	AlarmValue    float64 //
	RealPower     float64 //  实时功率
	ModbusAddress uint8   //
	ModelName     string  // 型号
	ShopNum       string  // 商铺号
	CT            int     // 1
	DefaultValue  float64 //
	StartValueV1  float64 //
	V1PrePrice    float64 //
	StartValueV2  float64 //
	V2PrePrice    float64 //
	StartValueV3  float64 //
	V3PrePrice    float64 //
	StartValueV4  float64 //
	V4PrePrice    float64 //
	StartValueV5  float64 //
	V5PrePrice    float64 //
	OnOff         int     // 合关
	AlarmStatus   bool    // 状态
	W1            float64 // 告警1
	W2            float64 // 告警2
	DefaultMoney  float64 //
	Remark        string  // 备注
	MUser         MUser   //
	MUserID       uint    // 那个租户的
	OnLine        int     //  2 离线 1 在线
	StartValue    float64 //

	V1 float64 //
	V2 float64 //
	V3 float64 //
	V4 float64 //
	V5 float64 //
	UA float64 //
	UB float64 //
	UC float64 //
	IA float64 //
	IB float64 //
	IC float64 //
	P  float64 //

	StartDate *time.Time
}

func CreateSuperTable(db *sql.DB, modelType int) {
	if modelType == "DTSD" {
		superTable := "CREATE TABLE DTSD (ts timestamp, degree double) TAGS(location binary(20), type int);"
		st := time.Now().Nanosecond()
		// create table
		res, err := db.Exec("create table " + demot + " (ts timestamp, id int, name binary(8), len tinyint, flag bool, notes binary(8), fv float, dv double)")
		checkErr(err, "create table db.Exec")

		affectd, err := res.RowsAffected()
		checkErr(err, "create table res.RowsAffected")

		et := time.Now().Nanosecond()
		fmt.Printf("create table result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1E9)
	}
}
