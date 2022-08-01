package seedwork

type AggregateVersionable interface {
	Version() uint
	IncreaseVersion()
}

func NewVersionedAggregate(version uint) VersionedAggregate {
	return VersionedAggregate{version: version}
}

type VersionedAggregate struct {
	version uint
}

func (a VersionedAggregate) Version() uint {
	return a.version
}

func (a *VersionedAggregate) IncreaseVersion() {
	a.version += 1
}
