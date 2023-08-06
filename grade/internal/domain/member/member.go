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
) (*Member, error) {
	return &Member{
		id:        id,
		status:    status,
		fullName:  fullName,
		createdAt: createdAt,
	}, nil
}

type Member struct {
	id        values.TenantMemberId
	status    values.Status
	fullName  values.FullName
	createdAt time.Time
	eventive  aggregate.EventiveEntity[aggregate.DomainEvent]
	aggregate.VersionedAggregate
}

func (m Member) Id() values.TenantMemberId {
	return m.id
}

func (m Member) PendingDomainEvents() []aggregate.DomainEvent {
	return m.eventive.PendingDomainEvents()
}

func (m *Member) ClearPendingDomainEvents() {
	m.eventive.ClearPendingDomainEvents()
}

func (m Member) Export(ex MemberExporterSetter) {
	ex.SetId(m.id)
	ex.SetStatus(m.status)
	ex.SetFullName(m.fullName)
	ex.SetCreatedAt(m.createdAt)
	ex.SetVersion(m.Version())
}

type MemberExporterSetter interface {
	SetId(id values.TenantMemberId)
	SetStatus(values.Status)
	SetFullName(values.FullName)
	SetCreatedAt(time.Time)
	SetVersion(uint)
}
