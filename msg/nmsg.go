package msg

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	kindOffset     = 0  // 8 bytes
	priceOffset    = 8  // 8 bytes
	amountOffset   = 16 // 8 bytes
	stockIdOffset  = 24 // 8 bytes
	traderIdOffset = 32 // 4 bytes
	tradeIdOffset  = 36 // 4 bytes
	binarySize     = 40
)

var binCoder = binary.LittleEndian

// Populate NMessage with *Message values
func Marshal(b []byte, m *Message) error {
	if len(b) != binarySize {
		return errors.New(fmt.Sprintf("Wrong sized byte buffer. Expecting %d, found %d", binarySize, len(b)))
	}
	binCoder.PutUint64(b[kindOffset:priceOffset], uint64(m.Kind))
	binCoder.PutUint64(b[priceOffset:amountOffset], uint64(m.Price))
	binCoder.PutUint64(b[amountOffset:stockIdOffset], uint64(m.Amount))
	binCoder.PutUint64(b[stockIdOffset:traderIdOffset], uint64(m.StockId))
	binCoder.PutUint32(b[traderIdOffset:tradeIdOffset], uint32(m.TraderId))
	binCoder.PutUint32(b[tradeIdOffset:binarySize], uint32(m.TradeId))
	return nil
}

// Populate *Message with NMessage values
func Unmarshal(b []byte, m *Message) error {
	if len(b) != binarySize {
		return errors.New(fmt.Sprintf("Wrong sized byte buffer. Expecting %d, found %d", binarySize, len(b)))
	}
	m.Kind = MsgKind(binCoder.Uint64(b[kindOffset:priceOffset]))
	m.Price = binCoder.Uint64(b[priceOffset:amountOffset])
	m.Amount = binCoder.Uint64(b[amountOffset:stockIdOffset])
	m.StockId = binCoder.Uint64(b[stockIdOffset:traderIdOffset])
	m.TraderId = binCoder.Uint32(b[traderIdOffset:tradeIdOffset])
	m.TradeId = binCoder.Uint32(b[tradeIdOffset:binarySize])
	return nil
}
