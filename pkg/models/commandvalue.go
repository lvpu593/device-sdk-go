package models

type ValueType int

type CommandValue struct {
	DeviceResourceName string

	Origin int64

	Type ValueType

	NumericValue []byte

	stringValue string

	BinValue []byte
}
