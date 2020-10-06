package paginator

// Result from orm.
type Result interface {
}

// PageResult from orm.
//
type PageResult struct {
	Page  int           `form:"page" json:"page" bson:"page" xml:"page" binding:"required"`
	Limit int           `form:"limit" json:"limit" bson:"limit" xml:"limit" binding:"required"`
	Total int           `form:"total" json:"total" bson:"total" xml:"total" `
	List  []interface{} `form:"list" json:"list" bson:"list" xml:"list"`
	Error error         `form:"error" json:"error" bson:"error" xml:"error"`
}

// CursorResult from orm.
// TODO
type CursorResult struct {
	Page  int           `form:"page" json:"page" bson:"page" xml:"page" binding:"required"`
	Limit int           `form:"limit" json:"limit" bson:"limit" xml:"limit" binding:"required"`
	Total int           `form:"total" json:"total" bson:"total" xml:"total" `
	List  []interface{} `form:"list" json:"list" bson:"list" xml:"list"`
	Error error         `form:"error" json:"error" bson:"error" xml:"error"`
}
