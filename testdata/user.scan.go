// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package testdata

import "database/sql"

func (s *Users) ScanRow(r *sql.Row) (*Users, error) {
	if err := r.Scan(
		s.Id,
		s.Username,
		s.Password,
		s.CreatedAt,
	); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Users) ScanRows(rs *sql.Rows) ([]Users, error) {
	structs := make([]Users, 0)
	var err error
	for rs.Next() {
		var s Users
		if err = rs.Scan(
			&s.Id,
			&s.Username,
			&s.Password,
			&s.CreatedAt,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}
