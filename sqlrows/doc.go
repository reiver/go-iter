/*
Package itersqlrows provides tools for creating iterators for *sql.Rows from the standard Go package "database/sql".

Example

	var rows *sql.Rows
	
	// ...
	
	iterator := itersqlrows.Rows{
		Rows: rows,
	}
*/
package itersqlrows
