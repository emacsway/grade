package repository

func NewStreamId(
	tenantId int,
	streamType string,
	streamId string,
) (StreamId, error) {
	return StreamId{
		tenantId:   tenantId,
		streamType: streamType,
		streamId:   streamId,
	}, nil
}

type StreamId struct {
	tenantId   int
	streamType string
	streamId   string
}

func (id StreamId) TenantId() int {
	return id.tenantId
}
func (id StreamId) StreamType() string {
	return id.streamType
}
func (id StreamId) StreamId() string {
	return id.streamId
}
