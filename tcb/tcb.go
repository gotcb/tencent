package tcb

import (
	"encoding/json"
	"strings"

	"github.com/gotcb/tencent/context"
	"github.com/gotcb/tencent/util"
)

const (
	//触发云函数
	invokeCloudFunctionURL = "https://tcb-api.tencentcloudapi.com/api/v2/envs/"
)

//Tcb tcb结构体
type Tcb struct {
	Context *context.Context
}

//Body 请求body
type Body struct {
	Query string `json:"query,omitempty"`
	Data  string `json:"data,omitempty"`
}

// type Inv struct {
// 	Data Id `json:"data,omitempty"`
// }
// type Id struct {
// 	Id string `json:"id,omitempty"`
// }

// 批量插入文档
type Data struct {
	InsertDatas []string `json:"data,omitempty"`
}

//NewTcb 创建tcb
func NewTcb(ct *context.Context) *Tcb {
	return &Tcb{
		Context: ct,
	}
}

//GetDocument 云函数调用
func (tcb *Tcb) GetDocument(collectionName, docID, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents/")
	build.WriteString(docID)
	build.WriteString("?transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	util := util.NewUtil(tcb.Context)

	data, err := util.SendPostURL("GET", urls, nil)

	if err != nil {
		return "", err
	}
	return data, nil
}

//UpdateDocument 云函数调用 单文档更新
func (tcb *Tcb) UpdateDocument(collectionName, docID, data, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents/")
	build.WriteString(docID)
	build.WriteString("?transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	util := util.NewUtil(tcb.Context)

	var by Body
	by.Data = data

	ques, err := json.Marshal(by)
	if err != nil {
		return "", err
	}
	datast, err := util.SendPostURL("PATCH", urls, ques)

	if err != nil {
		return "", err
	}
	return datast, nil
}

//SetDocument 单文档替换更新，若指定 docId 的 doc 已存在，则替换更新，否则插入新 doc
func (tcb *Tcb) SetDocument(collectionName, docID, data, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents/")
	build.WriteString(docID)
	build.WriteString("?transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	util := util.NewUtil(tcb.Context)

	var by Body
	by.Data = data

	ques, err := json.Marshal(by)
	if err != nil {
		return "", err
	}
	datast, err := util.SendPostURL("PUT", urls, ques)

	if err != nil {
		return "", err
	}
	return datast, nil
}

//InsertDocument 单文档插入
func (tcb *Tcb) InsertDocument(collectionName, docID, data, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents/")
	build.WriteString(docID)
	build.WriteString("?transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	util := util.NewUtil(tcb.Context)

	var by Body
	by.Data = data

	ques, err := json.Marshal(by)
	if err != nil {
		return "", err
	}
	datast, err := util.SendPostURL("POST", urls, ques)

	if err != nil {
		return "", err
	}
	return datast, nil
}

//DeleteDocument 单文档删除
func (tcb *Tcb) DeleteDocument(collectionName, docID, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents/")
	build.WriteString(docID)
	build.WriteString("?transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	util := util.NewUtil(tcb.Context)

	datast, err := util.SendPostURL("DELETE", urls, nil)

	if err != nil {
		return "", err
	}
	return datast, nil
}

//InsertDocuments 批量插入文档
func (tcb *Tcb) InsertDocuments(collectionName string, data []string, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents")
	build.WriteString("?transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	util := util.NewUtil(tcb.Context)

	var by Data
	by.InsertDatas = data

	ques, err := json.Marshal(by)
	if err != nil {
		return "", err
	}
	datast, err := util.SendPostURL("POST", urls, ques)

	if err != nil {
		return "", err
	}
	return datast, nil
}

//Find 查询数据库
func (tcb *Tcb) Find(collectionName, query, limit, skip, fields, sort, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents:find?limit=")
	build.WriteString(limit)
	build.WriteString("&skip=")
	build.WriteString(skip)

	build.WriteString("&fields=")
	build.WriteString(fields)

	build.WriteString("&sort=")
	build.WriteString(sort)
	build.WriteString("&transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	var by Body
	by.Query = query
	ques, err := json.Marshal(by)
	if err != nil {
		return "", err
	}
	util := util.NewUtil(tcb.Context)
	data, err := util.SendPostURL("POST", urls, ques)
	if err != nil {
		return "", err
	}
	return data, nil
}

//Count 查询数据库数量
func (tcb *Tcb) Count(collectionName, query, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents:count?transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	var by Body
	by.Query = query

	ques, err := json.Marshal(by)
	if err != nil {
		return "", err
	}
	util := util.NewUtil(tcb.Context)
	data, err := util.SendPostURL("POST", urls, ques)
	if err != nil {
		return "", err
	}
	return data, nil
}

//UpdateOne 批量查询并更新单文档
func (tcb *Tcb) UpdateOne(collectionName, query, data, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents:updateOne?transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	var by Body
	by.Query = query
	by.Data = data

	ques, err := json.Marshal(by)
	if err != nil {
		return "", err
	}
	util := util.NewUtil(tcb.Context)
	datast, err := util.SendPostURL("POST", urls, ques)
	if err != nil {
		return "", err
	}
	return datast, nil
}

//UpdateMany 批量查询并批量更新
func (tcb *Tcb) UpdateMany(collectionName, query, data, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents:updateMany?transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	var by Body
	by.Query = query
	by.Data = data

	ques, err := json.Marshal(by)
	if err != nil {
		return "", err
	}
	util := util.NewUtil(tcb.Context)
	datast, err := util.SendPostURL("POST", urls, ques)
	if err != nil {
		return "", err
	}
	return datast, nil
}

//DeleteOne 批量查询并删除单文档
func (tcb *Tcb) DeleteOne(collectionName, query, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents:deleteOne?transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	var by Body
	by.Query = query

	ques, err := json.Marshal(by)
	if err != nil {
		return "", err
	}
	util := util.NewUtil(tcb.Context)
	datast, err := util.SendPostURL("POST", urls, ques)
	if err != nil {
		return "", err
	}
	return datast, nil
}

//DeleteMany 批量查询并批量删除
func (tcb *Tcb) DeleteMany(collectionName, query, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/documents:deleteMany?transactionId=")
	build.WriteString(transactionID)
	urls := build.String()

	var by Body
	by.Query = query

	ques, err := json.Marshal(by)
	if err != nil {
		return "", err
	}
	util := util.NewUtil(tcb.Context)
	datast, err := util.SendPostURL("POST", urls, ques)
	if err != nil {
		return "", err
	}
	return datast, nil
}

//StartTransaction 开始事务
func (tcb *Tcb) StartTransaction(collectionName string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/transaction")

	urls := build.String()

	util := util.NewUtil(tcb.Context)
	datast, err := util.SendPostURL("POST", urls, nil)
	if err != nil {
		return "", err
	}
	return datast, nil
}

//CommitTransaction 提交事务
func (tcb *Tcb) CommitTransaction(collectionName, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/transaction/")
	build.WriteString(transactionID)
	build.WriteString(":commit")

	urls := build.String()

	util := util.NewUtil(tcb.Context)
	datast, err := util.SendPostURL("POST", urls, nil)
	if err != nil {
		return "", err
	}
	return datast, nil
}

//RollbackTransaction 回滚事务
func (tcb *Tcb) RollbackTransaction(collectionName, transactionID string) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/databases/")
	build.WriteString(collectionName)
	build.WriteString("/transaction/")
	build.WriteString(transactionID)
	build.WriteString(":rollback")
	urls := build.String()

	util := util.NewUtil(tcb.Context)
	datast, err := util.SendPostURL("POST", urls, nil)
	if err != nil {
		return "", err
	}
	return datast, nil
}

//函数执行
func (tcb *Tcb) Invoke(functionName string, data []byte) (string, error) {

	var build strings.Builder
	build.WriteString(invokeCloudFunctionURL)
	build.WriteString(tcb.Context.Envid)
	build.WriteString("/functions/")
	build.WriteString(functionName)
	build.WriteString(":invoke")
	urls := build.String()

	// ques, err := json.Marshal(data)
	// if err != nil {
	// 	return "", err
	// }
	util := util.NewUtil(tcb.Context)
	datast, err := util.SendPostURL("POST", urls, data)
	if err != nil {
		return "", err
	}
	return datast, nil
}
