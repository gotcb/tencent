package tencent

import (
	"github.com/gotcb/tencent/context"
	"github.com/gotcb/tencent/tcb"
)

//Config 配置
type Config struct {
	Timestamp int64  //当前时间戳
	SecretId  string //SecretId
	SecretKey string //secret
	Envid     string //环境id
}

//Tencent struct
type Tencent struct {
	Context *context.Context
}

// NewTencent init
func NewTencent(cfg *Config) *Tencent {
	context := new(context.Context)
	context.Timestamp = cfg.Timestamp
	context.SecretId = cfg.SecretId
	context.Envid = cfg.Envid
	context.SecretKey = cfg.SecretKey
	return &Tencent{context}
}

//GetAuth 初始化auth
func (tc *Tencent) GetAuth() error {
	return tc.Context.GetAuth()
}

//GetTcb 获取auth
func (tc *Tencent) GetTcb() *tcb.Tcb {
	return tcb.NewTcb(tc.Context)
}
