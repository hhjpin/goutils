package response

import (
	"github.com/hhjpin/goutils/errors"
)

type BaseResponse struct {
	errors.Error
	Data interface{} `json:"data"`
}

func (b *BaseResponse) Init(errCode errors.ErrCode, d ...interface{}) {

	err := errors.New(errCode)
	b.ErrCode = errCode
	b.ErrMsg = err.ErrMsg
	b.ErrMsgEn = err.ErrMsgEn

	if d != nil {
		if len(d) > 1 {
			b.Data = d
		} else if len(d) == 1 {
			b.Data = d[0]
		}
	}
	if b.Data == nil {
		b.Data = map[string]string{}
	}
}

func (b *BaseResponse) InitError(err error, d ...interface{}) {
	if e, ok := err.(errors.Error); ok {
		b.initRespError(e, d...)
	} else {
		b.initRespError(errors.NewFormat(1, err.Error()), d...)
	}
}

func (b *BaseResponse) initRespError(err errors.Error, d ...interface{}) {

	b.ErrCode = err.ErrCode
	b.ErrMsg = err.ErrMsg
	b.ErrMsgEn = err.ErrMsgEn

	if d != nil {
		if len(d) > 1 {
			b.Data = d
		} else if len(d) == 1 {
			b.Data = d[0]
		}
	}
	if b.Data == nil {
		b.Data = map[string]string{}
	}
}
