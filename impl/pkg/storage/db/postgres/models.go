// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package postgres

import ()

type PkarrRecord struct {
	ID    int32
	Key   []byte
	Value []byte
	Sig   []byte
	Seq   int64
}