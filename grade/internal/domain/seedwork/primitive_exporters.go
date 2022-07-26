package seedwork

import "database/sql/driver"

func NewUint8Exporter(value uint8) *Uint8Exporter {
	r := Uint8Exporter(value)
	return &r
}

type Uint8Exporter uint8

func (e *Uint8Exporter) SetState(value uint8) {
	*e = Uint8Exporter(value)
}
func (e Uint8Exporter) Value() (driver.Value, error) {
	return uint8(e), nil
}

func NewUintExporter(value uint) *UintExporter {
	r := UintExporter(value)
	return &r
}

type UintExporter uint

func (e *UintExporter) SetState(value uint) {
	*e = UintExporter(value)
}
func (e UintExporter) Value() (driver.Value, error) {
	return uint(e), nil
}

func NewUint64Exporter(value uint64) *Uint64Exporter {
	r := Uint64Exporter(value)
	return &r
}

type Uint64Exporter uint64

func (e *Uint64Exporter) SetState(value uint64) {
	*e = Uint64Exporter(value)
}
func (e Uint64Exporter) Value() (driver.Value, error) {
	return uint64(e), nil
}
