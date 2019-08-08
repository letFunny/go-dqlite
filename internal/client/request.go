package client

// DO NOT EDIT
//
// This file was generated by ./schema.sh

import (
	"github.com/canonical/go-dqlite/internal/bindings"
)

// EncodeLeader encodes a Leader request.
func EncodeLeader(request *Message) {
	request.putUint64(0)

	request.putHeader(bindings.RequestLeader)
}

// EncodeClient encodes a Client request.
func EncodeClient(request *Message, id uint64) {
	request.putUint64(id)

	request.putHeader(bindings.RequestClient)
}

// EncodeHeartbeat encodes a Heartbeat request.
func EncodeHeartbeat(request *Message, timestamp uint64) {
	request.putUint64(timestamp)

	request.putHeader(bindings.RequestHeartbeat)
}

// EncodeOpen encodes a Open request.
func EncodeOpen(request *Message, name string, flags uint64, vfs string) {
	request.putString(name)
	request.putUint64(flags)
	request.putString(vfs)

	request.putHeader(bindings.RequestOpen)
}

// EncodePrepare encodes a Prepare request.
func EncodePrepare(request *Message, db uint64, sql string) {
	request.putUint64(db)
	request.putString(sql)

	request.putHeader(bindings.RequestPrepare)
}

// EncodeExec encodes a Exec request.
func EncodeExec(request *Message, db uint32, stmt uint32, values NamedValues) {
	request.putUint32(db)
	request.putUint32(stmt)
	request.putNamedValues(values)

	request.putHeader(bindings.RequestExec)
}

// EncodeQuery encodes a Query request.
func EncodeQuery(request *Message, db uint32, stmt uint32, values NamedValues) {
	request.putUint32(db)
	request.putUint32(stmt)
	request.putNamedValues(values)

	request.putHeader(bindings.RequestQuery)
}

// EncodeFinalize encodes a Finalize request.
func EncodeFinalize(request *Message, db uint32, stmt uint32) {
	request.putUint32(db)
	request.putUint32(stmt)

	request.putHeader(bindings.RequestFinalize)
}

// EncodeExecSQL encodes a ExecSQL request.
func EncodeExecSQL(request *Message, db uint64, sql string, values NamedValues) {
	request.putUint64(db)
	request.putString(sql)
	request.putNamedValues(values)

	request.putHeader(bindings.RequestExecSQL)
}

// EncodeQuerySQL encodes a QuerySQL request.
func EncodeQuerySQL(request *Message, db uint64, sql string, values NamedValues) {
	request.putUint64(db)
	request.putString(sql)
	request.putNamedValues(values)

	request.putHeader(bindings.RequestQuerySQL)
}

// EncodeInterrupt encodes a Interrupt request.
func EncodeInterrupt(request *Message, db uint64) {
	request.putUint64(db)

	request.putHeader(bindings.RequestInterrupt)
}

// EncodeJoin encodes a Join request.
func EncodeJoin(request *Message, id uint64, address string) {
	request.putUint64(id)
	request.putString(address)

	request.putHeader(bindings.RequestJoin)
}

// EncodePromote encodes a Promote request.
func EncodePromote(request *Message, id uint64) {
	request.putUint64(id)

	request.putHeader(bindings.RequestPromote)
}

// EncodeRemove encodes a Remove request.
func EncodeRemove(request *Message, id uint64) {
	request.putUint64(id)

	request.putHeader(bindings.RequestRemove)
}
