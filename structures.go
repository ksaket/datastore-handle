package structures


type KindInfo struct {
	Name string
	Namespace string
}

type OpType struct {
	_ = iota
	UPDATE
	DELETE
	MODIFY_FIELD
	// UPSERT, INSERT ?
}

type Requisite struct {
	condition struct {
	relation struct {
		_ = iota
		EQUALS
		GREATER_THAN
		GREATER_EQ
		LESSER
		LESSER_EQ
	}
		rel relation
		param string
		value interface{}
	}

	conditions []condition
}