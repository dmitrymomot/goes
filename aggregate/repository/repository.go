package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/modernice/goes/aggregate"
	"github.com/modernice/goes/aggregate/query"
	"github.com/modernice/goes/aggregate/snapshot"
	"github.com/modernice/goes/aggregate/stream"
	"github.com/modernice/goes/event"
	equery "github.com/modernice/goes/event/query"
	"github.com/modernice/goes/event/query/version"
	"github.com/modernice/goes/helper/streams"
)

var (
	// ErrVersionNotFound is returned when trying to fetch an aggregate with a
	// version higher than the current version of the aggregate.
	ErrVersionNotFound = errors.New("version not found")

	// ErrDeleted is returned when trying to fetch an aggregate that has been soft-deleted.
	ErrDeleted = errors.New("aggregate was soft-deleted")
)

// Option is a repository option.
type Option func(*Repository)

// Repository provides an event-sourced aggregate repository for persisting and
// querying aggregates. It uses an event.Store to persist and query aggregates.
// Repository supports snapshots, hooks for inserting events, and query
// modifiers. It also supports deleting an aggregate by deleting its events from
// the event store. Use the Query method to query the event store for events
// that match a given query and use the returned Histories to build the current
// state of the queried aggregates.
type Repository struct {
	store          event.Store
	snapshots      snapshot.Store
	snapSchedule   snapshot.Schedule
	queryModifiers []func(context.Context, aggregate.Query, event.Query) (event.Query, error)
	beforeInsert   []func(context.Context, aggregate.Aggregate) error
	afterInsert    []func(context.Context, aggregate.Aggregate) error
	onFailedInsert []func(context.Context, aggregate.Aggregate, error) error
	onDelete       []func(context.Context, aggregate.Aggregate) error

	validateConsistency bool
}

// WithSnapshots returns an Option that add a Snapshot Store to a Repository.
//
// A Repository that has a Snapshot Store will fetch the latest valid Snapshot
// for an aggregate before fetching the necessary events to reconstruct the
// state of the Agrgegate.
//
// An optional Snapshot Schedule can be provided to instruct the Repository to
// make and save Snapshots into the Snapshot Store when appropriate:
//
//	var store snapshot.Store
//	r := repository.New(store, snapshot.Every(3))
//
// The example above will make a Snapshot of an aggregate every third version of
// the aggregate.
//
// Aggregates must implement snapshot.Marshaler & snapshot.Unmarshaler in order
// for Snapshots to work.
func WithSnapshots(store snapshot.Store, s snapshot.Schedule) Option {
	if store == nil {
		panic("nil Store")
	}
	return func(r *Repository) {
		r.snapshots = store
		r.snapSchedule = s
	}
}

// ValidateConsistency returns an Option that configures a [Repository] to
// validate the consistency of aggregate events before inserting the events into
// the event store. Defaults to true.
func ValidateConsistency(validate bool) Option {
	return func(r *Repository) {
		r.validateConsistency = validate
	}
}

// ModifyQueries returns an Option that adds mods as Query modifiers to a
// Repository. When the Repository builds a Query, it is passed to every
// modifier before the event store is queried.
func ModifyQueries(mods ...func(ctx context.Context, q aggregate.Query, prev event.Query) (event.Query, error)) Option {
	return func(r *Repository) {
		r.queryModifiers = append(r.queryModifiers, mods...)
	}
}

// BeforeInsert returns an Option that adds fn as a hook to a Repository. fn is
// called before the changes to an aggregate are inserted into the event store.
func BeforeInsert(fn func(context.Context, aggregate.Aggregate) error) Option {
	return func(r *Repository) {
		r.beforeInsert = append(r.beforeInsert, fn)
	}
}

// AfterInsert returns an Option that adds fn as a hook to a Repository. fn is
// called after the changes to an aggregate are inserted into the event store.
func AfterInsert(fn func(context.Context, aggregate.Aggregate) error) Option {
	return func(r *Repository) {
		r.afterInsert = append(r.afterInsert, fn)
	}
}

// OnFailedInsert returns an Option that adds fn as a hook to a Repository. fn
// is called when the Repository fails to insert the changes to an aggregate
// into the event store.
func OnFailedInsert(fn func(context.Context, aggregate.Aggregate, error) error) Option {
	return func(r *Repository) {
		r.onFailedInsert = append(r.onFailedInsert, fn)
	}
}

// OnDelete returns an Option that adds fn as a hook to a Repository. fn is
// called after an aggregate has been deleted.
func OnDelete(fn func(context.Context, aggregate.Aggregate) error) Option {
	return func(r *Repository) {
		r.onDelete = append(r.onDelete, fn)
	}
}

// New returns an event-sourced aggregate Repository. It uses the provided event
// Store to persist and query aggregates.
func New(store event.Store, opts ...Option) *Repository {
	return newRepository(store, opts...)
}

func newRepository(store event.Store, opts ...Option) *Repository {
	r := &Repository{
		store:               store,
		validateConsistency: true,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// Save saves the changes to an aggregate into the underlying event store and
// flushes its changes afterwards (by calling a.FlushChanges).
func (r *Repository) Save(ctx context.Context, a aggregate.Aggregate) error {
	if r.validateConsistency {
		id, name, version := a.Aggregate()
		ref := aggregate.Ref{Name: name, ID: id}
		if err := aggregate.ValidateConsistency(ref, version, a.AggregateChanges()); err != nil {
			return fmt.Errorf("validate consistency: %w", err)
		}
	}

	var snap bool
	if r.snapSchedule != nil && r.snapSchedule.Test(a) {
		snap = true
	}

	for _, fn := range r.beforeInsert {
		if err := fn(ctx, a); err != nil {
			return fmt.Errorf("BeforeInsert: %w", err)
		}
	}

	if err := r.store.Insert(ctx, a.AggregateChanges()...); err != nil {
		for _, fn := range r.onFailedInsert {
			if hookError := fn(ctx, a, err); hookError != nil {
				return fmt.Errorf("OnFailedInsert (%s): %w", err, hookError)
			}
		}

		return fmt.Errorf("insert events: %w", err)
	}

	for _, fn := range r.afterInsert {
		if err := fn(ctx, a); err != nil {
			return fmt.Errorf("AfterInsert: %w", err)
		}
	}

	if c, ok := a.(aggregate.Committer); ok {
		c.Commit()
	}

	if snap {
		if err := r.makeSnapshot(ctx, a); err != nil {
			return fmt.Errorf("make snapshot: %w", err)
		}
	}

	return nil
}

func (r *Repository) makeSnapshot(ctx context.Context, a aggregate.Aggregate) error {
	snap, err := snapshot.New(a)
	if err != nil {
		return err
	}
	if err = r.snapshots.Save(ctx, snap); err != nil {
		return fmt.Errorf("save snapshot: %w", err)
	}
	return nil
}

// Fetch fetches the events of the provided aggregate from the event store and
// applies them to it to build its current state.
//
// It is allowed to pass an aggregate that does't have any events in the event
// store yet.
//
// It is also allowed to pass an aggregate that has already events applied onto
// it. Only events with a version higher than the current version of the passed
// Aggregate are fetched from the event store.
func (r *Repository) Fetch(ctx context.Context, a aggregate.Aggregate) error {
	if _, ok := a.(snapshot.Target); ok && r.snapshots != nil {
		return r.fetchLatestWithSnapshot(ctx, a)
	}

	return r.fetch(ctx, a, equery.AggregateVersion(
		version.Min(aggregate.UncommittedVersion(a)+1),
	))
}

func (r *Repository) fetchLatestWithSnapshot(ctx context.Context, a aggregate.Aggregate) error {
	id, name, _ := a.Aggregate()

	snap, err := r.snapshots.Latest(ctx, name, id)
	if err != nil || snap == nil {
		return r.fetch(ctx, a, equery.AggregateVersion(
			version.Min(aggregate.UncommittedVersion(a)+1),
		))
	}

	if a, ok := a.(snapshot.Target); !ok {
		return fmt.Errorf("aggregate does not implement %T", a)
	} else {
		if err := snapshot.Unmarshal(snap, a); err != nil {
			return fmt.Errorf("unmarshal snapshot: %w", err)
		}
	}

	return r.fetch(ctx, a, equery.AggregateVersion(
		version.Min(aggregate.UncommittedVersion(a)+1),
	))
}

func (r *Repository) fetch(ctx context.Context, a aggregate.Aggregate, opts ...equery.Option) error {
	id, name, _ := a.Aggregate()

	opts = append([]equery.Option{
		equery.AggregateName(name),
		equery.AggregateID(id),
		equery.SortBy(event.SortAggregateVersion, event.SortAsc),
	}, opts...)

	events, err := r.queryEvents(ctx, equery.New(opts...))
	if err != nil {
		return fmt.Errorf("query events: %w", err)
	}

	if err = aggregate.ApplyHistory(a, events); err != nil {
		return fmt.Errorf("apply history: %w", err)
	}

	return nil
}

func (r *Repository) queryEvents(ctx context.Context, q equery.Query) ([]event.Event, error) {
	str, errs, err := r.store.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("query events: %w", err)
	}

	out := make([]event.Event, 0, len(str))
	var softDeleted bool
	if err := streams.Walk(ctx, func(evt event.Event) error {
		data := evt.Data()

		if data, ok := data.(aggregate.SoftDeleter); ok && data.SoftDelete() {
			softDeleted = true
		}

		if data, ok := data.(aggregate.SoftRestorer); ok && data.SoftRestore() {
			softDeleted = false
		}

		out = append(out, evt)

		return nil
	}, str, errs); err != nil {
		return out, err
	}

	if softDeleted {
		return out, ErrDeleted
	}

	return out, nil
}

// FetchVersion does the same as r.Fetch, but only fetches events up until the
// given version v. If the event store has no event for the provided aggregate
// with the requested version, ErrVersionNotFound is returned.
func (r *Repository) FetchVersion(ctx context.Context, a aggregate.Aggregate, v int) error {
	if v < 0 {
		v = 0
	}

	if r.snapshots != nil {
		return r.fetchVersionWithSnapshot(ctx, a, v)
	}

	return r.fetchVersion(ctx, a, v)
}

func (r *Repository) fetchVersionWithSnapshot(ctx context.Context, a aggregate.Aggregate, v int) error {
	id, name, _ := a.Aggregate()

	snap, err := r.snapshots.Limit(ctx, name, id, v)
	if err != nil || snap == nil {
		return r.fetchVersion(ctx, a, v)
	}

	if a, ok := a.(snapshot.Target); !ok {
		return fmt.Errorf("aggregate does not implement %T", a)
	} else {
		if err = snapshot.Unmarshal(snap, a); err != nil {
			return fmt.Errorf("unmarshal snapshot: %w", err)
		}
	}

	return r.fetchVersion(ctx, a, v)
}

func (r *Repository) fetchVersion(ctx context.Context, a aggregate.Aggregate, v int) error {
	if err := r.fetch(ctx, a, equery.AggregateVersion(
		version.Min(aggregate.UncommittedVersion(a)+1),
		version.Max(v),
	)); err != nil {
		return err
	}

	_, _, av := a.Aggregate()
	if av != v {
		return ErrVersionNotFound
	}

	return nil
}

// Delete deletes an aggregate by deleting its events from the event store.
func (r *Repository) Delete(ctx context.Context, a aggregate.Aggregate) error {
	id, name, _ := a.Aggregate()

	str, errs, err := r.store.Query(ctx, equery.New(
		equery.AggregateName(name),
		equery.AggregateID(id),
	))
	if err != nil {
		return fmt.Errorf("query events: %w", err)
	}

	for {
		if str == nil && errs == nil {
			break
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err, ok := <-errs:
			if !ok {
				errs = nil
				break
			}
			return fmt.Errorf("event stream: %w", err)
		case evt, ok := <-str:
			if !ok {
				str = nil
				break
			}
			if err = r.store.Delete(ctx, evt); err != nil {
				return fmt.Errorf("delete %q event (ID=%s): %w", evt.Name(), evt.ID(), err)
			}
		}
	}

	for _, fn := range r.onDelete {
		if err := fn(ctx, a); err != nil {
			return fmt.Errorf("OnDelete: %w", err)
		}
	}

	return nil
}

// Query queries the event store for events that match the given Query and
// returns a stream of aggregate Histories and errors. Use the returned
// Histories to build the current state of the queried aggregates:
//
//	var r *Repository
//	str, errs, err := r.Query(context.TODO(), query.New(...))
//	// handle err
//	histories, err := streams.Drain(context.TODO(), str, errs)
//	// handle err
//	for _, his := range histories {
//		aggregateName := his.AggregateName()
//		aggregateID := his.AggregateID()
//
//		// Create the aggregate from its name and UUID
//		foo := newFoo(aggregateID)
//
//		// Then apply its History
//		his.Apply(foo)
//	}
func (r *Repository) Query(ctx context.Context, q aggregate.Query) (<-chan aggregate.History, <-chan error, error) {
	eq, err := r.makeQuery(ctx, q)
	if err != nil {
		return nil, nil, fmt.Errorf("make query options: %w", err)
	}

	events, errs, err := r.store.Query(ctx, eq)
	if err != nil {
		return nil, nil, fmt.Errorf("query events: %w", err)
	}

	out, outErrors := stream.New(
		ctx,
		events,
		stream.Errors(errs),
		stream.Grouped(true),
		stream.Sorted(true),
	)

	return out, outErrors, nil
}

func (r *Repository) makeQuery(ctx context.Context, aq aggregate.Query) (event.Query, error) {
	opts := append(
		query.EventQueryOpts(aq),
		equery.SortByAggregate(),
	)

	var q event.Query = equery.New(opts...)
	var err error
	for _, mod := range r.queryModifiers {
		if q, err = mod(ctx, aq, q); err != nil {
			return q, fmt.Errorf("modify query: %w", err)
		}
	}

	return q, nil
}

// Use first fetches the aggregate a, then calls fn(a) and finally saves the
// aggregate. If the RetryUse() option is used, Use() is retried up to the
// configured maxTries option.
func (r *Repository) Use(ctx context.Context, a aggregate.Aggregate, fn func() error) error {
	var err error

	var trigger RetryTrigger
	var isRetryable IsRetryable

	if rp, ok := a.(Retryer); ok {
		trigger, isRetryable = rp.RetryUse()
	}

	for {
		if err != nil {
			if trigger == nil || isRetryable == nil || !isRetryable(err) {
				return err
			}

			if done := trigger.next(ctx); done != nil {
				return fmt.Errorf("%v: %w", done, err)
			}

			if discarder, ok := a.(ChangeDiscarder); ok {
				discarder.DiscardChanges()
			}
		}

		if err = r.Fetch(ctx, a); err != nil {
			err = fmt.Errorf("fetch aggregate: %w", err)
			continue
		}

		if err = fn(); err != nil {
			continue
		}

		if err = r.Save(ctx, a); err != nil {
			err = fmt.Errorf("save aggregate: %w", err)
			continue
		}

		return nil
	}
}
