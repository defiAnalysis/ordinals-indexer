// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/adshao/ordinals-indexer/internal/data/ent/inscription"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Inscription is the model entity for the Inscription schema.
type Inscription struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// InscriptionID holds the value of the "inscription_id" field.
	InscriptionID int64 `json:"inscription_id,omitempty"`
	// UID holds the value of the "uid" field.
	UID string `json:"uid,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// OutputValue holds the value of the "output_value" field.
	OutputValue uint64 `json:"output_value,omitempty"`
	// ContentLength holds the value of the "content_length" field.
	ContentLength uint64 `json:"content_length,omitempty"`
	// ContentType holds the value of the "content_type" field.
	ContentType string `json:"content_type,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// GenesisHeight holds the value of the "genesis_height" field.
	GenesisHeight uint64 `json:"genesis_height,omitempty"`
	// GenesisFee holds the value of the "genesis_fee" field.
	GenesisFee uint64 `json:"genesis_fee,omitempty"`
	// GenesisTx holds the value of the "genesis_tx" field.
	GenesisTx string `json:"genesis_tx,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// Output holds the value of the "output" field.
	Output string `json:"output,omitempty"`
	// Offset holds the value of the "offset" field.
	Offset       uint64 `json:"offset,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Inscription) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case inscription.FieldID, inscription.FieldInscriptionID, inscription.FieldOutputValue, inscription.FieldContentLength, inscription.FieldGenesisHeight, inscription.FieldGenesisFee, inscription.FieldOffset:
			values[i] = new(sql.NullInt64)
		case inscription.FieldUID, inscription.FieldAddress, inscription.FieldContentType, inscription.FieldGenesisTx, inscription.FieldLocation, inscription.FieldOutput:
			values[i] = new(sql.NullString)
		case inscription.FieldCreatedAt, inscription.FieldUpdatedAt, inscription.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Inscription fields.
func (i *Inscription) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case inscription.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case inscription.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case inscription.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case inscription.FieldInscriptionID:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field inscription_id", values[j])
			} else if value.Valid {
				i.InscriptionID = value.Int64
			}
		case inscription.FieldUID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[j])
			} else if value.Valid {
				i.UID = value.String
			}
		case inscription.FieldAddress:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[j])
			} else if value.Valid {
				i.Address = value.String
			}
		case inscription.FieldOutputValue:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field output_value", values[j])
			} else if value.Valid {
				i.OutputValue = uint64(value.Int64)
			}
		case inscription.FieldContentLength:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field content_length", values[j])
			} else if value.Valid {
				i.ContentLength = uint64(value.Int64)
			}
		case inscription.FieldContentType:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_type", values[j])
			} else if value.Valid {
				i.ContentType = value.String
			}
		case inscription.FieldTimestamp:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[j])
			} else if value.Valid {
				i.Timestamp = value.Time
			}
		case inscription.FieldGenesisHeight:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field genesis_height", values[j])
			} else if value.Valid {
				i.GenesisHeight = uint64(value.Int64)
			}
		case inscription.FieldGenesisFee:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field genesis_fee", values[j])
			} else if value.Valid {
				i.GenesisFee = uint64(value.Int64)
			}
		case inscription.FieldGenesisTx:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field genesis_tx", values[j])
			} else if value.Valid {
				i.GenesisTx = value.String
			}
		case inscription.FieldLocation:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[j])
			} else if value.Valid {
				i.Location = value.String
			}
		case inscription.FieldOutput:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field output", values[j])
			} else if value.Valid {
				i.Output = value.String
			}
		case inscription.FieldOffset:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field offset", values[j])
			} else if value.Valid {
				i.Offset = uint64(value.Int64)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Inscription.
// This includes values selected through modifiers, order, etc.
func (i *Inscription) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// Update returns a builder for updating this Inscription.
// Note that you need to call Inscription.Unwrap() before calling this method if this Inscription
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Inscription) Update() *InscriptionUpdateOne {
	return NewInscriptionClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Inscription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Inscription) Unwrap() *Inscription {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Inscription is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Inscription) String() string {
	var builder strings.Builder
	builder.WriteString("Inscription(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("inscription_id=")
	builder.WriteString(fmt.Sprintf("%v", i.InscriptionID))
	builder.WriteString(", ")
	builder.WriteString("uid=")
	builder.WriteString(i.UID)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(i.Address)
	builder.WriteString(", ")
	builder.WriteString("output_value=")
	builder.WriteString(fmt.Sprintf("%v", i.OutputValue))
	builder.WriteString(", ")
	builder.WriteString("content_length=")
	builder.WriteString(fmt.Sprintf("%v", i.ContentLength))
	builder.WriteString(", ")
	builder.WriteString("content_type=")
	builder.WriteString(i.ContentType)
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(i.Timestamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("genesis_height=")
	builder.WriteString(fmt.Sprintf("%v", i.GenesisHeight))
	builder.WriteString(", ")
	builder.WriteString("genesis_fee=")
	builder.WriteString(fmt.Sprintf("%v", i.GenesisFee))
	builder.WriteString(", ")
	builder.WriteString("genesis_tx=")
	builder.WriteString(i.GenesisTx)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(i.Location)
	builder.WriteString(", ")
	builder.WriteString("output=")
	builder.WriteString(i.Output)
	builder.WriteString(", ")
	builder.WriteString("offset=")
	builder.WriteString(fmt.Sprintf("%v", i.Offset))
	builder.WriteByte(')')
	return builder.String()
}

// Inscriptions is a parsable slice of Inscription.
type Inscriptions []*Inscription
