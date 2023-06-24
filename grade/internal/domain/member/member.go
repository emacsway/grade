package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

func NewMember(
	id TenantMemberId,
	status Status,
	fullName FullName,
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
	id           TenantMemberId
	status       Status
	fullName     FullName
	createdAt    time.Time
	eventSourced aggregate.EventSourcedAggregate
}

func (m Member) Id() TenantMemberId {
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
	SetId(id TenantMemberId)
	SetStatus(Status)
	SetFullName(FullName)
	SetVersion(uint)
	SetCreatedAt(time.Time)
}
