package datastorehandler

type KindInfo struct {
	Name      string
	Namespace string
}

type OpType int

const (
	_ OpType = iota
	UPDATE
	DELETE
	MODIFY_FIELD
	// UPSERT, INSERT ?
)

type Relation int

const (
	_ Relation = iota
	EQUALS
	GREATER_THAN
	GREATER_EQ
	LESSER
	LESSER_EQ
)

type Condition struct {
	rel   Relation
	param string
	value interface{}
}

type Requisite struct {
	conditions []Condition
}
