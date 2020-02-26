package sqlparser

import (
	"fmt"
	"strings"
)

type Prepare string

type Stypes string

const  (
	STATEMENT_SELECT Stypes="STATEMENT_SELECT"
	STATEMENT_INSERT Stypes="STATEMENT_INSERT"
)

type Statement struct {
	StatementType Stypes
}

const (
	PREPARE_SUCCESS Prepare="PREPARE_SUCCESS"
	PREPARE_UNRECOGNIZED_STATEMENT Prepare="PREPARE_UNRECOGNIZED_STATEMENT"
)

func PrepareResult(stmnt string)(Prepare,*Statement){
	var statement Statement
	if strings.Index(stmnt,"SELECT")==0{
		statement.StatementType=STATEMENT_SELECT
		return PREPARE_SUCCESS,&statement
	}
	if strings.Index(stmnt,"INSERT")==0{
		statement.StatementType=STATEMENT_INSERT
		return PREPARE_SUCCESS,&statement
	}
	return PREPARE_UNRECOGNIZED_STATEMENT,nil
}

func ExecuteStatement(s *Statement){
	switch s.StatementType {
	case STATEMENT_SELECT:
		fmt.Println("added select")
		break
	case STATEMENT_INSERT:
		fmt.Println("added insert")
		break
	}
}