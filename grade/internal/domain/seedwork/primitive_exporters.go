package seedwork

import "database/sql/driver"

func NewUint8Exporter(value uint8) *Uint8Exporter {
	r := Uint8Exporter(value)
	return &r
}

type Uint8Exporter uint8

func (ex *Uint8Exporter) SetState(value uint8) {
	*ex = Uint8Exporter(value)
}
func (ex Uint8Exporter) Value() (driver.Value, error) {
	return uint8(ex), nil
}

func NewUintExporter(value uint) *UintExporter {
	r := UintExporter(value)
	return &r
}

type UintExporter uint

func (ex *UintExporter) SetState(value uint) {
	*ex = UintExporter(value)
}
func (ex UintExporter) Value() (driver.Value, error) {
	return uint(ex), nil
}

func NewUint64Exporter(value uint64) *Uint64Exporter {
	r := Uint64Exporter(value)
	return &r
}

type Uint64Exporter uint64

func (ex *Uint64Exporter) SetState(value uint64) {
	*ex = Uint64Exporter(value)
}
func (ex Uint64Exporter) Value() (driver.Value, error) {
	return uint64(ex), nil
}

func NewStringExporter(value string) *StringExporter {
	r := StringExporter(value)
	return &r
}

type StringExporter string

func (ex *StringExporter) SetState(value string) {
	*ex = StringExporter(value)
}
func (ex StringExporter) Value() (driver.Value, error) {
	return string(ex), nil
}
