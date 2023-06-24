package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

func NewMember(
	id values.TenantMemberId,
	status values.Status,
	fullName values.FullName,
	createdAt time.Time,
) *Member {
	return &Member{
		id:        id,
		status:    status,
		fullName:  fullName,
		createdAt: createdAt,
	}
}

type Member struct {
	id           values.TenantMemberId
	status       values.Status
	fullName     values.FullName
	createdAt    time.Time
	eventSourced aggregate.EventSourcedAggregate
}

func (m Member) Id() values.TenantMemberId {
	return m.id
}

func (m Member) PendingDomainEvents() []aggregate.DomainEvent {
	return m.eventSourced.PendingDomainEvents()
}

func (m *Member) ClearPendingDomainEvents() {
	m.eventSourced.ClearPendingDomainEvents()
}

func (m Member) Version() uint {
	return m.eventSourced.Version()
}

func (m *Member) SetVersion(val uint) {
	m.eventSourced.SetVersion(val)
}

func (m Member) Export(ex MemberExporterSetter) {
	ex.SetId(m.id)
	ex.SetStatus(m.status)
	ex.SetFullName(m.fullName)
	ex.SetVersion(m.Version())
	ex.SetCreatedAt(m.createdAt)
}

type MemberExporterSetter interface {
	SetId(id values.TenantMemberId)
	SetStatus(values.Status)
	SetFullName(values.FullName)
	SetVersion(uint)
	SetCreatedAt(time.Time)
}
