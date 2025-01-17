// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AccessKey is an object representing the database table.
type AccessKey struct {
	Key        string    `boil:"key" json:"key" toml:"key" yaml:"key"`
	UID        int64     `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	ClientType string    `boil:"client_type" json:"client_type" toml:"client_type" yaml:"client_type"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *accessKeyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L accessKeyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AccessKeyColumns = struct {
	Key        string
	UID        string
	ClientType string
	CreatedAt  string
	UpdatedAt  string
}{
	Key:        "key",
	UID:        "uid",
	ClientType: "client_type",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

var AccessKeyTableColumns = struct {
	Key        string
	UID        string
	ClientType string
	CreatedAt  string
	UpdatedAt  string
}{
	Key:        "access_keys.key",
	UID:        "access_keys.uid",
	ClientType: "access_keys.client_type",
	CreatedAt:  "access_keys.created_at",
	UpdatedAt:  "access_keys.updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AccessKeyWhere = struct {
	Key        whereHelperstring
	UID        whereHelperint64
	ClientType whereHelperstring
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	Key:        whereHelperstring{field: "\"access_keys\".\"key\""},
	UID:        whereHelperint64{field: "\"access_keys\".\"uid\""},
	ClientType: whereHelperstring{field: "\"access_keys\".\"client_type\""},
	CreatedAt:  whereHelpertime_Time{field: "\"access_keys\".\"created_at\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"access_keys\".\"updated_at\""},
}

// AccessKeyRels is where relationship names are stored.
var AccessKeyRels = struct {
	UIDUser string
}{
	UIDUser: "UIDUser",
}

// accessKeyR is where relationships are stored.
type accessKeyR struct {
	UIDUser *User `boil:"UIDUser" json:"UIDUser" toml:"UIDUser" yaml:"UIDUser"`
}

// NewStruct creates a new relationship struct
func (*accessKeyR) NewStruct() *accessKeyR {
	return &accessKeyR{}
}

func (r *accessKeyR) GetUIDUser() *User {
	if r == nil {
		return nil
	}
	return r.UIDUser
}

// accessKeyL is where Load methods for each relationship are stored.
type accessKeyL struct{}

var (
	accessKeyAllColumns            = []string{"key", "uid", "client_type", "created_at", "updated_at"}
	accessKeyColumnsWithoutDefault = []string{"key", "uid", "client_type", "created_at", "updated_at"}
	accessKeyColumnsWithDefault    = []string{}
	accessKeyPrimaryKeyColumns     = []string{"key"}
	accessKeyGeneratedColumns      = []string{}
)

type (
	// AccessKeySlice is an alias for a slice of pointers to AccessKey.
	// This should almost always be used instead of []AccessKey.
	AccessKeySlice []*AccessKey
	// AccessKeyHook is the signature for custom AccessKey hook methods
	AccessKeyHook func(context.Context, boil.ContextExecutor, *AccessKey) error

	accessKeyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	accessKeyType                 = reflect.TypeOf(&AccessKey{})
	accessKeyMapping              = queries.MakeStructMapping(accessKeyType)
	accessKeyPrimaryKeyMapping, _ = queries.BindMapping(accessKeyType, accessKeyMapping, accessKeyPrimaryKeyColumns)
	accessKeyInsertCacheMut       sync.RWMutex
	accessKeyInsertCache          = make(map[string]insertCache)
	accessKeyUpdateCacheMut       sync.RWMutex
	accessKeyUpdateCache          = make(map[string]updateCache)
	accessKeyUpsertCacheMut       sync.RWMutex
	accessKeyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var accessKeyAfterSelectHooks []AccessKeyHook

var accessKeyBeforeInsertHooks []AccessKeyHook
var accessKeyAfterInsertHooks []AccessKeyHook

var accessKeyBeforeUpdateHooks []AccessKeyHook
var accessKeyAfterUpdateHooks []AccessKeyHook

var accessKeyBeforeDeleteHooks []AccessKeyHook
var accessKeyAfterDeleteHooks []AccessKeyHook

var accessKeyBeforeUpsertHooks []AccessKeyHook
var accessKeyAfterUpsertHooks []AccessKeyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AccessKey) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accessKeyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AccessKey) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accessKeyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AccessKey) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accessKeyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AccessKey) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accessKeyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AccessKey) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accessKeyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AccessKey) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accessKeyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AccessKey) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accessKeyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AccessKey) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accessKeyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AccessKey) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accessKeyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAccessKeyHook registers your hook function for all future operations.
func AddAccessKeyHook(hookPoint boil.HookPoint, accessKeyHook AccessKeyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		accessKeyAfterSelectHooks = append(accessKeyAfterSelectHooks, accessKeyHook)
	case boil.BeforeInsertHook:
		accessKeyBeforeInsertHooks = append(accessKeyBeforeInsertHooks, accessKeyHook)
	case boil.AfterInsertHook:
		accessKeyAfterInsertHooks = append(accessKeyAfterInsertHooks, accessKeyHook)
	case boil.BeforeUpdateHook:
		accessKeyBeforeUpdateHooks = append(accessKeyBeforeUpdateHooks, accessKeyHook)
	case boil.AfterUpdateHook:
		accessKeyAfterUpdateHooks = append(accessKeyAfterUpdateHooks, accessKeyHook)
	case boil.BeforeDeleteHook:
		accessKeyBeforeDeleteHooks = append(accessKeyBeforeDeleteHooks, accessKeyHook)
	case boil.AfterDeleteHook:
		accessKeyAfterDeleteHooks = append(accessKeyAfterDeleteHooks, accessKeyHook)
	case boil.BeforeUpsertHook:
		accessKeyBeforeUpsertHooks = append(accessKeyBeforeUpsertHooks, accessKeyHook)
	case boil.AfterUpsertHook:
		accessKeyAfterUpsertHooks = append(accessKeyAfterUpsertHooks, accessKeyHook)
	}
}

// One returns a single accessKey record from the query.
func (q accessKeyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AccessKey, error) {
	o := &AccessKey{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for access_keys")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AccessKey records from the query.
func (q accessKeyQuery) All(ctx context.Context, exec boil.ContextExecutor) (AccessKeySlice, error) {
	var o []*AccessKey

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AccessKey slice")
	}

	if len(accessKeyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AccessKey records in the query.
func (q accessKeyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count access_keys rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q accessKeyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if access_keys exists")
	}

	return count > 0, nil
}

// UIDUser pointed to by the foreign key.
func (o *AccessKey) UIDUser(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"uid\" = ?", o.UID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUIDUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (accessKeyL) LoadUIDUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAccessKey interface{}, mods queries.Applicator) error {
	var slice []*AccessKey
	var object *AccessKey

	if singular {
		var ok bool
		object, ok = maybeAccessKey.(*AccessKey)
		if !ok {
			object = new(AccessKey)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAccessKey)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAccessKey))
			}
		}
	} else {
		s, ok := maybeAccessKey.(*[]*AccessKey)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAccessKey)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAccessKey))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &accessKeyR{}
		}
		args = append(args, object.UID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &accessKeyR{}
			}

			for _, a := range args {
				if a == obj.UID {
					continue Outer
				}
			}

			args = append(args, obj.UID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.uid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.UIDUser = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UIDAccessKeys = append(foreign.R.UIDAccessKeys, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UID == foreign.UID {
				local.R.UIDUser = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UIDAccessKeys = append(foreign.R.UIDAccessKeys, local)
				break
			}
		}
	}

	return nil
}

// SetUIDUser of the accessKey to the related item.
// Sets o.R.UIDUser to related.
// Adds o to related.R.UIDAccessKeys.
func (o *AccessKey) SetUIDUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"access_keys\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"uid"}),
		strmangle.WhereClause("\"", "\"", 2, accessKeyPrimaryKeyColumns),
	)
	values := []interface{}{related.UID, o.Key}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UID = related.UID
	if o.R == nil {
		o.R = &accessKeyR{
			UIDUser: related,
		}
	} else {
		o.R.UIDUser = related
	}

	if related.R == nil {
		related.R = &userR{
			UIDAccessKeys: AccessKeySlice{o},
		}
	} else {
		related.R.UIDAccessKeys = append(related.R.UIDAccessKeys, o)
	}

	return nil
}

// AccessKeys retrieves all the records using an executor.
func AccessKeys(mods ...qm.QueryMod) accessKeyQuery {
	mods = append(mods, qm.From("\"access_keys\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"access_keys\".*"})
	}

	return accessKeyQuery{q}
}

// FindAccessKey retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAccessKey(ctx context.Context, exec boil.ContextExecutor, key string, selectCols ...string) (*AccessKey, error) {
	accessKeyObj := &AccessKey{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"access_keys\" where \"key\"=$1", sel,
	)

	q := queries.Raw(query, key)

	err := q.Bind(ctx, exec, accessKeyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from access_keys")
	}

	if err = accessKeyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return accessKeyObj, err
	}

	return accessKeyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AccessKey) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no access_keys provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(accessKeyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	accessKeyInsertCacheMut.RLock()
	cache, cached := accessKeyInsertCache[key]
	accessKeyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			accessKeyAllColumns,
			accessKeyColumnsWithDefault,
			accessKeyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(accessKeyType, accessKeyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(accessKeyType, accessKeyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"access_keys\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"access_keys\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into access_keys")
	}

	if !cached {
		accessKeyInsertCacheMut.Lock()
		accessKeyInsertCache[key] = cache
		accessKeyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AccessKey.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AccessKey) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	accessKeyUpdateCacheMut.RLock()
	cache, cached := accessKeyUpdateCache[key]
	accessKeyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			accessKeyAllColumns,
			accessKeyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update access_keys, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"access_keys\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, accessKeyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(accessKeyType, accessKeyMapping, append(wl, accessKeyPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update access_keys row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for access_keys")
	}

	if !cached {
		accessKeyUpdateCacheMut.Lock()
		accessKeyUpdateCache[key] = cache
		accessKeyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q accessKeyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for access_keys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for access_keys")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AccessKeySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accessKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"access_keys\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, accessKeyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in accessKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all accessKey")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AccessKey) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no access_keys provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(accessKeyColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	accessKeyUpsertCacheMut.RLock()
	cache, cached := accessKeyUpsertCache[key]
	accessKeyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			accessKeyAllColumns,
			accessKeyColumnsWithDefault,
			accessKeyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			accessKeyAllColumns,
			accessKeyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert access_keys, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(accessKeyPrimaryKeyColumns))
			copy(conflict, accessKeyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"access_keys\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(accessKeyType, accessKeyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(accessKeyType, accessKeyMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert access_keys")
	}

	if !cached {
		accessKeyUpsertCacheMut.Lock()
		accessKeyUpsertCache[key] = cache
		accessKeyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AccessKey record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AccessKey) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AccessKey provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), accessKeyPrimaryKeyMapping)
	sql := "DELETE FROM \"access_keys\" WHERE \"key\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from access_keys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for access_keys")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q accessKeyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no accessKeyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from access_keys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for access_keys")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AccessKeySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(accessKeyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accessKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"access_keys\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accessKeyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from accessKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for access_keys")
	}

	if len(accessKeyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AccessKey) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAccessKey(ctx, exec, o.Key)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccessKeySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AccessKeySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accessKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"access_keys\".* FROM \"access_keys\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accessKeyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AccessKeySlice")
	}

	*o = slice

	return nil
}

// AccessKeyExists checks if the AccessKey row exists.
func AccessKeyExists(ctx context.Context, exec boil.ContextExecutor, key string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"access_keys\" where \"key\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, key)
	}
	row := exec.QueryRowContext(ctx, sql, key)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if access_keys exists")
	}

	return exists, nil
}

// Exists checks if the AccessKey row exists.
func (o *AccessKey) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AccessKeyExists(ctx, exec, o.Key)
}
