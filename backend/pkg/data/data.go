// package data will provide wrappers to the type sql.*
package data

import "database/sql"

type NullString struct {
	sql.NullString
}

func ToNullString(s string) NullString {
	return NullString{
		sql.NullString{
			Valid:  true,
			String: s,
		},
	}
}

type NullInt64 struct {
	sql.NullInt64
}

func ToNullInt(i int) NullInt64 {
	return NullInt64{
		sql.NullInt64{
			Valid: true,
			Int64: int64(i),
		},
	}
}

type NullBool struct {
	sql.NullBool
}

func ToNullBool(b bool) NullBool {
	return NullBool{
		sql.NullBool{
			Valid: true,
			Bool:  b,
		},
	}
}

type NullFloat64 struct {
	sql.NullFloat64
}

func ToNullFloat64(f float64) NullFloat64 {
	return NullFloat64{
		sql.NullFloat64{
			Valid:   true,
			Float64: f,
		},
	}
}
