package main

const (
	scansText = `{{define "scans"}}// DON'T EDIT *** generated by sqlscan *** DON'T EDIT //

package {{.PackageName}}

import "database/sql"

{{range .Tokens}}
func {{$.Visibility}}can{{title .Name}}(r *sql.Row, s *{{.Name}}) error {
	return r.Scan({{range .Fields}}
		&s.{{.Name}},{{end}}
	)
}

func {{$.Visibility}}can{{title .Name}}s(rs *sql.Rows,ls []*{{.Name}}) error {
	var err error
	for rs.Next() {
		var s {{.Name}}
		if err = rs.Scan({{range .Fields}}
			&s.{{.Name}},{{end}}
		); err != nil {
			return err
		}
		ls = append(ls, &s)
	}
	return rs.Err()
}

{{end}}{{end}}`
)
