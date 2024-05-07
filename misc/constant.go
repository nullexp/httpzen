package misc

type QueryOperator string

const (
	QueryOperatorEqual           QueryOperator = "eq"
	QueryOperatorNotEqual        QueryOperator = "neq"
	QueryOperatorMoreThan        QueryOperator = "mr"
	QueryOperatorEqualOrMoreThan QueryOperator = "mroeq"
	QueryOperatorLessThan        QueryOperator = "ls"
	QueryOperatorEqualOrLessThan QueryOperator = "lsoeq"
	QueryOperatorContain         QueryOperator = "cn"
	QueryOperatorNotContain      QueryOperator = "ncn"
	QueryOperatorEmpty           QueryOperator = "empt"
	QueryOperatorNotEmpty        QueryOperator = "nempt"
	QueryOperatorLike            QueryOperator = "like"
)
