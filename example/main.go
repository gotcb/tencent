package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gotcb/tencent"
)

//Questions query参数
type Questions struct {
	Qaid    string `json:"qaid,omitempty"`
	Fieldid string `json:"fieldid,omitempty"`
}

//Questions query参数
type Datas struct {
	Qaid string `json:"qaid,omitempty"`
}

// func main() {
// 	config := &tencent.Config{
// 		Timestamp: time.Now().Unix(),
// 		SecretId:  "ewfwefwefewLHy******",
// 		SecretKey: "efefefefL1******",
// 	}
// 	tc := tencent.NewTencent(config)
// 	//授权签名
// 	err := tc.GetAuth()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	//获取tcb授权
// 	tcb := tc.GetTcb()
// 	// data, err := tcb.GetDocument(Envid, "questions", "0a3f93955ff8252200ccae8d710181a9")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// fmt.Println(data)

// 	que := Questions{
// 		Qaid: "123",
// 	}
// 	ques, err := json.Marshal(que)
// 	if err != nil {
// 		fmt.Printf("json.Marshal%v", err)
// 	}
// 	data, err := tcb.Find(Envid, "questions", string(ques), "100", "0", "", "", "")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(data)
// }

func main() {
	config := &tencent.Config{
		Envid:     "****-7gal*****",
		Timestamp: time.Now().Unix(),
		SecretId:  "ewfwefwefewLHy******",
 		SecretKey: "efefefefL1******",
	}
	tc := tencent.NewTencent(config)
	//授权签名
	err := tc.GetAuth()
	if err != nil {
		fmt.Println(err)
	}
	//获取tcb授权
	tcb := tc.GetTcb()
	//GetDocument *****************************************************************************************************************************
	// data, err := tcb.GetDocument( "questions", "0a3f93955ff8252200ccae8d710181a9")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)

	// que := Questions{
	// 	Qaid: "123",
	// }
	// ques, err := json.Marshal(que)
	// if err != nil {
	// 	fmt.Printf("json.Marshal%v", err)
	// }
	// data, err := tcb.Find("questions", string(ques), "100", "0", "", "", "")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//Count *****************************************************************************************************************************
	// que := Questions{
	// 	Qaid: "123",
	// }
	// ques, err := json.Marshal(que)
	// if err != nil {
	// 	fmt.Printf("json.Marshal%v", err)
	// }
	// data, err := tcb.Count( "questions", string(ques), "")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//UpdateOne *****************************************************************************************************************************
	// que := Questions{
	// 	Qaid: "123",
	// }

	// que1 := Datas{
	// 	Fieldname: "测试1",
	// }

	// ques, err := json.Marshal(que)
	// if err != nil {
	// 	fmt.Printf("json.Marshal%v", err)
	// }

	// ques1, err := json.Marshal(que1)
	// if err != nil {
	// 	fmt.Printf("json.Marshal%v", err)
	// }

	// data, err := tcb.UpdateOne("questions", string(ques), string(ques1), "")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//UpdateMany *****************************************************************************************************************************
	// que := Questions{
	// 	Qaid:    "123",
	// 	Fieldid: "1034-1609310682698-39711",
	// }

	// que1 := Datas{
	// 	Fieldname: "测试1",
	// 	Content:   "1；2；3；4；5",
	// }

	// ques, err := json.Marshal(que)
	// if err != nil {
	// 	fmt.Printf("json.Marshal%v", err)
	// }

	// ques1, err := json.Marshal(que1)
	// if err != nil {
	// 	fmt.Printf("json.Marshal%v", err)
	// }

	// data, err := tcb.UpdateMany("questions", string(ques), string(ques1), "")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//DeleteOne *****************************************************************************************************************************
	// que := Questions{
	// 	Qaid:    "123",
	// 	Fieldid: "1034-1609310682698-39711",
	// }

	// ques, err := json.Marshal(que)
	// if err != nil {
	// 	fmt.Printf("json.Marshal%v", err)
	// }

	// data, err := tcb.DeleteOne("questions", string(ques), "")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//DeleteMany *****************************************************************************************************************************
	// que := Questions{
	// 	Qaid:    "123",
	// 	Fieldid: "1034-1609310682698-39711",
	// }

	// ques, err := json.Marshal(que)
	// if err != nil {
	// 	fmt.Printf("json.Marshal%v", err)
	// }

	// data, err := tcb.DeleteMany("questions", string(ques), "")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(data)

	//GetDocument *****************************************************************************************************************************
	// data, err := tcb.GetDocument("questions", "0a3f93955ff8252200ccae8e68b5214f", "")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(data)

	//UpdateDocument *****************************************************************************************************************************
	// que := Datas{
	// 	Qaid: "123",
	// }
	// ques, err := json.Marshal(que)
	// if err != nil {
	// 	fmt.Printf("json.Marshal%v", err)
	// }
	// data, err := tcb.UpdateDocument("questions", "0a3f93955ff8252200ccae8e68b5214f", string(ques), "")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)

	//SetDocument *****************************************************************************************************************************
	// que := Datas{
	// 	Qaid: "789",
	// }
	// ques, err := json.Marshal(que)
	// if err != nil {
	// 	fmt.Printf("json.Marshal%v", err)
	// }
	// data, err := tcb.SetDocument("questions", "0a3f93955ff8252200ccae8f3bb567dc", string(ques), "")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)

	//InsertDocument *****************************************************************************************************************************
	que := Datas{
		Qaid: "789",
	}
	ques, err := json.Marshal(que)
	if err != nil {
		fmt.Printf("json.Marshal%v", err)
	}
	data, err := tcb.InsertDocument("questions", "0a3f93955ff8252200ccae904e8cd346sdgs", string(ques), "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	//InsertDocuments *****************************************************************************************************************************
	// que := Datas{
	// 	Qaid: "123465",
	// }
	// ques, err := json.Marshal(que)
	// if err != nil {
	// 	fmt.Printf("json.Marshal%v", err)
	// }
	// var str []string
	// str = append(str, string(ques))
	// str = append(str, string(ques))
	// str = append(str, string(ques))

	// data, err := tcb.InsertDocuments("questions", str, "")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)

	//DeleteDocument *****************************************************************************************************************************
	// data, err := tcb.DeleteDocument("questions", "0a3f93955ff8252200ccae8f3bb567dc", "")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)
	
	// 函数执行
	//var g Inv
	//g.Data.Id = "990e1249-5388-4263-ba20-07e98f21610e"
	//g.Data = aaa
	//ques, err := json.Marshal(g)
	//if err != nil {
	//	fmt.Printf("json.Marshal%v", err)
	//}
	//fmt.Println(string(ques))
	//data, err := tcb.Invoke("getmationinfo", ques)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(data)

}
