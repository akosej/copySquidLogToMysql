package models

import "database/sql"

type Report struct {
	ID         uint           `db:"id"`
	Time       sql.NullTime   `db:"time"`
	Elapsed    sql.NullInt64  `db:"elapsed"`
	Remotehost sql.NullString `db:"remotehost"`
	CodeStatus sql.NullString `db:"code_status"`
	Bytes      sql.NullInt64  `db:"bytes"`
	Method     sql.NullString `db:"method"`
	URL        sql.NullString `db:"url"`
	Rfc931     sql.NullString `db:"rfc931"`
	PeerStatus sql.NullString `db:"peer_status"`
	Type       sql.NullString `db:"type"`
}
