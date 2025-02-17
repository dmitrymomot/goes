package aggregate

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/modernice/goes/event"
)

const (
	// ID means there is an inconsistency in the aggregate ids.
	InconsistentID = ConsistencyKind(iota + 1)

	// Name means there is an inconsistency in the aggregate names.
	InconsistentName

	// Version means there is an inconsistency in the event versions.
	InconsistentVersion

	// Time means there is an inconsistency in the event times.
	InconsistentTime
)

// Error is a consistency error.
type ConsistencyError struct {
	// Kind is the kind of incosistency.
	Kind ConsistencyKind
	// Aggregate is the handled aggregate.
	Aggregate Ref
	// CurrentVersion is the current version of the aggregate (without the tested changes).
	CurrentVersion int
	// Events are the tested events.
	Events []event.Event
	// EventIndex is the index of the event that caused the Error.
	EventIndex int
}

// ConsistencyKind is the kind of inconsistency.
type ConsistencyKind int

// IsConsistencyError returns whether an error was caused by an inconsistency in
// the events of an aggregate. An error is considered a consistency error if it
// either unwraps to a *ConsistencyError or if it has an IsConsistencyError() bool
// method that return true for the given error.
func IsConsistencyError(err error) bool {
	var cerr *ConsistencyError
	if errors.As(err, &cerr) {
		return true
	}

	var ierr interface{ IsConsistencyError() bool }
	if errors.As(err, &ierr) {
		return ierr.IsConsistencyError()
	}

	return false
}

// ConsistencyOption is an option for consistency validation.
type ConsistencyOption func(*consistencyValidation)

// IgnoreTime returns a ConsistencyOption that disables validation of event times.
func IgnoreTime(ignore bool) ConsistencyOption {
	return func(cfg *consistencyValidation) {
		cfg.ignoreTime = ignore
	}
}

type consistencyValidation struct {
	ignoreTime bool
}

// Validate tests the consistency of aggregate changes (events).
//
// The provided events are valid if they are correctly sorted by both version
// and time. No two events may have the same version or time, and their versions
// must be greater than 0.
func ValidateConsistency[Data any, Events ~[]event.Of[Data]](ref Ref, currentVersion int, events Events, opts ...ConsistencyOption) error {
	var cfg consistencyValidation
	for _, opt := range opts {
		opt(&cfg)
	}

	aevents := make([]event.Event, len(events))
	for i, evt := range events {
		aevents[i] = event.Any(evt)
	}

	var hasPrevEvent bool
	var prevEvent event.Event
	var prevVersion int

	for i, evt := range aevents {
		eid, ename, ev := evt.Aggregate()
		if eid != ref.ID {
			return &ConsistencyError{
				Kind:           InconsistentID,
				Aggregate:      ref,
				CurrentVersion: currentVersion,
				Events:         aevents,
				EventIndex:     i,
			}
		}
		if ename != ref.Name {
			return &ConsistencyError{
				Kind:           InconsistentName,
				Aggregate:      ref,
				CurrentVersion: currentVersion,
				Events:         aevents,
				EventIndex:     i,
			}
		}
		if ev <= 0 {
			return &ConsistencyError{
				Kind:           InconsistentVersion,
				Aggregate:      ref,
				CurrentVersion: currentVersion,
				Events:         aevents,
				EventIndex:     i,
			}
		}
		if ev <= currentVersion {
			return &ConsistencyError{
				Kind:           InconsistentVersion,
				Aggregate:      ref,
				CurrentVersion: currentVersion,
				Events:         aevents,
				EventIndex:     i,
			}
		}
		if hasPrevEvent && ev <= prevVersion {
			return &ConsistencyError{
				Kind:           InconsistentVersion,
				Aggregate:      ref,
				CurrentVersion: currentVersion,
				Events:         aevents,
				EventIndex:     i,
			}
		}
		if hasPrevEvent && !cfg.ignoreTime {
			nano := evt.Time().UnixNano()
			prevNano := prevEvent.Time().UnixNano()
			if nano <= prevNano {
				return &ConsistencyError{
					Kind:           InconsistentTime,
					Aggregate:      ref,
					CurrentVersion: currentVersion,
					Events:         aevents,
					EventIndex:     i,
				}
			}
		}
		prevEvent = evt
		prevVersion = ev
		hasPrevEvent = true
	}
	return nil
}

// Event return the first event that caused an inconsistency.
func (err *ConsistencyError) Event() event.Event {
	if err.EventIndex < 0 || err.EventIndex >= len(err.Events) {
		return nil
	}
	return err.Events[err.EventIndex]
}

// Error returns a string representation of the *ConsistencyError. The returned
// string describes the inconsistency error in detail, including the invalid
// AggregateID, AggregateName, AggregateVersion or Time. The method is used to
// provide an error message for a *ConsistencyError.
func (err *ConsistencyError) Error() string {
	evt := err.Event()
	var (
		id   uuid.UUID
		name string
		v    int
	)
	if evt != nil {
		id, name, v = evt.Aggregate()
	}

	var (
		aid   uuid.UUID
		aname string
	)

	aid, aname, _ = err.Aggregate.Aggregate()

	switch err.Kind {
	case InconsistentID:
		return fmt.Sprintf(
			"consistency: %q event has invalid AggregateID. want=%s got=%s",
			evt.Name(), aid, id,
		)
	case InconsistentName:
		return fmt.Sprintf(
			"consistency: %q event has invalid AggregateName. want=%s got=%s",
			evt.Name(), aname, name,
		)
	case InconsistentVersion:
		return fmt.Sprintf(
			"consistency: %q event has invalid AggregateVersion. want >=%d got=%d",
			evt.Name(), err.CurrentVersion+1, v,
		)
	case InconsistentTime:
		return fmt.Sprintf(
			"consistency: %q event has invalid Time. want=after %v got=%v",
			evt.Name(), err.Events[err.EventIndex-1].Time(), evt.Time(),
		)
	default:
		return fmt.Sprintf("consistency: invalid inconsistency kind=%d", err.Kind)
	}
}

// IsConsistencyError implements error.Is.
func (err *ConsistencyError) IsConsistencyError() bool {
	return true
}

// String returns a string representation of the ConsistencyKind
// [ConsistencyKind]. It returns one of the following strings:
// "<InconsistentID>", "<InconsistentName>", "<InconsistentVersion>",
// "<InconsistentTime>", or "<UnknownInconsistency>".
func (k ConsistencyKind) String() string {
	switch k {
	case InconsistentID:
		return "<InconsistentID>"
	case InconsistentName:
		return "<InconsistentName>"
	case InconsistentVersion:
		return "<InconsistentVersion>"
	case InconsistentTime:
		return "<InconsistentTime>"
	default:
		return "<UnknownInconsistency>"
	}
}

// func currentVersion(a Aggregate) int {
// 	_, _, v := a.Aggregate()
// 	return v + len(a.AggregateChanges())
// }
