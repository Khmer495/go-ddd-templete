// Code generated by entc, DO NOT EDIT.

package auth

const (
	// Label holds the string label denoting the auth type in the database.
	Label = "auth"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldFirebaseUserID holds the string denoting the firebase_user_id field in the database.
	FieldFirebaseUserID = "firebase_user_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the auth in the database.
	Table = "auths"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "auths"
	// UserInverseTable is the table name for the User model.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for auth fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldFirebaseUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
