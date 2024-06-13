// Code generated by entc, DO NOT EDIT.

package address

import (
	"time"
)

const (
	// Label holds the string label denoting the address type in the database.
	Label = "address"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPostCode holds the string denoting the post_code field in the database.
	FieldPostCode = "post_code"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the core edge name in mutations.
	EdgeUser = "core"
	// Table holds the table name of the address in the database.
	Table = "addresses"
	// UserTable is the table the holds the core relation/edge.
	UserTable = "addresses"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "core" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the core relation/edge.
	UserColumn = "user_addresses"
)

// Columns holds all SQL columns for address fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldMobile,
	FieldAddress,
	FieldPostCode,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "addresses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_addresses",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
