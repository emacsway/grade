package aggregate

type AggregateVersioner interface {
	Version() uint
	SetVersion(uint)
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

func (a *VersionedAggregate) NextVersion() uint {
	a.version += 1
	return a.version
}

func (a *VersionedAggregate) SetVersion(val uint) {
	a.version = val
}

func (a VersionedAggregate) Export(ex VersionedAggregateExporterSetter) {
	ex.SetVersion(a.Version())
}

type VersionedAggregateExporterSetter interface {
	SetVersion(uint)
}
