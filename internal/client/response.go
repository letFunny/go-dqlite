package client

// DO NOT EDIT
//
// This file was generated by ./schema.sh

import (
	"fmt"

	"github.com/canonical/go-dqlite/internal/bindings"
)

// DecodeFailure decodes a Failure response.
func DecodeFailure(response *Message) (code uint64, message string, err error) {
	mtype, _ := response.getHeader()

	if mtype == bindings.ResponseFailure {
		e := ErrRequest{}
		e.Code = response.getUint64()
		e.Description = response.getString()
                err = e
                return
	}

	if mtype != bindings.ResponseFailure {
		err = fmt.Errorf("unexpected response type %d", mtype)
                return
	}

	code = response.getUint64()
	message = response.getString()

	return
}

// DecodeWelcome decodes a Welcome response.
func DecodeWelcome(response *Message) (heartbeatTimeout uint64, err error) {
	mtype, _ := response.getHeader()

	if mtype == bindings.ResponseFailure {
		e := ErrRequest{}
		e.Code = response.getUint64()
		e.Description = response.getString()
                err = e
                return
	}

	if mtype != bindings.ResponseWelcome {
		err = fmt.Errorf("unexpected response type %d", mtype)
                return
	}

	heartbeatTimeout = response.getUint64()

	return
}

// DecodeServer decodes a Server response.
func DecodeServer(response *Message) (address string, err error) {
	mtype, _ := response.getHeader()

	if mtype == bindings.ResponseFailure {
		e := ErrRequest{}
		e.Code = response.getUint64()
		e.Description = response.getString()
                err = e
                return
	}

	if mtype != bindings.ResponseServer {
		err = fmt.Errorf("unexpected response type %d", mtype)
                return
	}

	address = response.getString()

	return
}

// DecodeServers decodes a Servers response.
func DecodeServers(response *Message) (servers Servers, err error) {
	mtype, _ := response.getHeader()

	if mtype == bindings.ResponseFailure {
		e := ErrRequest{}
		e.Code = response.getUint64()
		e.Description = response.getString()
                err = e
                return
	}

	if mtype != bindings.ResponseServers {
		err = fmt.Errorf("unexpected response type %d", mtype)
                return
	}

	servers = response.getServers()

	return
}

// DecodeDb decodes a Db response.
func DecodeDb(response *Message) (id uint32, err error) {
	mtype, _ := response.getHeader()

	if mtype == bindings.ResponseFailure {
		e := ErrRequest{}
		e.Code = response.getUint64()
		e.Description = response.getString()
                err = e
                return
	}

	if mtype != bindings.ResponseDb {
		err = fmt.Errorf("unexpected response type %d", mtype)
                return
	}

	id = response.getUint32()
	response.getUint32()

	return
}

// DecodeStmt decodes a Stmt response.
func DecodeStmt(response *Message) (db uint32, id uint32, params uint64, err error) {
	mtype, _ := response.getHeader()

	if mtype == bindings.ResponseFailure {
		e := ErrRequest{}
		e.Code = response.getUint64()
		e.Description = response.getString()
                err = e
                return
	}

	if mtype != bindings.ResponseStmt {
		err = fmt.Errorf("unexpected response type %d", mtype)
                return
	}

	db = response.getUint32()
	id = response.getUint32()
	params = response.getUint64()

	return
}

// DecodeEmpty decodes a Empty response.
func DecodeEmpty(response *Message) (err error) {
	mtype, _ := response.getHeader()

	if mtype == bindings.ResponseFailure {
		e := ErrRequest{}
		e.Code = response.getUint64()
		e.Description = response.getString()
                err = e
                return
	}

	if mtype != bindings.ResponseEmpty {
		err = fmt.Errorf("unexpected response type %d", mtype)
                return
	}

	response.getUint64()

	return
}

// DecodeResult decodes a Result response.
func DecodeResult(response *Message) (result Result, err error) {
	mtype, _ := response.getHeader()

	if mtype == bindings.ResponseFailure {
		e := ErrRequest{}
		e.Code = response.getUint64()
		e.Description = response.getString()
                err = e
                return
	}

	if mtype != bindings.ResponseResult {
		err = fmt.Errorf("unexpected response type %d", mtype)
                return
	}

	result = response.getResult()

	return
}

// DecodeRows decodes a Rows response.
func DecodeRows(response *Message) (rows Rows, err error) {
	mtype, _ := response.getHeader()

	if mtype == bindings.ResponseFailure {
		e := ErrRequest{}
		e.Code = response.getUint64()
		e.Description = response.getString()
                err = e
                return
	}

	if mtype != bindings.ResponseRows {
		err = fmt.Errorf("unexpected response type %d", mtype)
                return
	}

	rows = response.getRows()

	return
}
