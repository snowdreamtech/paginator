package paginator

const (
	orderAsc  = "asc"
	orderDesc = "desc"
)

// paginator   paginator
type paginator struct {
	DB      *interface{}
	Query   Query
	Options Options
	Result  Result
}

// Paginator   Paginator
type Paginator interface {
	// Paginate  return result.
	PaginateWithResult(query Query, options Options) Result

	// Paginate  return result.
	Paginate(page int, limit int) *interface{}
}
