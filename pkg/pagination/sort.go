package pagination

//go:generate gomodifytags -all --quiet -w -file sort.go -clear-tags
//go:generate gomodifytags -all --quiet --skip-unexported -w -file sort.go -add-tags query,form,json,xml

type Sort struct {
	Query   string            `query:"query" form:"query" json:"query" xml:"query"`
	Columns map[string]string `query:"columns" form:"columns" json:"columns" xml:"columns"`
}
