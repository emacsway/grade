package aggregate

func NewCausalDependency(
	aggregateId any,
	aggregateType string,
	aggregateVersion uint,
) CausalDependency {
	return CausalDependency{
		aggregateId:      aggregateId,
		aggregateType:    aggregateType,
		aggregateVersion: aggregateVersion,
	}
}

type CausalDependency struct {
	aggregateId      any
	aggregateType    string
	aggregateVersion uint
}

func (d CausalDependency) AggregateId() any {
	return d.aggregateId
}

func (d CausalDependency) AggregateType() string {
	return d.aggregateType
}

func (d CausalDependency) AggregateVersion() uint {
	return d.aggregateVersion
}

func (d CausalDependency) Export(ex CausalDependencyExporterSetter) {
	ex.SetAggregateId(d.aggregateId)
	ex.SetAggregateType(d.aggregateType)
	ex.SetAggregateVersion(d.aggregateVersion)
}

type CausalDependencyExporterSetter interface {
	SetAggregateId(any)
	SetAggregateType(string)
	SetAggregateVersion(uint)
}

type CausalDependencyExporter struct {
	AggregateId      any
	AggregateType    string
	AggregateVersion uint
}

func (ex *CausalDependencyExporter) SetAggregateId(val any) {
	ex.AggregateId = val
}
func (ex *CausalDependencyExporter) SetAggregateType(val string) {
	ex.AggregateType = val
}
func (ex *CausalDependencyExporter) SetAggregateVersion(val uint) {
	ex.AggregateVersion = val
}
