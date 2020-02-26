package main

import (
	"bufio"
	"fmt"
	"github.com/JohnGeorge47/simpldb/sqlparser"
	"os"
	"strings"
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
	for  {
		printPrompt()
		text,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println("There was an error reading input")
			os.Exit(1)
		}
		input:=strings.Replace(text, "\n", "", -1)
		if input=="exit"{
			os.Exit(1)
		}
		fmt.Println(input)
		prep,stmnt:=sqlparser.PrepareResult(input)
		switch prep {
		case sqlparser.PREPARE_SUCCESS:
			break
		case sqlparser.PREPARE_UNRECOGNIZED_STATEMENT:
			fmt.Println("unrecognized keyword")
			continue
		}
        sqlparser.ExecuteStatement(stmnt)
		fmt.Println("statement has been executed")
	}
}

func printPrompt(){
	fmt.Print("db>")

}