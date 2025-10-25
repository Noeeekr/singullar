package scan

import (
	"database/sql"

	"github.com/Noeeekr/borm"
)

type ScannerFunc func(*sql.Rows) error
type DatabaseScanner[Target any] struct {
	Scanner ScannerFunc
}

func Scanner[Target any](ScannerFunc ScannerFunc) *DatabaseScanner[Target] {
	return &DatabaseScanner[Target]{
		Scanner: ScannerFunc,
	}
}

func (d *DatabaseScanner[Target]) On(target *[]*Target) borm.ReturnScanner {
	return func(rows *sql.Rows) (found bool, err error) {
		defer rows.Close()
		if target == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for rows.Next() {
			if err := d.Scanner(rows); err != nil {
				return false, err
			}
		}

		if err := rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		return len(*target) != 0, nil
	}
}
