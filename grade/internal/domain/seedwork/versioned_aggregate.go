package seedwork

type AggregateVersionable interface {
	GetVersion() uint
	IncreaseVersion()
}

func NewVersionedAggregate(version uint) (VersionedAggregate, error) {
	return VersionedAggregate{version: version}, nil
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
