package seedwork

type AggregateVersionable interface {
	GetVersion() uint
	IncreaseVersion()
}

func NewVersionedAggregate(version uint) VersionedAggregate {
	return VersionedAggregate{version: version}
}

type VersionedAggregate struct {
	version uint
}

func (a VersionedAggregate) GetVersion() uint {
	return a.version
}

func (a *VersionedAggregate) IncreaseVersion() {
	a.version += 1
}
