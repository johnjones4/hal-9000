module main

go 1.17

replace github.com/johnjones4/hal-9000/server/hal9000 => ../hal9000

require github.com/johnjones4/hal-9000/server/hal9000 v0.0.0-00010101000000-000000000000

require (
	github.com/cdipaolo/goml v0.0.0-20210723214924-bf439dd662aa // indirect
	golang.org/x/text v0.3.7 // indirect
)