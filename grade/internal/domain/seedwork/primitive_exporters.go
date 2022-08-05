package seedwork

import (
	"database/sql/driver"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

type Uint8Exporter uint8

func (ex *Uint8Exporter) SetState(value uint8) {
	*ex = Uint8Exporter(value)
}
func (ex Uint8Exporter) Value() (driver.Value, error) {
	return uint8(ex), nil
}

type UintExporter uint

func (ex *UintExporter) SetState(value uint) {
	*ex = UintExporter(value)
}
func (ex UintExporter) Value() (driver.Value, error) {
	return uint(ex), nil
}

type Uint64Exporter uint64

func (ex *Uint64Exporter) SetState(value uint64) {
	*ex = Uint64Exporter(value)
}
func (ex Uint64Exporter) Value() (driver.Value, error) {
	return uint64(ex), nil
}

type UuidExporter uuid.Uuid

func (ex *UuidExporter) SetState(value uuid.Uuid) {
	*ex = UuidExporter(value)
}
func (ex UuidExporter) Value() (driver.Value, error) {
	return uuid.Uuid(ex), nil
}

type StringExporter string

func (ex *StringExporter) SetState(value string) {
	*ex = StringExporter(value)
}
func (ex StringExporter) Value() (driver.Value, error) {
	return string(ex), nil
}
