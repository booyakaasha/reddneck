package db

type Row interface {
	Scan(dest ...any) error
}
