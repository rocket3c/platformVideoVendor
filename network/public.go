package network

import (
    "fmt"
	"io"
	"errors"
	"crypto/md5"
	"time"
    "encoding/json"
)

// var (
// 	Website 						= "avia"
// 	Uppername 						= "dyihao"
// 	Createmember_url 				= "http://linkapi.tcy789.com/app/WebService/JSON/display.php/CreateMember"
// 	Login_url 						= "http://888.cbapi01.com/app/WebService/JSON/display.php/Login"
// 	Login2_url 						= "http://888.cbapi01.com/app/WebService/JSON/display.php/Login2"
// 	Logout_url					    = "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/Logout"
// 	Checkuserbanace_url 			= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/CheckUsrBalance"
// 	Transfer_url					= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/Transfer"
// 	CheckTransfer_url				= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/CheckTransfer"
// 	TransferRecord_url				= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/TransferRecord"
// 	PlayGame_url					= "http://888.cbapi01.com/app/WebService/JSON/display.php/PlayGame"
// 	PlayGameByH5_url 				= "http://888.cbapi01.com/app/WebService/JSON/display.php/PlayGameByH5"
// 	BetRecord_url					= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/BetRecord"
// 	BetRecordByModifiedDate3_url 	= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/BetRecordByModifiedDate3"
// 	GetJPHistory_url 				= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/GetJPHistory"
// 	ForwardGameH5By5_url 			= "http://888.cbapi01.com/app/WebService/JSON/display.php/ForwardGameH5By5"
// 	WagersRecordBy38_url 			= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/WagersRecordBy38"
// 	WagersRecordBy30_url 			= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/WagersRecordBy30"
// 	ForwardGameH5By30_url			= "http://888.cbapi01.com/app/WebService/JSON/display.php/ForwardGameH5By30"
// 	ForwardGameH5By38_url			= "http://888.cbapi01.com/app/WebService/JSON/display.php/ForwardGameH5By38"
// 	GetFishEventHistory_url 		= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/GetFishEventHistory"
// 	FishEventUrl_url  			    = "http://888.cbapi01.com/app/WebService/JSON/display.php/FishEventUrl"
// 	GetSKEventHistory_url 			= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/GetSKEventHistory"
// 	SKEventUrl_url			        = "http://888.cbapi01.com/app/WebService/JSON/display.php/SKEventUrl"
// 	SportEventUrl_url			    = "http://888.cbapi01.com/app/WebService/JSON/display.php/SKEventUrl"
// 	GetSportEventHistory_url    	= "http://linkapi.cbapi01.com/app/WebService/JSON/display.php/GetSportEventHistory"
// )

 const debug = false     // debug 打印开关
 const Debug = true     // main测试程序 debug 打印开关


//通用参数
const (
	Website 						= "avia"
	Uppername 						= "dyihao"
	BaseA_url						= "http://linkapi.tcy789.com/"    //only createmember use
	BaseB_url					    = "http://linkapi.cbapi01.com/"
	BaseC_url						= "http://888.cbapi01.com/"
	route_url						= "app/WebService/JSON/display.php/"
)

//get 请求地址
const (
	Createmember_url				= BaseA_url + route_url + "CreateMember"
	Login_url 						= BaseC_url + route_url + "Login"
	Login2_url 						= BaseC_url + route_url + "Login2"
	Logout_url					    = BaseB_url + route_url + "Logout"
	Checkuserbanace_url 			= BaseB_url + route_url + "CheckUsrBalance"
	Transfer_url					= BaseB_url + route_url + "Transfer"
	CheckTransfer_url				= BaseB_url + route_url + "CheckTransfer"
	TransferRecord_url				= BaseB_url + route_url + "TransferRecord"
	PlayGame_url					= BaseC_url + route_url + "PlayGame"
	PlayGameByH5_url 				= BaseC_url + route_url + "PlayGameByH5"
	BetRecord_url					= BaseB_url + route_url + "BetRecord"
	BetRecordByModifiedDate3_url 	= BaseB_url + route_url + "BetRecordByModifiedDate3"
	GetJPHistory_url 				= BaseB_url + route_url + "GetJPHistory"
	ForwardGameH5By5_url 			= BaseC_url + route_url + "ForwardGameH5By5"
	WagersRecordBy38_url 			= BaseB_url + route_url + "WagersRecordBy38"
	WagersRecordBy30_url 			= BaseB_url + route_url + "WagersRecordBy30"
	ForwardGameH5By30_url			= BaseC_url + route_url + "ForwardGameH5By30"
	ForwardGameH5By38_url			= BaseC_url + route_url + "ForwardGameH5By38"
	GetFishEventHistory_url 		= BaseB_url + route_url + "GetFishEventHistory"
	FishEventUrl_url  			    = BaseC_url + route_url + "FishEventUrl"
	GetSKEventHistory_url 			= BaseB_url + route_url + "GetSKEventHistory"
	SKEventUrl_url			        = BaseC_url + route_url + "SKEventUrl"
	SportEventUrl_url			    = BaseC_url + route_url + "SKEventUrl"
	GetSportEventHistory_url    	= BaseB_url + route_url + "GetSportEventHistory"
)

// keyB 值
const (
	Createmember_keyB 				= "3QcgFxyY0"
	Login_keyB 						= "fV98jAu"
	Login2_keyB 					= "fV98jAu"
	Logout_keyB 					= "x2b7x"
	CheckUsrBalance_keyB 			= "7pxyd9c0a"
	Transfer_keyB 					= "10WyHdOdZ"
	CheckTransfer_keyB 				= "5Jr57Ya8c7"
	TransferRecord_keyB 			= "5Jr57Ya8c7"
	PlayGame_keyB 					= "05Rz1lv"
	PlayGameByH5_keyB 				= "05Rz1lv"
	BetRecord_keyB 					= "6kqBB1"
	BetRecordByModifiedDate3_keyB 	= "6kqBB1"
	GetJPHistory_keyB 				= "6kqBB1"
	ForwardGameH5By5_keyB 			= "05Rz1lv"
	WagersRecordBy38_keyB 			= "6kqBB1"
	WagersRecordBy30_keyB 			= "6kqBB1"
	ForwardGameH5By30_keyB 			= "05Rz1lv"
	ForwardGameH5By38_keyB 			= "05Rz1lv"
	GetFishEventHistory_keyB 		= "6kqBB1"
	FishEventUrl_keyB 				= "05Rz1lv"
	GetSKEventHistory_keyB 			= "6kqBB1"
	SKEventUrl_keyB				    = "05Rz1lv"
	SportEventUrl_keyB              = "05Rz1lv"
	GetSportEventHistory_keyB       = "6kqBB1"
)

/*******************  request 结构数据 ********************************/
type CreateMember struct {
	Website  			string     //avia
	Username 	 		string
	Uppername  			string
	Password   			string
	Key   				string   
}

type RKLogin struct {
	Website  			string     //avia
	Username 	 		string
	Uppername  			string
	Lang				string
	Page_site			string	
	Page_present		string
	Maintenance_page	string
	Key					string 
}

type RKLogin2 struct {
	Website  			string     //avia
	Username 	 		string
	Uppername  			string
	Lang				string
	Key					string 
}

type RKLogout struct {
	Website  			string     //avia
	Username 	 		string
	Key					string 
}

type RKCheckUsrBalance struct {
	Website  			string     //avia
	Username 	 		string
	Uppername  			string
    Page 				int
    Pagelimit           int
    Key					string
}
type RKTransfer struct {
	Website  			string  
	Username 	 		string
	Uppername  			string
    Remitno 			string    	//转帐序号(唯一值)
	Action              string  	//IN(转入额度) OUT(转出额度)
	Remit				int			//转帐额度(正整数)
    Key					string
}
type RKCheckTransfer struct {
	Website  			string    
	Transid 	 		string      //转帐序号，对应Transfer API中的remitno
    Key					string      //A= 无意义字串长度8码
									//B=MD5(website + KeyB + YYYYMMDD) 
									//C=无意义字串长度8码
}
type RKTransferRecord struct {
	Website  			string   
	Username 	 		string
	Uppername  			string
    Transid 			string   	//转帐序号，对应Transfer API中的remitno
	Transtype           string  	//IN转入;OUT转出
	Date_start			string  	//开始日期ex:2012/03/21、2012-03-21
    Date_end			string  	//结束日期ex:2012/03/21、2012-03-21
    Start_hhmmss        string   	//开始时间ex:00:00:00
	End_hhmmss			string      //结束时间ex:23:59:59
	Page      			int   	    //查询页数
	Pagelimit			int   		//每页数量
    Key					string    	//key=A+B+C(验证码组合方式)
}

type RKPlayGame struct {
	Website  			string    
	Username 	 		string
	Gamekind  			int
    Gametype 			int
	Gamecode           	int
	Lang				string
    Key					string  	//key=A+B+C(验证码组合方式)
}

type RKPlayGameH5 struct {
	Website  			string    
	Username 	 		string
	Gamekind  			string
    Gametype 			int
	Gamecode           	int
	Lang				string
    Key					string  	//key=A+B+C(验证码组合方式)
}

type RKBetRecord struct {
	Website  			string    
	Username 	 		string
	Uppername  			string
    Rounddate 			string      //日期ex:2012/03/21、2012-03-21 (gamekind=5，只可取7日内 的资料)
	Starttime			string		//开始时间ex:00:00:00(BB体育无效)
	Endtime				string		//结束时间ex:23:59:59(BB体育无效)
	Gamekind			int 		//游戏种类(1:BB体育、3:BB视讯、5:BB电子、12:BB彩票、 99:BB小费)
	Subgamekind			int
	Gametype			int
	Page      			int   		//查询页数
	Pagelimit			int   		//每页数量
    Key					string  	//key=A+B+C(验证码组合方式)
}
type RKBetRecordByModifiedDate3 struct {
	Website  			string     		//avia
	Start_date 	 		string		//开始日期ex:2012/03/21、2012-03-21
	End_date  			string		//开始日期ex:2012/03/21、2012-03-21
	Starttime			string		//开始时间ex:00:00:00
	Endtime				string		//结束时间ex:23:59:59
	Gamekind			int 			//游戏种类(1:BB体育、3:BB视讯、5:BB电子、12:BB彩票、 99:BB小费)
	Subgamekind			int
	Gametype			int
	Page      			int   			 //查询页数
	Pagelimit			int   			//每页数量
    Key					string  		 // key=A+B+C(验证码组合方式)
										//A= 无意义字串长度1码
										//B=MD5(website+ username + KeyB + YYYYMMDD) C=无意义字串长度8码
}	
type RKGetJPHistory struct {
	Website  			string     	
	Start_date			string		//开始日期ex:2012/03/21、2012-03-21
	End_date			string		//结束日期ex:2012/03/21、2012-03-21
	Starttime			string		//开始时间ex:00:00:00(BB体育无效)
	Endtime				string		//结束时间ex:23:59:59(BB体育无效)
	Jptype			    int 		//奖项:1=>Grand、2=>Major、3=>Minor、4=>Mini
	Page      			int   		//查询页数
	Pagelimit			int   		//每页数量
    Key					string  	//key=A+B+C(验证码组合方式)
}
type RKForwardGameH5By5 struct {
	Website  			string   
	Username 	 		string
	Uppername  			string
    Gametype 			int
    Key					string   	
}

type RKWagersRecordBy38 struct {
	Website  			string    
	Action				string 
	Username 	 		string
	Uppername  			string
	Start_date			string
	End_date			string
	Starttime			string
	Endtime				string
	Gametype			int
    Page 				int
    Pagelimit           int
    Key					string   	
}
type RKWagersRecordBy30 struct {
	Website  			string   
	Action				string 
	Uppername  			string
	Start_date			string
	End_date			string
	Starttime			string
	Endtime				string
	Gametype			int
    Page 				int
    Pagelimit           int
    Key					string   
}
type RKForwardGameH5By30 struct {
	Website  			string    
	Username 	 		string
	Uppername  			string
    Gametype 			int
	Key					string
}
				
type RKForwardGameH5By38 struct {
	Website  			string   
	Username 	 		string
	Uppername  			string
    Gametype 			int
	Key					string
}						
type RKGetFishEventHistory struct {
	Website  			string     
	Rounddate 	 		string
	Starttime  			string
	Endtime				string
    Page 				int
    Pagelimit           int
	Key					string
}						
						
type RKFishEventUrl struct {
	Website  			string  
	Username 	 		string
	Lang  				string
	Key					string
}		
	 				
type RKGetSKEventHistory struct {
	Website  			string    
	Rounddate			string
	Starttime			string
	Endtime				string
    Page 				int
    Pagelimit           int
	Key					string
}		
	 				
type RKSKEventUrl struct {
	Website  			string   
	Username 	 		string
	Lang  				string
	Key					string
}

type RKSportEventUrl struct {
	Website  			string    
	Username 	 		string
	Lang  				string
	Key					string
}

type RKGetSportEventHistory struct {
	Website  			string   
	Rounddate			string
	Starttime			string
	Endtime				string
    Page 				int
    Pagelimit           int
	Key					string
}		
	 

/*******************  response 结构数据 **************************/
type DateResponse struct {
	Code 		string `json:"Code"`
	Message 	string `json:"Message"`
}

//createmember
type CreateResponse struct {
	Result     bool          `json:"result"`
	Date       DateResponse  `json:"data"`
}
//login
type LoginResponse struct {
	Result     bool          `json:"result"`
	Date       DateResponse  `json:"data"`
}
//login2
type Login2DateResponse struct {
	Code 		int    `json:"Code"`
	Message 	string `json:"Message"`
}
type Login2Response struct {
	Result     bool         	   `json:"result"`
	Date       Login2DateResponse  `json:"data"`
}
//logout
type LogoutResponse struct {
	Result     bool          `json:"result"`
	Date       DateResponse  `json:"data"`
}
//checkusrbanlance 
type CheckUsrData struct {
	LoginName    string  	`json:"LoginName"`
	Currency     string 	`json:"Currency"`
	Balance	     int		`json:"Balance"`
	TotalBalance int 	`json:"TotalBalance"`
}
type CheckUsrPagination struct{
	Page    		int 	`json:"Page"`
	PageLimit     	int 	`json:"PageLimit"`
	TotalNumber	    int	    `json:"TotalNumber"`
	TotalPage 		int 	`json:"TotalPage"`
}
type CheckUsrResponse struct {
	Result     bool         		`json:"result"`
	Date       []CheckUsrData  		`json:"data"`
	Pagination CheckUsrPagination 	`json:"pagination"`
}

//transfer
type TransferResponse struct {
	Result     bool          `json:"result"`
	Date       DateResponse  `json:"data"`
}

//check transfer
type CheckTransferdate struct {
	Code 		string		`json:"code"`
	Message 	string 		`json:"message"`
	Status      int    		`json:"status"`    
}

type CheckTransferResponse struct {
	Result     bool         	 `json:"result"`
	CheckDate CheckTransferdate  `json:"data"`
}

//TransferRecord
type TransferRecordData struct {
	UserName    	string  	`json:"UserName"`
	CreateTime     	string 		`json:"CreateTime"`
	TransType	    string		`json:"TransType"`
	Amount 			string 		`json:"Amount"`
	Balance			string 		`json:"Balance"`
	Currency		string 		`json:"Currency"`
	TransID  		string 		`json:"TransID"`   
}
type TransferRecordPagination struct{
	Page    		int 	`json:"Page"`
	PageLimit     	int 	`json:"PageLimit"`
	TotalNumber	    int	    `json:"TotalNumber"`
	TotalPage 		int 	`json:"TotalPage"`
}
type TransferRecordResponse struct {
	Result     bool         		    `json:"result"`
	Data       []TransferRecordData  	`json:"data"`
	Pagination TransferRecordPagination `json:"pagination"`
}

//PlayGameResponse
type PlayGameResponse struct {
	Result     bool          `json:"result"`
	Date       DateResponse  `json:"data"`
}

//BetRecord
type BetRecordData struct {
	UserName    	string  	`json:"UserName"`
	WagersID     	string 		`json:"WagersID"`
	WagersDate	    string		`json:"WagersDate"`
	SerialID 		string 		`json:"SerialID"`
	RoundNo			string 		`json:"RoundNo"`
	GameType		string 		`json:"GameType"`
	WagerDetail  	string 		`json:"WagerDetail"`   
	GameCode		string		`json:"GameCode"`   
	Result 			string 		`json:"Result"`   
	ResultType 		string 		`json:"ResultType"`   
	Card 			string 		`json:"Card"`   
	BetAmount 		string 		`json:"BetAmount"`   
	Payoff 			string		`json:"Payoff"`   
	Currency 		string 		`json:"Currency"`   
	ExchangeRate 	string 		`json:"ExchangeRate"`   
	Commissionable 	string 		`json:"Commissionable"`   
	Origin 			string 		`json:"Origin"`   
}
type BetRecordPagination struct{
	Page    		int 	`json:"Page"`
	PageLimit     	int 	`json:"PageLimit"`
	TotalNumber	    string	`json:"TotalNumber"`
	TotalPage 		int  	`json:"TotalPage"`
}
type BetRecordResponse struct {
	Result     bool         		`json:"result"`
	Data       []BetRecordData  	`json:"data"`
	Pagination BetRecordPagination  `json:"pagination"`
}

//BetRecordByModifiedDate3
type BetRecordByModifiedDate3Data struct {
	UserName    	string  	`json:"UserName"`
	WagersID     	string 		`json:"WagersID"`
	WagersDate	    string		`json:"WagersDate"`
	SerialID 		string 		`json:"SerialID"`
	RoundNo			string 		`json:"RoundNo"`
	GameType		string 		`json:"GameType"`
	WagerDetail  	string 		`json:"WagerDetail"`   
	GameCode		string		`json:"GameCode"`   
	Result 			string 		`json:"Result"`   
	ResultType 		string 		`json:"ResultType"`   
	Card 			string 		`json:"Card"`   
	BetAmount 		string 		`json:"BetAmount"`   
	Payoff 			string		`json:"Payoff"`   
	Currency 		string 		`json:"Currency"`   
	ExchangeRate 	string 		`json:"ExchangeRate"`   
	Commissionable 	string 		`json:"Commissionable"` 
	Commission      string 		`json:"Commission"` 
	Origin 			string 		`json:"Origin"`   
	ModifiedDate	string      `json:"ModifiedDate"`
}
type BetRecordByModifiedDate3Pagination struct{
	Page    		int 	`json:"Page"`
	PageLimit     	int 	`json:"PageLimit"`
	TotalNumber	    int	    `json:"TotalNumber"`
	TotalPage 		int 	`json:"TotalPage"`
}
type BetRecordByModifiedDate3Response struct {
	Result     bool         		`json:"result"`
	Data       BetRecordData  		`json:"data"`
	Pagination BetRecordPagination  `json:"pagination"`
}

//GetJPHistory
type GetJPHistoryData struct {
	WagersID		string		`json:"WagersID"`
	JPTypeID		string		`json:"JPTypeID"`
	UserName    	string  	`json:"UserName"`
	WagersDate	    string		`json:"WagersDate"`
	JPAmount		string		`json:"JPAmount"`
	GameType		string 		`json:"GameType"`
}
type GetJPHistoryPagination struct{
	Page    		int 	`json:"Page"`
	PageLimit     	int 	`json:"PageLimit"`
	TotalNumber	    string	`json:"TotalNumber"`
	TotalPage 		int 	`json:"TotalPage"`
}
type GetJPHistoryResponse struct {
	Result     bool         		   `json:"result"`
	Data       GetJPHistoryData  	   `json:"data"`
	Pagination GetJPHistoryPagination  `json:"pagination"`
}

//WagersRecordBy38
type WagersRecordBy38Data struct {
	UserName    	string  	`json:"UserName"`
	WagersID		string		`json:"WagersID"`
	WagersDate	    string		`json:"WagersDate"`
	GameType		string 		`json:"GameType"`
	Result			string 		`json:"Result"`
	SerialID		string 		`json:"SerialID"`
	BetAmount		string 		`json:"BetAmount"`
	Payoff			string 		`json:"Payoff"`
	Currency		string 		`json:"Currency"`
	ExchangeRate	string 		`json:"ExchangeRate"`
	Commissionable	string 		`json:"Commissionable"`
	ModifiedDate	string 		`json:"ModifiedDate"`	
}
type WagersRecordBy38Pagination struct{
	Page    		int 	`json:"Page"`
	PageLimit     	int 	`json:"PageLimit"`
	TotalNumber	    string  `json:"TotalNumber"`
	TotalPage 		int 	`json:"TotalPage"`
}
type WagersRecordBy38Response struct {
	Result     bool         		       `json:"result"`
	Data       []WagersRecordBy38Data  	   `json:"data"`
	Pagination WagersRecordBy38Pagination  `json:"pagination"`
}

//WagersRecordBy30
type WagersRecordBy30Data struct {
	UserName    	string  	`json:"UserName"`
	WagersID		string		`json:"WagersID"`
	WagersDate	    string		`json:"WagersDate"`
	SerialID		string 		`json:"SerialID"`
	GameType		string 		`json:"GameType"`
	Result			string 		`json:"Result"`
	BetAmount		string 		`json:"BetAmount"`
	Commissionable	string 		`json:"Commissionable"`
	Payoff			string 		`json:"Payoff"`
	Currency		string 		`json:"Currency"`
	ExchangeRate	string 		`json:"ExchangeRate"`
	ModifiedDate	string 		`json:"ModifiedDate"`	
}
type WagersRecordBy30Pagination struct{
	Page    		int 	`json:"Page"`
	PageLimit     	int 	`json:"PageLimit"`
	TotalNumber	    string  `json:"TotalNumber"`
	TotalPage 		int 	`json:"TotalPage"`
}
type WagersRecordBy30Response struct {
	Result     bool         		       `json:"result"`
	Data       []WagersRecordBy30Data  	   `json:"data"`
	Pagination WagersRecordBy30Pagination  `json:"pagination"`
}

//RKGetFishEvent
type GetFishEventData struct {
	ID				string 		`json:"ID"`
	UserName    	string  	`json:"UserName"`
	CreateTime		string		`json:"CreateTime"`
	Amount	        string		`json:"Amount"`
}
type GetFishEventPagination struct{
	Page    		int 	`json:"Page"`
	PageLimit     	int 	`json:"PageLimit"`
	TotalNumber	    string  `json:"TotalNumber"`
	TotalPage 		int 	`json:"TotalPage"`
}
type GetFishEventResponse struct {
	Result     bool         		   `json:"result"`
	Data       []GetFishEventData      `json:"data"`
	Pagination GetFishEventPagination  `json:"pagination"`
}

//GetSKEventHistory
type GetSKEventHistoryData struct {
	ID				string 		`json:"ID"`
	UserName    	string  	`json:"UserName"`
	CreateTime		string		`json:"CreateTime"`
	Amount	        string		`json:"Amount"`
}
type GetSKEventHistoryPagination struct{
	Page    		int 	`json:"Page"`
	PageLimit     	int 	`json:"PageLimit"`
	TotalNumber	    string  `json:"TotalNumber"`
	TotalPage 		int 	`json:"TotalPage"`
}
type GetSKEventHistoryResponse struct {
	Result     bool         		         `json:"result"`
	Data       []GetSKEventHistoryData    	 `json:"data"`
	Pagination GetSKEventHistoryPagination   `json:"pagination"`
}
//GetSKEventHistory response struct 
type GetSportEventHistoryData struct {
	ID				string 		`json:"ID"`
	UserName    	string  	`json:"UserName"`
	CreateTime		string		`json:"CreateTime"`
	Amount	        string		`json:"Amount"`
}
type GetSportEventHistoryPagination struct{
	Page    		int 	`json:"Page"`
	PageLimit     	int 	`json:"PageLimit"`
	TotalNumber	    int	    `json:"TotalNumber"`
	TotalPage 		int 	`json:"TotalPage"`
}
type GetSportEventHistoryResponse struct {
	Result     bool         		         `json:"result"`
	Data       []GetSKEventHistoryData    	 `json:"data"`
	Pagination GetSKEventHistoryPagination   `json:"pagination"`
}

//公共接口API
/***************************
*  字符串转换为MD5字符串
*initstr： 需要进行MD5的字符串
* return:  返回MD5字符串 如：7ac66c0f148de9519b8bd264312c4d64
****************************/
func RuokChangeMD5(initstr string) (error,string) {
	if initstr == "" {
		return errors.New("initstr string is null, please check! "),""
	}
	str := md5.New()
    io.WriteString(str, initstr)   //initstr
  
	bytestr := fmt.Sprintf("%x", str.Sum(nil))  //w.Sum(nil)将w的hash转成[]byte格式
    md5str := string(bytestr[:])  //字符数转换为字符串

    return nil,md5str
}

/***************************
* 将发送请求的结构转换为JSON 
*reqdata： 请求的结构体数据
* return:  返回Json字符串
****************************/
func RuokStructTOJSON(reqdata interface{}) string{
	//TODO
	// if reqdata == nil {
	// 	fmt.Println("request date is null \n")
	// 	return errors.New("input string is null")
	// }
	  data, _:= json.Marshal(reqdata)
	  OutJSON := string(data)
	  
	  return OutJSON
}

/***************************
* 获取美国东部时间，以纽约时间为准
* return:  返回时间字符串 如：20180515
****************************/
 func RuokGetTimestamps() string {

	formate:="20060102"  //输出时间格式
	now := time.Now()
	
	local, err := time.LoadLocation("America/New_York")   // 美国东部时间，获取纽约西四区时间
    if err != nil {
        fmt.Println(err)
	}
	date := fmt.Sprintf("%s",now.In(local).Format(formate))
	
	if debug {
		fmt.Println(now.In(local).Format("2006-01-02 15:04:05"))
	}
	
	return date 
 }

/************ 生成KEY *******
* return:  返回时间字符串 如：o748a578593afdfefc94f6bcb7aa1b695ruokplat
****************************/
//验证码(需全小写)，组成方式如下: key=A+B+C(验证码组合方式)
//A= 无意义字串长度7码
//B= MD5(website+ username + KeyB + YYYYMMDD) 
//C= 无意义字串长度1码 YYYYMMDD为美东时间(GMT-4)(20180508)
func RuokInstallQuestKey(Astr,Cstr,md5str string) (error,string) {
	
	if md5str == "" || Astr == "" || Cstr == "" {
		return errors.New("md5 Astr or Cstr string is null, can't install key"),""
	}
	Bstr :=md5str
	endstr := Astr + Bstr + Cstr
	
	return nil,endstr
}

// 测试参数需要时间  str1 返回格式： 2018-05-17 str2返回格式：15:04:05
func Rkt_Date_time() (string,string) {
	// str1 := fmt.Sprintf("%s",time.Now().Format("2006-01-02 15:04:05"))
	str1 := fmt.Sprintf("%s",time.Now().Format("2006-01-02"))
	str2 := fmt.Sprintf("%s",time.Now().Format("15:04:05"))

	if debug {
		fmt.Println(str1)
		fmt.Println(str2)
	}
	return str1, str2
 }