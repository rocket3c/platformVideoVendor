package network

import (
	"fmt"
)

func Testinterface() {	
	 //test time 
	// time := RuokGetTimestamps()
	// fmt.Println(time)   //test time 
	
	Date, Day := Rkt_Date_time()
	if Date == "" || Day == "" {
		fmt.Println("time string error!")
	}
	if Debug {
		fmt.Println("************test createmember:************")
	}
	createmember := &CreateMember{}

	createmember.Website = Website
	createmember.Username = "testname"
	createmember.Uppername = "dyihao"
	createmember.Password = "password123"

	_,err := Ruok_CreateMember(createmember)
	if err != nil {
		fmt.Println(err)
	}

	if Debug {
		fmt.Println("\n************test login:************")
	}
	Login := &RKLogin{} 

	Login.Website = Website
	Login.Username = "testname"
	 Login.Uppername = "dyihao"
	 Login.Lang ="zh-cn"
	
     err1 := Ruok_Login(Login)
	if err1 != nil {
		fmt.Println(err1)
	}
	
	if Debug {
		fmt.Println("\n************test login2:************")
	}
	Login2 := &RKLogin2{} 

	Login2.Website = Website
	Login2.Username = "testname"
	 Login2.Uppername = "dyihao"
	 Login2.Lang ="zh-cn"

	_,err2 := Ruok_Login2(Login2)
	if err2 != nil {
		fmt.Println(err2)
	}
	if Debug {
		fmt.Println("\n********test logout:************")
	}
	Logout := &RKLogout{} 

	Logout.Website = Website
	Logout.Username = "testname"

	_,err3 := Ruok_Logout(Logout)
	if err3 != nil {
		fmt.Println(err3)
	}

	if Debug {
		fmt.Println("\n********test CheckUsrbanlance:************")
	}
	CheckUsr := &RKCheckUsrBalance{} 

	CheckUsr.Website = Website
	CheckUsr.Username = "testname"
	CheckUsr.Uppername = "dyihao"
	
	 var CheckUsrBalance = &CheckUsrResponse{} //[]byte
	 CheckUsrBalance,err4 := Ruok_CheckUsrBalance(CheckUsr)
	if err4 != nil {
		fmt.Println(err4)
	}
	if Debug {
		fmt.Println(CheckUsrBalance)
	}

	if Debug {
		fmt.Println("\n********test transfer:************")
	}
	Transfer := &RKTransfer{} 

	Transfer.Website = Website
	Transfer.Username = "testname"
	Transfer.Uppername = "dyihao"
	Transfer.Remitno = "9223372036854775806"
	Transfer.Action = "OUT"
	Transfer.Remit = 1

	 var rkTransfer = &TransferResponse{} //[]byte
	 rkTransfer,err5 := Ruok_Transfer(Transfer)
	if err5!= nil {
		fmt.Println(err5)
	}
	if Debug {
		fmt.Println(rkTransfer)
	}
	
	//TODO   test
	if Debug {
		fmt.Println("\n********test CheckTransfer:************")
	}
	CheckTransfer := &RKCheckTransfer{} 
	CheckTransfer.Website = Website
	CheckTransfer.Transid = "9223372036854775806"

	 var rkCheckTransfer = &CheckTransferResponse{} //[]byte
	 rkCheckTransfer,err6 := Ruok_CheckTransfer(CheckTransfer)
	if err6!= nil {
		fmt.Println(err6)
	}
	if Debug {
		fmt.Println(rkCheckTransfer)
	}

	if Debug {
		fmt.Println("\n********test TransferRecord:************")
	}
	TransferRecord := &RKTransferRecord{} 
	TransferRecord.Website = Website
	TransferRecord.Username = "testname"
	TransferRecord.Uppername = "dyihao"
	TransferRecord.Transid = "9223372036854775806"
	TransferRecord.Transtype = "OUT"
	TransferRecord.Date_start = Date
	TransferRecord.Date_end = Date
	TransferRecord.Start_hhmmss = Day
	TransferRecord.End_hhmmss = Day
	TransferRecord.Page = 1
	TransferRecord.Pagelimit = 100

	 var rkTransferRecord = &TransferRecordResponse{} //[]byte
	 rkTransferRecord,err7 := Ruok_TransferRecord(TransferRecord)
	if err7 != nil {
		fmt.Println(err7)
	}
	if Debug {
		fmt.Println(rkTransferRecord)
	}
	if Debug {
		fmt.Println("\n********test PlayGame:************")
	}
	PlayGame := &RKPlayGame{} 

	PlayGame.Website = Website
	PlayGame.Username = "testname"
	PlayGame.Gamekind = 3
	PlayGame.Gametype = 3001
	PlayGame.Gamecode = 1
	PlayGame.Lang = "zh-cn"

	 //var rkPlayGame = &network.PlayGameResponse{} //[]byte
	 err8 := Ruok_PlayGame(PlayGame)
	if err8 != nil {
		fmt.Println(err8)
	}

	if Debug {
		fmt.Println("\n********test BetRecord:************")
	}
	BetRecord := &RKBetRecord{} 

	BetRecord.Website = Website
	BetRecord.Username = "testname"
	BetRecord.Uppername = "dyihao"
	BetRecord.Rounddate = Date
	BetRecord.Starttime = Day
	BetRecord.Endtime = Day
	BetRecord.Gamekind = 3
	BetRecord.Subgamekind = 3
	BetRecord.Gametype = 1
	BetRecord.Page = 1
	BetRecord.Pagelimit =100

	var rkBetRecord = &BetRecordResponse{} //[]byte
	 rkBetRecord,err9 := Ruok_BetRecord(BetRecord)
	if err9 != nil {
		fmt.Println(err9)
	}
	if Debug {
		fmt.Println(rkBetRecord)
	}

	if Debug {
		fmt.Println("\n********test BetRecordByModifiedDate3:************")
	}
	BetRecordByModifiedDate3 := &RKBetRecordByModifiedDate3{} 

	BetRecordByModifiedDate3.Website = Website
	BetRecordByModifiedDate3.Start_date = Date
	BetRecordByModifiedDate3.End_date = Date
	BetRecordByModifiedDate3.Starttime =  Day
	BetRecordByModifiedDate3.Endtime = Day
	BetRecordByModifiedDate3.Gamekind = 3
	BetRecordByModifiedDate3.Subgamekind = 3
	BetRecordByModifiedDate3.Gametype =3
	BetRecordByModifiedDate3.Page = 1
	BetRecordByModifiedDate3.Pagelimit =100

	var rkBetRecordByModifiedDate3 = &BetRecordByModifiedDate3Response{} //[]byte
	rkBetRecordByModifiedDate3,err10 := Ruok_BetRecordByModifiedDate3(BetRecordByModifiedDate3)
	if err10 != nil {
		fmt.Println(err10)
	}
	if Debug {
		fmt.Println(rkBetRecordByModifiedDate3)
	}
	
	if Debug {
	fmt.Println("\n********test GetJPHistory:************")
	}
	GetJPHistory := &RKGetJPHistory{} 

	GetJPHistory.Website = Website
	GetJPHistory.Start_date = Date
	GetJPHistory.End_date = Date
	GetJPHistory.Starttime = Day
	GetJPHistory.Endtime = Day
	GetJPHistory.Jptype = 1
	GetJPHistory.Page = 1
	GetJPHistory.Pagelimit =100

	var rkGetJPHistory = &GetJPHistoryResponse{} //[]byte
	rkGetJPHistory,err11 := Ruok_GetJPHistory(GetJPHistory)
	if err11 != nil {
		fmt.Println(err11)
	}
	if Debug {
		fmt.Println(rkGetJPHistory)
	}
	if Debug {
	fmt.Println("\n********test WagersRecordBy38:************")
	}
	WagersRecordBy38 := &RKWagersRecordBy38{} 

	WagersRecordBy38.Website = Website
	WagersRecordBy38.Action = "BetTime"
	WagersRecordBy38.Username = "testname"
	WagersRecordBy38.Uppername = "dyihao"
	WagersRecordBy38.Start_date = Date
	WagersRecordBy38.End_date = Date
	WagersRecordBy38.Starttime = Day
	WagersRecordBy38.Endtime = Day
	WagersRecordBy38.Gametype = 3
	WagersRecordBy38.Page = 1
	WagersRecordBy38.Pagelimit =100

	var rkWagersRecordBy38 = &WagersRecordBy38Response{} //[]byte
	rkWagersRecordBy38,err12 := Ruok_WagersRecordBy38(WagersRecordBy38)
	if err12 != nil {
		fmt.Println(err12)
	}
	if Debug {
		fmt.Println(rkWagersRecordBy38)
	}
	if Debug {
		fmt.Println("\n********test WagersRecordBy30:************")
	}
	WagersRecordBy30 := &RKWagersRecordBy30{} 

	WagersRecordBy30.Website = Website
	WagersRecordBy30.Action = "BetTime"
	WagersRecordBy30.Uppername = "dyihao"
	WagersRecordBy30.Start_date = Date
	WagersRecordBy30.End_date = Date
	WagersRecordBy30.Starttime = Day
	WagersRecordBy30.Endtime = Day
	WagersRecordBy30.Gametype = 3
	WagersRecordBy30.Page = 1
	WagersRecordBy30.Pagelimit =100

	var rkWagersRecordBy30 = &WagersRecordBy30Response{} //[]byte
	rkWagersRecordBy30,err13 := Ruok_WagersRecordBy30(WagersRecordBy30)
	if err13 != nil {
		fmt.Println(err13)
	}
	if Debug {
		fmt.Println(rkWagersRecordBy30)
	}
	if Debug {
		fmt.Println("\n********test GetFishEventHistory:************")
	}
	GetFishEventHistory := &RKGetFishEventHistory{} 

	GetFishEventHistory.Website = Website
	GetFishEventHistory.Rounddate = Date
	GetFishEventHistory.Starttime = Day
	GetFishEventHistory.Endtime = Day
	GetFishEventHistory.Page = 1
	GetFishEventHistory.Pagelimit =100

	var rkGetFishEventHistory = &GetFishEventResponse{} //[]byte
	rkGetFishEventHistory,err14 := Ruok_GetFishEventHistory(GetFishEventHistory)
	if err14 != nil {
		fmt.Println(err14)
	}
	if Debug {
		fmt.Println(rkGetFishEventHistory)
	}
	if Debug {
		fmt.Println("\n********test GetSKEventHistory:************")
	}
	GetSKEventHistory := &RKGetSKEventHistory{} 

	GetSKEventHistory.Website = Website
	GetSKEventHistory.Rounddate = Date
	GetSKEventHistory.Starttime = Day
	GetSKEventHistory.Endtime = Day
	GetSKEventHistory.Page = 1
	GetSKEventHistory.Pagelimit =100

	var rkGetSKEventHistory = &GetSKEventHistoryResponse{} //[]byte
	rkGetSKEventHistory,err15 := Ruok_GetSKEventHistory(GetSKEventHistory)
	if err15 != nil {
		fmt.Println(err15)
	}
	if Debug {
		fmt.Println(rkGetSKEventHistory)
	}
	if Debug {
		fmt.Println("\n********test GetSportEventHistory:************")
	}
	GetSportEventHistory := &RKGetSportEventHistory{} 

	GetSportEventHistory.Website = Website
	GetSportEventHistory.Rounddate = Date
	GetSportEventHistory.Starttime = Day
	GetSportEventHistory.Endtime = Day
	GetSportEventHistory.Page = 1
	GetSportEventHistory.Pagelimit =100

	var rkGetSportEventHistory = &GetSportEventHistoryResponse{} //[]byte
	rkGetSportEventHistory,err16 := Ruok_GetSportEventHistory(GetSportEventHistory)
	if err16 != nil {
		fmt.Println(err16)
	}
	if Debug {
		fmt.Println(rkGetSportEventHistory)
	}	
}