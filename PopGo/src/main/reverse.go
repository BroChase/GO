package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"stack"
)

var expression_error error = errors.New("expression error")
var buffer string
var index int = 0
var expect_operand bool

//func evaluate() (int, error)
//evaluates the integer expression in Buffer and returns the result
//an Expression error is raised if the expression has an error
func evaluate() (float64, error) {
	var operand_stack = stack.New()
	var operator_stack = stack.New()
	var err error
	//precedence := func (op byte) int{...}
	//Returns the precedence of Operator. Raises Exception_error if
	//Operator is not a known operator.
	// '+' | '-' => 1
	// '*' | '-' => 2
	// '#' | '(' => 0
	//others raise exception
	precedence := func(op byte) int {
		//todo case statements
		switch op {
		case '+':
			return 1
		case '-':
			return 1
		case '*':
			return 2
		case '/':
			return 2
		case '#':
			return 0
		case '(':
			return 0
		default:
			panic("Illegal operator")
		}
	}

	//apply := func() error {..}
	//Applies the top operator on the Operator_Stack to its right and left
	//Operands on the operand stack
	apply := func() error {
		var op interface{}
		var left, right interface{}
		var err error
		//pop the operator off into op
		if err = operator_stack.Pop(&op); err != nil {
			panic("Pop failed operator error")
		}
		//pop the operand off into the right
		if err = operand_stack.Pop(&right); err != nil {
			panic("Pop failed operand error")
		}
		//pop the operand off into the left
		if err = operand_stack.Pop(&left); err != nil {
			panic("Pop failed operand error")
		}
		var R float64
		var L float64
		//fmt.Println(reflect.TypeOf(right))
		//tests the reflect of the interface to see if the value is a int or float
		//if it is a int then the reflect value is stored then that value is stored as a float
		v := reflect.TypeOf(right)
		if v.Kind() == reflect.Int {
			u := reflect.ValueOf(right)
			k := u.Interface()
			R = float64(k.(int))
			//fmt.Println(reflect.TypeOf(R))
		} else {
			R = right.(float64)
		}
		w := reflect.TypeOf(left)
		if w.Kind() == reflect.Int {
			e := reflect.ValueOf(left)
			q := e.Interface()
			L = float64(q.(int))
		} else {
			L = left.(float64)
		}

		//t := L+R
		//fmt.Println(t)

		switch op.(byte) {
		case '+':
			if err = operand_stack.Push(L + R); err != nil {
				panic("Error pushing value from addition")
			}
		case '-':
			if err = operand_stack.Push(L - R); err != nil {
				panic("Error pushing value from subtraction")
			}
		case '*':
			if err = operand_stack.Push(L * R); err != nil {
				panic("Error pushing value from Multiplication")
			}
		case '/':
			if err = operand_stack.Push(L / R); err != nil {
				panic("Error pushing value from Division")
			}
		default:
			panic("Illegal Operator")
		}
		return nil
	}
	//func() {...}()
	//Recover from a panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%v\n", r)
		}
	}()
	var Dec bool = false
	//process the expression left to right on character at a time.
	for index < len(buffer) {
		switch buffer[index] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			{
				//The character starts an operand. Extract it and push it on the operand stack
				value := 0
				var Pos int = 0
				var Count int = 0
				var Devi float64 = 1
				//loop repeats getting the values of the numbers in the string.
				//if the number has a decimal "float" then track where that decimal is and then convert the integer number
				//that is created into a float ex: "1.23 gets read into value 123 and then devided by 100 to give 1.23"
				for index < len(buffer) && buffer[index] >= '0' && buffer[index] <= '9' || index < len(buffer) && buffer[index] == '.' {
					if buffer[index] == '.' {
						Pos = Count
						Dec = true //decimal point was detected track its position in the "float"
						index += 1
					}
					value = value*10 + int(buffer[index]-'0')
					Count += 1
					index += 1
				}
				Pos = Count - Pos
				if Dec {
					for Pos != 0 {
						Devi = Devi * 10
						Pos -= 1
					}
					pvalue := float64(value) / Devi
					//fmt.Println(pvalue)
					if err = operand_stack.Push(pvalue); err != nil {
						panic("Failed to push float value")
					}
				} else {
					pvalue := value
					//fmt.Println(pvalue)
					if err = operand_stack.Push(pvalue); err != nil {
						panic("Failed to push int value")
					}
				}
				//fmt.Println("break")
				//operand_stack.Push(value)
				//expect an operator after an operand set operand expected to false and reset the Dec detection to false
				expect_operand = false
				Dec = false
			}
		case '+', '-', '*', '/':
			{
				//The character is an operator.  Apply any pending operators
				//on the operator_stack whos precedence is greater than or equal
				//to this operator.  Then, push the operator on the operator stack

				//if the top of the stack does not == nil then check the precedence of the operators
				//if the operator has a higher or equal precedence then do work else move on.
				//fmt.Println("operator")

				if operator_stack.Top() != nil {
					if precedence(buffer[index]) <= precedence(operator_stack.Top().(byte)) {
						apply()
					}
				}

				//push operator onto operator stack
				if err = operator_stack.Push(buffer[index]); err != nil {
					panic("Failed to push Operator to stack")
				}

				if expect_operand {
					return 0, expression_error
				}
				//expect an operand after operator.  set to true
				expect_operand = true
				//inc to next value in buffer
				index += 1
			}
		case ')':
			{
				//if the character is a right paren then loop back through the stack to
				//till a left paren doing work
				var ex interface{}

				ex = uint8(40)
				for operator_stack.Top() != ex {
					if precedence(operator_stack.Top().(byte)) > precedence(ex.(byte)) {
						apply()
					}
				}
				operator_stack.PopOff()
				index += 1
			}
		case '(':
			{
				//fmt.Println("left parens")
				if err = operator_stack.Push(buffer[index]); err != nil {
					panic("Failed to push LParens to stack")
				}
				index += 1
			}
		case ' ':
			{
				//The character is a space. Ignore it
				index += 1
			}
		//the character is something unexpected.  Return an expression error
		default:
			return 0, expression_error
		}
	}
	//we are at the end of the expression.  Apply all of the pending operators.  The operand stack must have exactly
	//one value, which is returned
	for !operator_stack.IsEmpty() {
		apply()
	}

	var value interface{}
	if err = operand_stack.Pop(&value); err != nil {
		panic("Pop Failed at result")
	}
	return value.(float64), nil
}

func main() {
	var result float64
	fmt.Println("Please input an operation expression without any space:")
	scanner := bufio.NewScanner(os.Stdin)
	//Process all of the expressions in standard input
	for scanner.Scan() {
		//read the next expression, evaluate it, and print the result.
		buffer = scanner.Text()
		fmt.Println(buffer)
		index = 0
		expect_operand = true
		var err error
		if result, err = evaluate(); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
		fmt.Printf("Result: %v\n", result)
	}
}
