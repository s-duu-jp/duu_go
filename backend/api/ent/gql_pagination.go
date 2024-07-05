// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/organization"
	"api/ent/photo"
	"api/ent/user"
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// OrganizationEdge is the edge representation of Organization.
type OrganizationEdge struct {
	Node   *Organization `json:"node"`
	Cursor Cursor        `json:"cursor"`
}

// OrganizationConnection is the connection containing edges to Organization.
type OrganizationConnection struct {
	Edges      []*OrganizationEdge `json:"edges"`
	PageInfo   PageInfo            `json:"pageInfo"`
	TotalCount int                 `json:"totalCount"`
}

func (c *OrganizationConnection) build(nodes []*Organization, pager *organizationPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Organization
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Organization {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Organization {
			return nodes[i]
		}
	}
	c.Edges = make([]*OrganizationEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &OrganizationEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// OrganizationPaginateOption enables pagination customization.
type OrganizationPaginateOption func(*organizationPager) error

// WithOrganizationOrder configures pagination ordering.
func WithOrganizationOrder(order *OrganizationOrder) OrganizationPaginateOption {
	if order == nil {
		order = DefaultOrganizationOrder
	}
	o := *order
	return func(pager *organizationPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultOrganizationOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithOrganizationFilter configures pagination filter.
func WithOrganizationFilter(filter func(*OrganizationQuery) (*OrganizationQuery, error)) OrganizationPaginateOption {
	return func(pager *organizationPager) error {
		if filter == nil {
			return errors.New("OrganizationQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type organizationPager struct {
	reverse bool
	order   *OrganizationOrder
	filter  func(*OrganizationQuery) (*OrganizationQuery, error)
}

func newOrganizationPager(opts []OrganizationPaginateOption, reverse bool) (*organizationPager, error) {
	pager := &organizationPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultOrganizationOrder
	}
	return pager, nil
}

func (p *organizationPager) applyFilter(query *OrganizationQuery) (*OrganizationQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *organizationPager) toCursor(o *Organization) Cursor {
	return p.order.Field.toCursor(o)
}

func (p *organizationPager) applyCursors(query *OrganizationQuery, after, before *Cursor) (*OrganizationQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultOrganizationOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *organizationPager) applyOrder(query *OrganizationQuery) *OrganizationQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultOrganizationOrder.Field {
		query = query.Order(DefaultOrganizationOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *organizationPager) orderExpr(query *OrganizationQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultOrganizationOrder.Field {
			b.Comma().Ident(DefaultOrganizationOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Organization.
func (o *OrganizationQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...OrganizationPaginateOption,
) (*OrganizationConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newOrganizationPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if o, err = pager.applyFilter(o); err != nil {
		return nil, err
	}
	conn := &OrganizationConnection{Edges: []*OrganizationEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := o.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if o, err = pager.applyCursors(o, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		o.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := o.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	o = pager.applyOrder(o)
	nodes, err := o.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// OrganizationOrderFieldName orders Organization by name.
	OrganizationOrderFieldName = &OrganizationOrderField{
		Value: func(o *Organization) (ent.Value, error) {
			return o.Name, nil
		},
		column: organization.FieldName,
		toTerm: organization.ByName,
		toCursor: func(o *Organization) Cursor {
			return Cursor{
				ID:    o.ID,
				Value: o.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f OrganizationOrderField) String() string {
	var str string
	switch f.column {
	case OrganizationOrderFieldName.column:
		str = "name"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f OrganizationOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *OrganizationOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("OrganizationOrderField %T must be a string", v)
	}
	switch str {
	case "name":
		*f = *OrganizationOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid OrganizationOrderField", str)
	}
	return nil
}

// OrganizationOrderField defines the ordering field of Organization.
type OrganizationOrderField struct {
	// Value extracts the ordering value from the given Organization.
	Value    func(*Organization) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) organization.OrderOption
	toCursor func(*Organization) Cursor
}

// OrganizationOrder defines the ordering of Organization.
type OrganizationOrder struct {
	Direction OrderDirection          `json:"direction"`
	Field     *OrganizationOrderField `json:"field"`
}

// DefaultOrganizationOrder is the default ordering of Organization.
var DefaultOrganizationOrder = &OrganizationOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &OrganizationOrderField{
		Value: func(o *Organization) (ent.Value, error) {
			return o.ID, nil
		},
		column: organization.FieldID,
		toTerm: organization.ByID,
		toCursor: func(o *Organization) Cursor {
			return Cursor{ID: o.ID}
		},
	},
}

// ToEdge converts Organization into OrganizationEdge.
func (o *Organization) ToEdge(order *OrganizationOrder) *OrganizationEdge {
	if order == nil {
		order = DefaultOrganizationOrder
	}
	return &OrganizationEdge{
		Node:   o,
		Cursor: order.Field.toCursor(o),
	}
}

// PhotoEdge is the edge representation of Photo.
type PhotoEdge struct {
	Node   *Photo `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// PhotoConnection is the connection containing edges to Photo.
type PhotoConnection struct {
	Edges      []*PhotoEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *PhotoConnection) build(nodes []*Photo, pager *photoPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Photo
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Photo {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Photo {
			return nodes[i]
		}
	}
	c.Edges = make([]*PhotoEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &PhotoEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// PhotoPaginateOption enables pagination customization.
type PhotoPaginateOption func(*photoPager) error

// WithPhotoOrder configures pagination ordering.
func WithPhotoOrder(order *PhotoOrder) PhotoPaginateOption {
	if order == nil {
		order = DefaultPhotoOrder
	}
	o := *order
	return func(pager *photoPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultPhotoOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithPhotoFilter configures pagination filter.
func WithPhotoFilter(filter func(*PhotoQuery) (*PhotoQuery, error)) PhotoPaginateOption {
	return func(pager *photoPager) error {
		if filter == nil {
			return errors.New("PhotoQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type photoPager struct {
	reverse bool
	order   *PhotoOrder
	filter  func(*PhotoQuery) (*PhotoQuery, error)
}

func newPhotoPager(opts []PhotoPaginateOption, reverse bool) (*photoPager, error) {
	pager := &photoPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultPhotoOrder
	}
	return pager, nil
}

func (p *photoPager) applyFilter(query *PhotoQuery) (*PhotoQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *photoPager) toCursor(ph *Photo) Cursor {
	return p.order.Field.toCursor(ph)
}

func (p *photoPager) applyCursors(query *PhotoQuery, after, before *Cursor) (*PhotoQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultPhotoOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *photoPager) applyOrder(query *PhotoQuery) *PhotoQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultPhotoOrder.Field {
		query = query.Order(DefaultPhotoOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *photoPager) orderExpr(query *PhotoQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultPhotoOrder.Field {
			b.Comma().Ident(DefaultPhotoOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Photo.
func (ph *PhotoQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...PhotoPaginateOption,
) (*PhotoConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newPhotoPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if ph, err = pager.applyFilter(ph); err != nil {
		return nil, err
	}
	conn := &PhotoConnection{Edges: []*PhotoEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := ph.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if ph, err = pager.applyCursors(ph, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		ph.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := ph.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	ph = pager.applyOrder(ph)
	nodes, err := ph.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// PhotoOrderFieldName orders Photo by name.
	PhotoOrderFieldName = &PhotoOrderField{
		Value: func(ph *Photo) (ent.Value, error) {
			return ph.Name, nil
		},
		column: photo.FieldName,
		toTerm: photo.ByName,
		toCursor: func(ph *Photo) Cursor {
			return Cursor{
				ID:    ph.ID,
				Value: ph.Name,
			}
		},
	}
	// PhotoOrderFieldURL orders Photo by url.
	PhotoOrderFieldURL = &PhotoOrderField{
		Value: func(ph *Photo) (ent.Value, error) {
			return ph.URL, nil
		},
		column: photo.FieldURL,
		toTerm: photo.ByURL,
		toCursor: func(ph *Photo) Cursor {
			return Cursor{
				ID:    ph.ID,
				Value: ph.URL,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f PhotoOrderField) String() string {
	var str string
	switch f.column {
	case PhotoOrderFieldName.column:
		str = "name"
	case PhotoOrderFieldURL.column:
		str = "url"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f PhotoOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *PhotoOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("PhotoOrderField %T must be a string", v)
	}
	switch str {
	case "name":
		*f = *PhotoOrderFieldName
	case "url":
		*f = *PhotoOrderFieldURL
	default:
		return fmt.Errorf("%s is not a valid PhotoOrderField", str)
	}
	return nil
}

// PhotoOrderField defines the ordering field of Photo.
type PhotoOrderField struct {
	// Value extracts the ordering value from the given Photo.
	Value    func(*Photo) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) photo.OrderOption
	toCursor func(*Photo) Cursor
}

// PhotoOrder defines the ordering of Photo.
type PhotoOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *PhotoOrderField `json:"field"`
}

// DefaultPhotoOrder is the default ordering of Photo.
var DefaultPhotoOrder = &PhotoOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &PhotoOrderField{
		Value: func(ph *Photo) (ent.Value, error) {
			return ph.ID, nil
		},
		column: photo.FieldID,
		toTerm: photo.ByID,
		toCursor: func(ph *Photo) Cursor {
			return Cursor{ID: ph.ID}
		},
	},
}

// ToEdge converts Photo into PhotoEdge.
func (ph *Photo) ToEdge(order *PhotoOrder) *PhotoEdge {
	if order == nil {
		order = DefaultPhotoOrder
	}
	return &PhotoEdge{
		Node:   ph,
		Cursor: order.Field.toCursor(ph),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	reverse bool
	order   *UserOrder
	filter  func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption, reverse bool) (*userPager, error) {
	pager := &userPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) (*UserQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultUserOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *userPager) applyOrder(query *UserQuery) *UserQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(DefaultUserOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *userPager) orderExpr(query *UserQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := u.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if u, err = pager.applyCursors(u, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	u = pager.applyOrder(u)
	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// UserOrderFieldSid orders User by sid.
	UserOrderFieldSid = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.Sid, nil
		},
		column: user.FieldSid,
		toTerm: user.BySid,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Sid,
			}
		},
	}
	// UserOrderFieldUID orders User by uid.
	UserOrderFieldUID = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.UID, nil
		},
		column: user.FieldUID,
		toTerm: user.ByUID,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.UID,
			}
		},
	}
	// UserOrderFieldName orders User by name.
	UserOrderFieldName = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.Name, nil
		},
		column: user.FieldName,
		toTerm: user.ByName,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Name,
			}
		},
	}
	// UserOrderFieldEmail orders User by email.
	UserOrderFieldEmail = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.Email, nil
		},
		column: user.FieldEmail,
		toTerm: user.ByEmail,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Email,
			}
		},
	}
	// UserOrderFieldRoleType orders User by role_type.
	UserOrderFieldRoleType = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.RoleType, nil
		},
		column: user.FieldRoleType,
		toTerm: user.ByRoleType,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.RoleType,
			}
		},
	}
	// UserOrderFieldStatusType orders User by status_type.
	UserOrderFieldStatusType = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.StatusType, nil
		},
		column: user.FieldStatusType,
		toTerm: user.ByStatusType,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.StatusType,
			}
		},
	}
	// UserOrderFieldOauthType orders User by oauth_type.
	UserOrderFieldOauthType = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.OauthType, nil
		},
		column: user.FieldOauthType,
		toTerm: user.ByOauthType,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.OauthType,
			}
		},
	}
	// UserOrderFieldSub orders User by sub.
	UserOrderFieldSub = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.Sub, nil
		},
		column: user.FieldSub,
		toTerm: user.BySub,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Sub,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserOrderField) String() string {
	var str string
	switch f.column {
	case UserOrderFieldSid.column:
		str = "sid"
	case UserOrderFieldUID.column:
		str = "uid"
	case UserOrderFieldName.column:
		str = "name"
	case UserOrderFieldEmail.column:
		str = "email"
	case UserOrderFieldRoleType.column:
		str = "role_type"
	case UserOrderFieldStatusType.column:
		str = "status_type"
	case UserOrderFieldOauthType.column:
		str = "oauth_type"
	case UserOrderFieldSub.column:
		str = "sub"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserOrderField %T must be a string", v)
	}
	switch str {
	case "sid":
		*f = *UserOrderFieldSid
	case "uid":
		*f = *UserOrderFieldUID
	case "name":
		*f = *UserOrderFieldName
	case "email":
		*f = *UserOrderFieldEmail
	case "role_type":
		*f = *UserOrderFieldRoleType
	case "status_type":
		*f = *UserOrderFieldStatusType
	case "oauth_type":
		*f = *UserOrderFieldOauthType
	case "sub":
		*f = *UserOrderFieldSub
	default:
		return fmt.Errorf("%s is not a valid UserOrderField", str)
	}
	return nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	// Value extracts the ordering value from the given User.
	Value    func(*User) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) user.OrderOption
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.ID, nil
		},
		column: user.FieldID,
		toTerm: user.ByID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
