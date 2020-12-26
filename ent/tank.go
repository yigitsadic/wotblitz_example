// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/yigitsadic/wotblitz_example/ent/tank"
)

// Tank is the model entity for the Tank schema.
type Tank struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Tier holds the value of the "tier" field.
	Tier int `json:"tier,omitempty"`
	// IsPremium holds the value of the "isPremium" field.
	IsPremium bool `json:"isPremium,omitempty"`
	// TankClass holds the value of the "tankClass" field.
	TankClass string `json:"tankClass,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TankQuery when eager-loading is set.
	Edges TankEdges `json:"edges"`
}

// TankEdges holds the relations/edges for other nodes in the graph.
type TankEdges struct {
	// FromTankId holds the value of the fromTankId edge.
	FromTankId []*Tank
	// NextTanks holds the value of the nextTanks edge.
	NextTanks []*Tank
	// NextTankId holds the value of the nextTankId edge.
	NextTankId []*Tank
	// PreviousTanks holds the value of the previousTanks edge.
	PreviousTanks []*Tank
	// Modules holds the value of the modules edge.
	Modules []*Module
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// FromTankIdOrErr returns the FromTankId value or an error if the edge
// was not loaded in eager-loading.
func (e TankEdges) FromTankIdOrErr() ([]*Tank, error) {
	if e.loadedTypes[0] {
		return e.FromTankId, nil
	}
	return nil, &NotLoadedError{edge: "fromTankId"}
}

// NextTanksOrErr returns the NextTanks value or an error if the edge
// was not loaded in eager-loading.
func (e TankEdges) NextTanksOrErr() ([]*Tank, error) {
	if e.loadedTypes[1] {
		return e.NextTanks, nil
	}
	return nil, &NotLoadedError{edge: "nextTanks"}
}

// NextTankIdOrErr returns the NextTankId value or an error if the edge
// was not loaded in eager-loading.
func (e TankEdges) NextTankIdOrErr() ([]*Tank, error) {
	if e.loadedTypes[2] {
		return e.NextTankId, nil
	}
	return nil, &NotLoadedError{edge: "nextTankId"}
}

// PreviousTanksOrErr returns the PreviousTanks value or an error if the edge
// was not loaded in eager-loading.
func (e TankEdges) PreviousTanksOrErr() ([]*Tank, error) {
	if e.loadedTypes[3] {
		return e.PreviousTanks, nil
	}
	return nil, &NotLoadedError{edge: "previousTanks"}
}

// ModulesOrErr returns the Modules value or an error if the edge
// was not loaded in eager-loading.
func (e TankEdges) ModulesOrErr() ([]*Module, error) {
	if e.loadedTypes[4] {
		return e.Modules, nil
	}
	return nil, &NotLoadedError{edge: "modules"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tank) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullInt64{},  // tier
		&sql.NullBool{},   // isPremium
		&sql.NullString{}, // tankClass
		&sql.NullString{}, // country
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tank fields.
func (t *Tank) assignValues(values ...interface{}) error {
	if m, n := len(values), len(tank.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		t.Name = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field tier", values[1])
	} else if value.Valid {
		t.Tier = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field isPremium", values[2])
	} else if value.Valid {
		t.IsPremium = value.Bool
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field tankClass", values[3])
	} else if value.Valid {
		t.TankClass = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field country", values[4])
	} else if value.Valid {
		t.Country = value.String
	}
	return nil
}

// QueryFromTankId queries the fromTankId edge of the Tank.
func (t *Tank) QueryFromTankId() *TankQuery {
	return (&TankClient{config: t.config}).QueryFromTankId(t)
}

// QueryNextTanks queries the nextTanks edge of the Tank.
func (t *Tank) QueryNextTanks() *TankQuery {
	return (&TankClient{config: t.config}).QueryNextTanks(t)
}

// QueryNextTankId queries the nextTankId edge of the Tank.
func (t *Tank) QueryNextTankId() *TankQuery {
	return (&TankClient{config: t.config}).QueryNextTankId(t)
}

// QueryPreviousTanks queries the previousTanks edge of the Tank.
func (t *Tank) QueryPreviousTanks() *TankQuery {
	return (&TankClient{config: t.config}).QueryPreviousTanks(t)
}

// QueryModules queries the modules edge of the Tank.
func (t *Tank) QueryModules() *ModuleQuery {
	return (&TankClient{config: t.config}).QueryModules(t)
}

// Update returns a builder for updating this Tank.
// Note that, you need to call Tank.Unwrap() before calling this method, if this Tank
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tank) Update() *TankUpdateOne {
	return (&TankClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Tank) Unwrap() *Tank {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tank is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tank) String() string {
	var builder strings.Builder
	builder.WriteString("Tank(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", tier=")
	builder.WriteString(fmt.Sprintf("%v", t.Tier))
	builder.WriteString(", isPremium=")
	builder.WriteString(fmt.Sprintf("%v", t.IsPremium))
	builder.WriteString(", tankClass=")
	builder.WriteString(t.TankClass)
	builder.WriteString(", country=")
	builder.WriteString(t.Country)
	builder.WriteByte(')')
	return builder.String()
}

// Tanks is a parsable slice of Tank.
type Tanks []*Tank

func (t Tanks) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
