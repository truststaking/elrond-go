package blockAPI

import (
	"errors"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go/process"
)

var (
	errNilArgAPIBlockProcessor   = errors.New("nil arg api block processor")
	errNilTransactionUnmarshaler = errors.New("nil transaction unmarshaler")
	errNilLogsRepository         = errors.New("nil logs repository")
)

func checkNilArg(arg *ArgAPIBlockProcessor) error {
	if arg == nil {
		return errNilArgAPIBlockProcessor
	}
	if check.IfNil(arg.Uint64ByteSliceConverter) {
		return process.ErrNilUint64Converter
	}
	if check.IfNil(arg.Store) {
		return process.ErrNilStorage
	}
	if check.IfNil(arg.Marshalizer) {
		return process.ErrNilMarshalizer
	}
	if check.IfNil(arg.HistoryRepo) {
		return process.ErrNilHistoryRepository
	}
	if check.IfNil(arg.TxUnmarshaller) {
		return errNilTransactionUnmarshaler
	}
	if check.IfNil(arg.Hasher) {
		return process.ErrNilHasher
	}
	if check.IfNil(arg.AddressPubkeyConverter) {
		return process.ErrNilPubkeyConverter
	}
	if check.IfNil(arg.LogsRepository) {
		return errNilLogsRepository
	}

	return nil
}
