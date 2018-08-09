package cqrs

type EventStore interface {
	Persist(string,[]EventWrapper)
	Load(string) []EventWrapper
}

type MemEventStore struct {
	eventMap map[string][]EventWrapper
}

func (s *MemEventStore) Persist(aggregateId string, newEvents []EventWrapper) {
	events := s.eventMap[aggregateId]
	if events == nil{
		events = make([]EventWrapper,0)
	}
	events = append(events, newEvents...)
	s.eventMap[aggregateId] = events
}
func (s *MemEventStore) Load(aggregateId string) []EventWrapper {
	return s.eventMap[aggregateId]
}

func NewMemEventStore() EventStore {
	return &MemEventStore{make(map[string][]EventWrapper)}
}