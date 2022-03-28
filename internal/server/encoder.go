package server

import (
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

type responseFailed struct {
	Code int32  `json:"error_code"`
	Msg  string `json:"error_msg"`
}
type responseSuccess struct {
	Code int32       `json:"error_code"`
	Msg  string      `json:"error_msg"`
	Data interface{} `json:"data"`
}

func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	reply := &responseSuccess{
		Data: v,
	}
	codec := encoding.GetCodec("json")
	data, err := codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	reply := &responseFailed{
		Code: se.Code,
		Msg:  se.Message,
	}
	codec := encoding.GetCodec("json")
	data, err := codec.Marshal(reply)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	//1000以下为http code,设置header http status, 1000以上为业务code
	if se.Code < 1000 {
		w.WriteHeader(int(se.Code))
	}
	_, err = w.Write(data)
	if err != nil {
		return
	}
}
