package main

import (
	"fmt"
)

func main() {
	// 匿名结构体
	var config struct {
		APIKEY string 
		USERNAME string 
		AGE int
	}

	config.APIKEY = "123admin"
	config.USERNAME = "clarence"
	config.AGE = 20

	temp := struct {
		code uint
		message string 
	} {
		200,
		"success",
	}

	fmt.Printf("config=%v.\ntemp=%v.", config, temp)
// 	config={123admin clarence 20}.
// 	temp={200 success}.
}

func CheckReading(ctx *gin.Context) {
	in := &CheckReadingOption{}
	if err != ctx.ShouldBindQuery(in); err != nil {
		fx.BadRequest(ctx, err.Error())
		return 
	}
	if err := CheckReading(db, in.RoomID, in.Reading); err != nil {
		fx.BadRequest(ctx, err.Error())
		return 
	}
	l := struct {
		message string
		Reading meter_reading.Reading
		RoomID uint 
	} {
		"success",
		in.Reading,
		in.RoomID,
	}
	fx.Succ(ctx, l)
}
