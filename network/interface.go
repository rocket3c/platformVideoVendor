package network

import (
	"fmt"
	"io/ioutil"
	"errors"
	"net/http"
	"encoding/json"
	//"crypto/md5"
	//network "github.com/ruoklive/platformGame/network"
)

func Ruok_CreateMember(create *CreateMember) (*CreateResponse,error) {

	//生成一个MD5字符串
	initstr := Website + create.Username + Createmember_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}
  
	// 根据规则生成一个KEY
	astr := "ruokplt"  //无意义字串长度7码
	cstr := "o"        //无意义字串长度1码
	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)
	if err != nil {
		return nil,errors.New("install key can't finish")
	}
	
	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%s&password=%s&key=%s",
		Createmember_url, Website, create.Username, Uppername, create.Password, Key)

	//test print 
	if debug { 
		fmt.Println("request url:")
		fmt.Println(url)
	}
	client := &http.Client{}

	request, _ := http.NewRequest("GET", url, nil)
	
	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("createmember response do process fail")
	}	
	defer response.Body.Close()

	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
        bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("createmember response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	//test print
	if debug { 
		fmt.Println("response string:")
		fmt.Println(string(bodyBytes))
	}

	//解析接收到是json
	var RkCreateMember  CreateResponse
    err = json.Unmarshal(bodyBytes, &RkCreateMember)

	fmt.Printf("result:%+v code:%s message:%s\n",RkCreateMember.Result, RkCreateMember.Date.Code,RkCreateMember.Date.Message)
    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}	
	return &RkCreateMember, nil
}
func Ruok_Login(Login *RKLogin) error {

	//生成一个MD5字符串
	initstr := Website + Login.Username + Login_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return errors.New("md5 string is null")
	}
 
	// 根据规则生成一个KEY
	astr := "ruokplat"  //无意义字串长度8码
	cstr := "o"        //无意义字串长度1码
	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return errors.New("install key can't finish")
	}
	if debug {
		fmt.Println(Login.Key)
	}
	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%s&lang=%s&key=%s",
		Login_url, Website, Login.Username, Uppername, Login.Lang, Key)  //get 数据拼接

	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}
	client := &http.Client{}
	
	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return errors.New("Login NewRequest fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return errors.New("Login response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	//var bodyBytes []byte

    if response.StatusCode == 200 {
		_, err = ioutil.ReadAll(response.Body)
        //bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return errors.New("Login response readall process fail")
    } else {
        return fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	//test print
	// fmt.Println("bodyBytes:")
	// fmt.Println(string(bodyBytes))

	//解析接收到是json
	// var Rklogin  LoginResponse
    // err = json.Unmarshal(bodyBytes, &Rklogin)

	// fmt.Printf("result:%+v code:%s message:%s\n",Rklogin.Result, Rklogin.Date.Code,Rklogin.Date.Message)
    // if err != nil {
    //     return nil, fmt.Errorf("unable to parse the JSON response:", err)
	// }	
	return nil
}
func Ruok_Login2(Login2 *RKLogin2) (*Login2Response, error) {

	//生成一个MD5字符串
	initstr := Website + Login2.Username + Login2_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}
  
	astr := "ruokplat"  //无意义字串长度8码
	cstr := "o"        //无意义字串长度1码
	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return nil,errors.New("install key can't finish")
	}
	
	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%s&lang=%s&key=%s",
		Login2_url, Website, Login2.Username, Uppername, Login2.Lang, Key)  //get 数据拼接

	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}
	client := &http.Client{}
	
	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("Login2 NewRequest fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("Login response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
        bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("Login2 response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	//test print
	if debug {
		fmt.Println("response string:")
		fmt.Println(string(bodyBytes))
	}
	//解析接收到是json
	var Login2Res  Login2Response
	err = json.Unmarshal(bodyBytes, &Login2Res)
    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	//return:  result:false code:44000 message:Key error
	fmt.Printf("result:%+v code:%v message:%s\n",Login2Res.Result, Login2Res.Date.Code,Login2Res.Date.Message)

	return &Login2Res, nil
}
func Ruok_Logout(Logout *RKLogout) (*LogoutResponse,error) {
	
	//生成一个MD5字符串
	initstr := Website + Logout.Username + Logout_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}
	astr := "ruokpla"  //无意义字串长度7码
	cstr := "ruokpl"  //无意义字串长度6码

	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return nil,errors.New("install key can't finish")
	}
	
	url := fmt.Sprintf("%s?website=%s&username=%s&key=%s",Logout_url, Website, Logout.Username, Key)  //get 数据拼接
	
	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}
	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("logout newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("logout response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("Logout response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	if debug {
		fmt.Println("response string:")
	 	fmt.Println(string(bodyBytes))
	}
	//解析接收到是json
	var Rklogout  LogoutResponse
	err = json.Unmarshal(bodyBytes, &Rklogout)
    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	//return:  result:false code:44000 message:Key error
	fmt.Printf("result:%+v code:%s message:%s\n",Rklogout.Result, Rklogout.Date.Code,Rklogout.Date.Message)

	return &Rklogout, nil
}
func Ruok_CheckUsrBalance(CheckUsrBalance *RKCheckUsrBalance) (*CheckUsrResponse,error) {
	
	//生成一个MD5字符串
	initstr := Website + CheckUsrBalance.Username + CheckUsrBalance_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}
	
	astr := "plat"     //无意义字串长度4码
	cstr := "ruokplt"  //无意义字串长度7码

	err,CheckUsrBalance.Key = RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return nil,errors.New("install key can't finish")
	}
	
	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%v&page=%v&pagelimit=%v&key=%s",
		Checkuserbanace_url, Website, CheckUsrBalance.Username, CheckUsrBalance.Uppername,
		CheckUsrBalance.Page, CheckUsrBalance.Pagelimit, CheckUsrBalance.Key)  //get 数据拼接
	
	//test print
	if debug { 
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("checkbalance newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("checkbalance response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("checkusrbalance response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	if debug {
		fmt.Println("response string:")
		 fmt.Println(string(bodyBytes))
	}

	//解析接收到是json
	var RkCheckUsr  CheckUsrResponse
	err = json.Unmarshal(bodyBytes, &RkCheckUsr)
    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	//return:  result:false code:44000 message:Key error
	fmt.Printf("test print result:%+v loginname:%s Page:%d\n",RkCheckUsr.Result, RkCheckUsr.Date[0].LoginName, RkCheckUsr.Pagination.Page)
	return &RkCheckUsr, nil
}
func Ruok_Transfer(Transfer *RKTransfer) (*TransferResponse,error) {

	initstr := Website + Transfer.Username + Transfer.Remitno + Transfer_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}

	astr := "ruokplatg"  						//无意义字串长度9码
	cstr := "ruok" 								 //无意义字串长度4码
	
	 errkey,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if errkey != nil {
		return nil,errors.New("install key can't finish")
	}
	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%v&remitno=%v&action=%v&remit=%v&key=%s",
		Transfer_url, Website, Transfer.Username, Transfer.Uppername, Transfer.Remitno, Transfer.Action, Transfer.Remitno, Key)  //get 数据拼接
	
	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("transfer newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("transfer response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("transfer response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}

	if debug {
		fmt.Println("response string:")
	 	fmt.Println(string(bodyBytes))
	}
	//解析接收到的json
	var Rktransfer  TransferResponse
	err = json.Unmarshal(bodyBytes, &Rktransfer)

	//test print
	fmt.Printf("result:%+v code:%s message:%s\n",Rktransfer.Result, Rktransfer.Date.Code,Rktransfer.Date.Message)

    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	return &Rktransfer, nil
}
func Ruok_CheckTransfer(CheckTransfer *RKCheckTransfer) (*CheckTransferResponse,error) {
	
	initstr := Website + CheckTransfer_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}
	astr := "ruokplat"  						     //无意义字串长度8码
	cstr := "ruokacde" 								 //无意义字串长度8码
	
	 errkey,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if errkey != nil {
		return nil,errors.New("install key can't finish")
	}
	url := fmt.Sprintf("%s?website=%s&transid=%s&key=%s",CheckTransfer_url,Website,CheckTransfer.Transid , Key)  //get 数据拼接
	
	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}
	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, err
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("checktransfer response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("checktranfer response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}

	if debug {
		fmt.Println("response string:")
		fmt.Println(string(bodyBytes))
	}

	//解析接收到的json
	var RkCheckTransfer CheckTransferResponse
	err = json.Unmarshal(bodyBytes, &RkCheckTransfer)

	//test print
	//fmt.Printf("result:%+v code:%s message:%s\n",Rktransfer.Result, Rktransfer.Date.Code,Rktransfer.Date.Message)

    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	return &RkCheckTransfer, nil
}
func Ruok_TransferRecord(TransferRecord *RKTransferRecord) (*TransferRecordResponse,error) {
	//md5生成
	initstr := Website + TransferRecord.Username + TransferRecord_keyB + RuokGetTimestamps()

	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}
	astr := "ruokplat"  //无意义字串长度8码
	cstr := "ruokplta"  //无意义字串长度8码

	err,key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return nil,errors.New("install key can't finish")
	}
	
	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%v&transid=%v&transtype=%v&date_start=%v&date_end=%v&start_hhmmss=%v&end_hhmmss=%v&page=%v&pagelimit=%v&key=%v",
		TransferRecord_url, Website, TransferRecord.Username, TransferRecord.Uppername, TransferRecord.Transid, TransferRecord.Transtype,
		TransferRecord.Date_start,TransferRecord.Date_end, TransferRecord.Start_hhmmss, TransferRecord.End_hhmmss,TransferRecord.Page, TransferRecord.Pagelimit, key)  //get 数据拼接
	
	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("transferrecord newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("transferrecord response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("transferrecord response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}

	if debug {
		fmt.Println("response string:")
	    fmt.Println(string(bodyBytes))
	}

	//解析接收到是json
	var RkTransferRecord  TransferRecordResponse
	err = json.Unmarshal(bodyBytes, &RkTransferRecord)
    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	//return:  result:false code:44000 message:Key error
	fmt.Printf("result:%+v\n",RkTransferRecord.Result)

	return &RkTransferRecord, nil
}
func Ruok_PlayGame(PlayGame *RKPlayGame) error { 

	//md5生成
	initstr := Website + PlayGame.Username + PlayGame_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return errors.New("md5 string is null")
	}
 
	// 根据规则生成一个KEY
	astr := "ruokplat"  //无意义字串长度8码
	cstr := "oasfxbde"        //无意义字串长度8码
	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&username=%s&gamekind=%v&gametype=%v&gamecode=%v&lang=%v&key=%s",
		PlayGame_url, Website, PlayGame.Username, PlayGame.Gamekind, PlayGame.Gametype, PlayGame.Gamecode, PlayGame.Lang, Key)  //get 数据拼接

	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}
	
	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return errors.New("playgame newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return errors.New("playgame response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
    if response.StatusCode == 200 {
		_, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return errors.New("playgame response readall process fail")
    } else {
        return fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	return nil
}
func Ruok_PlayGameByH5(PlayGameByH5 *RKPlayGameH5) error { 

	//md5生成
	initstr := Website + PlayGameByH5.Username + PlayGameByH5_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return errors.New("md5 string is null")
	}
	// 根据规则生成一个KEY
	astr := "ruokplat"  		 //无意义字串长度8码
	cstr := "oasfxbde"       	 //无意义字串长度8码
	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&username=%s&gamekind=%s&gametype=%d&gamecode=%d&lang=%s&key=%s",
		PlayGameByH5_url,Website,PlayGameByH5.Username,PlayGameByH5.Gamekind,PlayGameByH5.Gametype,PlayGameByH5.Lang,Key)  //get 数据拼接

	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}
	
	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return errors.New("playgamebyh5 newresponse process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return errors.New("playgamebyh5 response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
    if response.StatusCode == 200 {
		_, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return errors.New("playgamebyh5 response readall process fail")
    } else {
        return fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	return nil
}
func Ruok_BetRecord(BetRecord *RKBetRecord) (*BetRecordResponse,error) {

	//md5生成
	initstr := Website + BetRecord.Username + BetRecord_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}
	astr := "o"  						//无意义字串长度1码
	cstr := "ruokplat" 								 //无意义字串长度8码
	
	 errkey,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if errkey != nil {
		return nil,errors.New("install key can't finish")
	}
	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%s&rounddate=%v&starttime=%v&endtime=%v&gamekind=%v&subgamekind=%v&gametype=%v&page=%v&pagelimit=%v&key=%s",
		BetRecord_url, Website,BetRecord.Username, BetRecord.Uppername, BetRecord.Rounddate, BetRecord.Starttime, BetRecord.Endtime, BetRecord.Gamekind,
		BetRecord.Subgamekind, BetRecord.Gametype, BetRecord.Page, BetRecord.Pagelimit, Key)  //get 数据拼接
	
	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("betrecord newresponse process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("betrecord response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("betrecord response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	if debug {
		fmt.Println("response string:")
	 	fmt.Println(string(bodyBytes))
	}

	//解析接收到的json
	var RkBetRecord BetRecordResponse
	err = json.Unmarshal(bodyBytes, &RkBetRecord)

    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	return &RkBetRecord, nil
}
func Ruok_BetRecordByModifiedDate3(BetRecordByModifiedDate3 *RKBetRecordByModifiedDate3) (*BetRecordByModifiedDate3Response,error) {

	//md5生成
	initstr := Website + BetRecordByModifiedDate3_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}

	astr := "o"  						//无意义字串长度1码
	cstr := "ruokplat" 								 //无意义字串长度8码
	
	 errkey,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if errkey != nil {
		return nil,errors.New("install key can't finish")
	}
	url := fmt.Sprintf("%s?website=%s&start_date=%v&end_date=%s&starttime=%v&endtime=%v&gamekind=%v&subgamekind=%v&gametype=%v&page=%v&pagelimit=%v&key=%s",
		BetRecordByModifiedDate3_url,Website,BetRecordByModifiedDate3.Start_date,BetRecordByModifiedDate3.End_date,BetRecordByModifiedDate3.Starttime,
		BetRecordByModifiedDate3.Endtime,BetRecordByModifiedDate3.Gamekind,BetRecordByModifiedDate3.Subgamekind,BetRecordByModifiedDate3.Gametype,
		BetRecordByModifiedDate3.Page,BetRecordByModifiedDate3.Pagelimit,Key)  //get 数据拼接
	
	//test print
	if debug { 
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("betrecordbymodifieddate3 newrequest process fail") 
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("betrecordbymodifieddate3 response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("betrecordbymodifieddate3 response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}

	if debug {
		fmt.Println("response string:")
	 	fmt.Println(string(bodyBytes))
	}

	//解析接收到的json
	var RKBetRecordByModifiedDate3 BetRecordByModifiedDate3Response
	err = json.Unmarshal(bodyBytes, &BetRecordByModifiedDate3)

    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	return &RKBetRecordByModifiedDate3, nil
}
func Ruok_GetJPHistory(GetJPHistory *RKGetJPHistory) (*GetJPHistoryResponse,error) {

	//md5生成
	initstr := Website + GetJPHistory_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}
	astr := "o"  						//无意义字串长度1码
	cstr := "ruokplat" 								 //无意义字串长度8码
	
	 errkey,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if errkey != nil {
		return nil,errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&start_date=%+v&end_date=%s&starttime=%v&endtime=%v&jptype=%v&page=%v&pagelimit=%v&key=%s",
		GetJPHistory_url,Website,GetJPHistory.Start_date,GetJPHistory.End_date,GetJPHistory.Starttime,
		GetJPHistory.Endtime,GetJPHistory.Jptype,GetJPHistory.Page,GetJPHistory.Pagelimit,Key)  //get 数据拼接
	
	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("getjphistory newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("getjphistory response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("getjphistory response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}

	if debug {
		fmt.Println("response string:")
	 	fmt.Println(string(bodyBytes))
	}

	//解析接收到的json
	var RkGetJPHistory GetJPHistoryResponse
	err = json.Unmarshal(bodyBytes, &RkGetJPHistory)

    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	return &RkGetJPHistory, nil
}
func Ruok_ForwardGameH5By5(ForwardGameH5By5 *RKForwardGameH5By5) error {

	//md5生成
	initstr := Website + ForwardGameH5By5.Username + ForwardGameH5By5_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return errors.New("md5 string is null")
	}
 
	// 根据规则生成一个KEY
	astr := "ruokplat"  //无意义字串长度8码
	cstr := "oRUISDFG"        //无意义字串长度8码

	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%s&gametype=%s&key=%s",
		ForwardGameH5By5_url,Website,ForwardGameH5By5.Username,Uppername,ForwardGameH5By5.Gametype,Key)  //get 数据拼接

	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}
	
	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return  errors.New("forwardgameh5by5 newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return errors.New("forwardgameh5by5 response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
    if response.StatusCode == 200 {
	   //bodyBytes, err = ioutil.ReadAll(response.Body)
	_, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return errors.New("forwardgameh5by5 response readall process fail")
    } else {
        return fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	return nil
}
func Ruok_WagersRecordBy38(WagersRecordBy38 *RKWagersRecordBy38) (*WagersRecordBy38Response,error) {
	
    //md5生成
	initstr := Website +  WagersRecordBy38_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}
 
	astr := "o"  						//无意义字串长度1码
	cstr := "ruokplat" 					 //无意义字串长度8码
	
	 errkey,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if errkey != nil {
		return nil,errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&action=%v&username=%v&uppername=%v&start_date=%+v&end_date=%s&starttime=%v&endtime=%v&gametype=%v&page=%v&pagelimit=%v&key=%s",
		WagersRecordBy38_url,Website,WagersRecordBy38.Action,WagersRecordBy38.Username,WagersRecordBy38.Uppername, WagersRecordBy38.Start_date,WagersRecordBy38.End_date,
		WagersRecordBy38.Starttime,WagersRecordBy38.Endtime,WagersRecordBy38.Gametype,WagersRecordBy38.Page,WagersRecordBy38.Pagelimit,Key)  //get 数据拼接
	
	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("wagerrecordby38 newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("wagerrecordby38 response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("wagerrecordby38 response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	if debug {
		fmt.Println("response string:")
	 	fmt.Println(string(bodyBytes))
	}

	//解析接收到的json
	var RkWagersRecordBy38 WagersRecordBy38Response
	err = json.Unmarshal(bodyBytes, &RkWagersRecordBy38)

    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	return &RkWagersRecordBy38, nil
}
func Ruok_WagersRecordBy30(WagersRecordBy30 *RKWagersRecordBy30) (*WagersRecordBy30Response,error) {

	//md5生成
	initstr := Website + WagersRecordBy30_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}

	astr := "o"  												//无意义字串长度1码
	cstr := "ruokplat" 										    //无意义字串长度8码
	
	 errkey,Key := RuokInstallQuestKey(astr,cstr,md5vaule)      // 根据规则生成一个KEY	
	if errkey != nil {
		return nil,errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&action=%v&uppername=%v&start_date=%+v&end_date=%s&starttime=%v&endtime=%v&gametype=%v&page=%v&pagelimit=%v&key=%s",
		WagersRecordBy30_url,Website,WagersRecordBy30.Action,WagersRecordBy30.Uppername,WagersRecordBy30.Start_date,WagersRecordBy30.End_date,
		WagersRecordBy30.Starttime,WagersRecordBy30.Endtime,WagersRecordBy30.Gametype,WagersRecordBy30.Page,WagersRecordBy30.Pagelimit,Key)  //get 数据拼接
	
	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("wagerrecordby30 newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("wagerrecordby30 response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("wagerrecordby30 response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}

	if debug {
		fmt.Println("response string:")
	 	fmt.Println(string(bodyBytes))
	}
	//解析接收到的json
	var RkWagersRecordBy30 WagersRecordBy30Response
	err = json.Unmarshal(bodyBytes, &RkWagersRecordBy30)

    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	return &RkWagersRecordBy30, nil
}
func Ruok_ForwardGameH5By30(ForwardGameH5By30 *RKForwardGameH5By30) error {

	//md5生成
	initstr := Website + ForwardGameH5By30.Username+ ForwardGameH5By30_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return errors.New("md5 string is null")
	}
 
	// 根据规则生成一个KEY
	astr := "ruokplat"  //无意义字串长度8码
	cstr := "oRUISDFG"        //无意义字串长度8码

	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%s&gametype=%s&key=%s",
		ForwardGameH5By30_url,Website,ForwardGameH5By30.Username,Uppername,ForwardGameH5By30.Gametype,Key)  //get 数据拼接

	client := &http.Client{}
	
	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return  errors.New("forwardgameh5y30 newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return errors.New("forwardgameh5y30 response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	//var bodyBytes []byte
    if response.StatusCode == 200 {
	   //bodyBytes, err = ioutil.ReadAll(response.Body)
	_, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return errors.New("forwardgameh5y30 response readall process fail")
    } else {
        return fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	return nil
}
func Ruok_ForwardGameH5By38(ForwardGameH5By38 *RKForwardGameH5By38) error {

	//md5生成
	initstr := Website + ForwardGameH5By38.Username + ForwardGameH5By38_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return errors.New("md5 string is null")
	}
 
	// 根据规则生成一个KEY
	astr := "ruokplat"  		//无意义字串长度8码
	cstr := "oRUISDFG"          //无意义字串长度8码

	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%s&gametype=%s&key=%s",
		ForwardGameH5By38_url,Website,ForwardGameH5By38.Username,Uppername,ForwardGameH5By38.Gametype,Key)  //get 数据拼接

	client := &http.Client{}
	
	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return  errors.New("forwardgameh5y38 newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return errors.New("forwardgameh5y38 response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
    if response.StatusCode == 200 {
	_, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return errors.New("forwardgameh5y38 response readall process fail")
    } else {
        return fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	return nil
}
func Ruok_GetFishEventHistory(GetFishEventHistory *RKGetFishEventHistory) (*GetFishEventResponse,error) {

	initstr := Website + GetFishEventHistory_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}

	astr := "o"  						//无意义字串长度1码
	cstr := "ruokplat" 								 //无意义字串长度8码
	
	 errkey,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if errkey != nil {
		return nil,errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&rounddate=%v&starttime=%v&endtime=%v&page=%v&pagelimit=%v&key=%s",
		GetFishEventHistory_url,Website,GetFishEventHistory.Rounddate,GetFishEventHistory.Starttime,
		GetFishEventHistory.Endtime,GetFishEventHistory.Page,GetFishEventHistory.Pagelimit,Key)  //get 数据拼接
	
	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}
	
	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("getfishevethistoy newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("getfishevethistoy response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("getfishevethistoy response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}

	if debug {
		fmt.Println("response string:")
	 	fmt.Println(string(bodyBytes))
	}

	//解析接收到的json
	var RKGetFishEvent GetFishEventResponse
	err = json.Unmarshal(bodyBytes, &RKGetFishEvent)

    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	return &RKGetFishEvent, nil
}
func Ruok_FishEventUrl(FishEventUrl *RKFishEventUrl) error {

	initstr := Website + FishEventUrl.Username + FishEventUrl_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return errors.New("md5 string is null")
	}
	// 根据规则生成一个KEY
	astr := "ruokplat"  //无意义字串长度8码
	cstr := "oRUISDFG"        //无意义字串长度8码

	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%s&lang=%s&key=%s",
		FishEventUrl_url,Website,FishEventUrl.Username,Uppername,FishEventUrl.Lang,Key)  //get 数据拼接

	client := &http.Client{}
	
	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return  errors.New("getfishevet newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return errors.New("getfishevet response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
    if response.StatusCode == 200 {
	_, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return errors.New("getfishevet response readall process fail")
    } else {
        return fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	return nil
}
func Ruok_GetSKEventHistory(GetSKEventHistory *RKGetSKEventHistory) (*GetSKEventHistoryResponse,error) {

	initstr := Website +  GetSKEventHistory_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}
	astr := "o"  						//无意义字串长度1码
	cstr := "ruokplat" 								 //无意义字串长度8码
	
	 errkey,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if errkey != nil {
		return nil,errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&rounddate=%v&starttime=%v&endtime=%v&page=%v&pagelimit=%v&key=%s",
		GetSKEventHistory_url,Website,GetSKEventHistory.Rounddate,GetSKEventHistory.Starttime,
		GetSKEventHistory.Endtime,GetSKEventHistory.Page,GetSKEventHistory.Pagelimit,Key)  //get 数据拼接
	
	//test print
	if debug { 
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("getskeventistory newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("getskeventistory response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("getskeventistory response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}

	if debug {
		fmt.Println("response string:")
	 	fmt.Println(string(bodyBytes))
	}

	//解析接收到的json
	var RKGetSKEventHistory GetSKEventHistoryResponse
	err = json.Unmarshal(bodyBytes, &RKGetSKEventHistory)

    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	return &RKGetSKEventHistory, nil
}
func Ruok_SKEventUrl(SKEventUrl *RKSKEventUrl) error {

	initstr := Website + SKEventUrl.Username + SKEventUrl_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return errors.New("md5 string is null")
	}
 
	// 根据规则生成一个KEY
	astr := "ruokplat"  //无意义字串长度8码
	cstr := "oRUISDFG"        //无意义字串长度8码

	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%s&lang=%s&key=%s",
		SKEventUrl_url,Website,SKEventUrl.Username,Uppername,SKEventUrl.Lang,Key)  //get 数据拼接

	client := &http.Client{}
	
	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return  errors.New("skeventurl newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return errors.New("skeventurl response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
    if response.StatusCode == 200 {
	_, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return errors.New("skeventurl response readall process fail")
    } else {
        return fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	return nil
}
func Ruok_GetSportEventHistory(GetSportEventHistory *RKGetSportEventHistory) (*GetSportEventHistoryResponse,error) {

	initstr := Website + GetSportEventHistory_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return nil,errors.New("md5 string is null")
	}

	astr := "o"  						//无意义字串长度1码
	cstr := "ruokplat" 								 //无意义字串长度8码
	
	 errkey,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if errkey != nil {
		return nil,errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&rounddate=%v&starttime=%v&endtime=%v&page=%v&pagelimit=%v&key=%s",
		GetSportEventHistory_url,Website,GetSportEventHistory.Rounddate,GetSportEventHistory.Starttime,
		GetSportEventHistory.Endtime,GetSportEventHistory.Page,GetSportEventHistory.Pagelimit,Key)  //get 数据拼接
	
	//test print 
	if debug {
		fmt.Println("request url:")
		fmt.Println(url)
	}

	client := &http.Client{}

	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil, errors.New("getsporteventhistory newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return nil,errors.New("getsporteventhistory response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
	var bodyBytes []byte

    if response.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return nil, errors.New("getsporteventhistory response readall process fail")
    } else {
        return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	if debug {
		fmt.Println("response string:")
	 	fmt.Println(string(bodyBytes))
	}

	//解析接收到的json
	var RKGetSportEventHistory GetSportEventHistoryResponse
	err = json.Unmarshal(bodyBytes, &RKGetSportEventHistory)

    if err != nil {
        return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}
	return &RKGetSportEventHistory, nil
}
func Ruok_SportEventUrl(SportEventUrl *RKSportEventUrl) error {

	initstr := Website + SportEventUrl.Username + SportEventUrl_keyB + RuokGetTimestamps()
	err,md5vaule := RuokChangeMD5(initstr)
	if err != nil {
		return errors.New("md5 string is null")
	}
	// 根据规则生成一个KEY
	astr := "ruokplat"  //无意义字串长度8码
	cstr := "oRUISDFG"        //无意义字串长度8码

	err,Key := RuokInstallQuestKey(astr,cstr,md5vaule)   // 根据规则生成一个KEY	
	if err != nil {
		return errors.New("install key can't finish")
	}

	url := fmt.Sprintf("%s?website=%s&username=%s&uppername=%s&lang=%s&key=%s",
		SportEventUrl_url,Website,SportEventUrl.Username,Uppername,SportEventUrl.Lang,Key)  //get 数据拼接

	client := &http.Client{}
	
	//建立连接请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return  errors.New("getsporteventurl newrequest process fail")
    }

	response, err := client.Do(request)
	if err != nil {
    	return errors.New("getsporteventurl response do process fail")
	}	
	defer response.Body.Close()
	
	//接收body
    if response.StatusCode == 200 {
	_, err = ioutil.ReadAll(response.Body)
    } else if err != nil {
        return errors.New("getsporteventurl response readall process fail")
    } else {
        return fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}
	return nil
}