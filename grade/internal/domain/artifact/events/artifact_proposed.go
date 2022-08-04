package events

type ArtifactProposed struct {
}

// EventType should be used instead of Invoke(Aggregate) approach
func (e ArtifactProposed) EventType() string {
	return "ArtifactProposed"
}

func (e ArtifactProposed) EventVersion() uint8 {
	return 1
}
