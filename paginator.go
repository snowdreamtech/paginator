package paginator

const (
	// DefaultPage default page
	DefaultPage = 1

	// DefaultLimit default limit
	DefaultLimit = 10

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
	PaginateWithResult(list *[]interface{}, page int, limit int) Result

	// Paginate  return result.
	Paginate(page int, limit int) *interface{}
}
