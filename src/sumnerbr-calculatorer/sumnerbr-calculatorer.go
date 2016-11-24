// BRIAN SUMNER
// UCDENVER CSCI3415-001
// FALL 2016
// PA03 - SUMNERBR-CALCULATORER
// NOTE: A CALCULATORER DOES THE CALCULATORING.

// (MODIFIED FROM INSTRUCTOR EXAMPLE USING INSTRUCTOR STACK PACKAGE)
// REFERENCE USED: https://golang.org/doc/effective_go.html
// REFERENCE USED: https://blog.golang.org/laws-of-reflection

package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"stack"
)

var operatorStack = stack.NewStack()
var operandStack = stack.NewStack()

// PANICKER PANICS.
// PANICKER DOES ALL THE PANICKING.
// DO NOT PANIC UNLESS PANICKER PANICS.
// PANICKER IS THE ONLY DRAMA QUEEN HERE.
// NOTE:  IT IS NOT PANICKER'S JOB TO FIX EXPRESSIONS.
func panicker(inputString string) {

	// CAUTION:  PANICKER MAY BECOME VERY UPSET.
	PANICMSG_AAAAAGGHH := " YOU'VE KILLED ME!  YOU DID THIS TO ME!!  YOU DID THIS!!!\n\n"

	// CHECK NOT EMPTY STRING
	if inputString == "" {
		panic("\n\nERR_01:  OH NO!  I'VE DIED OF BOREDOM!!!\n\n")
	}

	// CHECK IF SPECIAL CASE
	if inputString == "/0" {
		panic("\n\nERR_11:  I DON'T KNOW HOW TO DIVIDE BY ZERO, ARE YOU CRAZY?!\n         I CAN'T GO ON LIVING LIKE THIS!!!\n\n")
	} else if inputString == "???" {
		panic("\n\nERR_12:  I DON'T UNDERSTAND WHAT YOU'RE TRYING TO DO TO ME!  LEAVE ME ALONE ALREADY!!!\n\n")
	}

	// CHECK ALL CHARACTERS ARE VALID
	// CHECK NUMBER OF OPEN PARENTHESES NOT NEGATIVE
	// CHECK NUMBER OF DECIMAL POINTS PER OPERAND
	// CHECK ARITHMETIC SYNTAX IS CORRECT
	openParentheses := int(0)
	activeDecimalPoints := int(0)
	isIncompleteOperand := false
	isNegativeOperand := false
	doesNeedOperand := true
	doesNeedOperator := false

	for _, inputChar := range inputString {

		switch inputChar {

		// FOUND OPERAND
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':

			// CHECK IF NEEDS OPERATOR INSTEAD
			if doesNeedOperator {
				panic("\n\nERR_21: " + PANICMSG_AAAAAGGHH)
			}

			// LET OPERAND CONTINUE
			isIncompleteOperand = true

			// NEXT CHARACTER COULD BUT NEED NOT BE OPERAND
			doesNeedOperand = false

		// FOUND DECIMAL POINT
		case '.':

			// CHECK IF NEEDS OPERATOR INSTEAD
			if doesNeedOperator {
				panic("\n\nERR_31: " + PANICMSG_AAAAAGGHH)
			}

			// CHECK IF OPERAND ALREADY CONTAINS DECIMAL POINT
			if activeDecimalPoints > 0 {
				panic("\n\nERR_32: " + PANICMSG_AAAAAGGHH)
			}

			// INCREMENT DECIMAL POINT COUNTER
			activeDecimalPoints++

		// FOUND SPACE CHAR
		case ' ':

			// NOTE: NEGATIVE OPERAND FUNCTIONALITY NOT IMPLEMENTED
			if isNegativeOperand {
				panic("\n\nERR_41" + PANICMSG_AAAAAGGHH)
			}

			// IF SPACE FOUND AFTER OPERAND THEN PREVIOUS OPERAND IS COMPLETE
			if isIncompleteOperand {
				doesNeedOperator = true
				doesNeedOperand = false
			}

			// RESET DECIMAL POINT COUNTER
			activeDecimalPoints = 0

		// FOUND OPERATOR
		case '+', '-', '*', '/':

			// NOTE: NEGATIVE OPERAND FUNCTIONALITY NOT IMPLEMENTED
			if doesNeedOperand && inputChar == '-' {

				// THIS PROBABLY INDICATES NEGATIVE VALUE FOR OPERAND
				panic("\n\nERR_51:  THEY DIDN'T TEACH ME NEGATIVE NUMBERS!  I JUST CAN'T GO ON LIVING!!!\n\n")
			} else {

				// CHECK IF NEEDS OPERAND INSTEAD
				if doesNeedOperand {
					panic("\n\nERR_52: " + PANICMSG_AAAAAGGHH)
				}

				// PREVIOUS OPERAND IS COMPLETE
				isIncompleteOperand = false

				// DOES NOT NEED ANOTHER OPERATOR
				doesNeedOperator = false

				// NEEDS OPERAND NEXT
				doesNeedOperand = true

				// RESET DECIMAL POINT COUNTER
				activeDecimalPoints = 0
			}

		// FOUND BEGINPARENS
		case '(':

			// NOTE: BEGINPARENS MAY ONLY FOLLOW AN OPERATOR
			if isIncompleteOperand || doesNeedOperator {
				panic("\n\nERR_61: " + PANICMSG_AAAAAGGHH)
			}

			// INCREMENT OPEN PARENTHESES COUNTER
			openParentheses++

			// PREVIOUS OPERAND IS COMPLETE
			isIncompleteOperand = false

			// DOES NOT NEED ANOTHER OPERATOR
			doesNeedOperator = false

			// RESET DECIMAL POINT COUNTER
			activeDecimalPoints = 0

		// FOUND ENDPARENS
		case ')':

			// CHECK IF NUMBER OF ENDPARENS EXCEEDS NUMBER OF BEGINPARENS
			if openParentheses < 1 {
				panic("\n\nERR_71: " + PANICMSG_AAAAAGGHH)
			}

			// DECREMENT OPEN PARENS COUNTER
			openParentheses--

			// RESET DECIMAL POINT COUNTER
			activeDecimalPoints = 0

			// PREVIOUS OPERAND IS COMPLETE
			isIncompleteOperand = false

			// NEEDS OPERATOR NEXT
			doesNeedOperator = true

		// FOUND INVALID CHARACTER
		default:
			panic("\n\nERR_81: " + PANICMSG_AAAAAGGHH)
		}
	}

	switch {
	// CHECK IF LINE ENDS WITH OPERATOR INSTEAD OF OPERAND
	case doesNeedOperand:
		panic("\n\nERR_91: " + PANICMSG_AAAAAGGHH)
	}
}

// PARENTHESIZER PARENTHESIZES.
// SPECIFICALLY, PARENTHESIZER COMPLETES INCOMPLETE PARETHESIZATION.
// IF NUM_BEGINPARENS < NUM_ENDPARENS THEN PANICKER WILL PANIC.
// IF PANICKER PANICS, PARENTHESIZER WILL NOT PARENTHESIZE.
// NOTE:  IT'S NOT PANICKER'S JOB TO PARENTHESIZE EXPRESSIONS.
// NOTE:  THAT'S WHAT PARENTHESIZER IS FOR.
// ALSO:  PARENTHESIZER ENSURES ALL EXPRESSIONS BEGIN WITH A BEGINPARENS.
func parenthesizer(inputString string) (outputString string) {

	// DETERMINE IF FIRST CHAR IS BEGINPARENS
	isFirstCharParens := false
	if inputString[0] == '(' {
		isFirstCharParens = true
	}

	// COUNT OPEN PARENTHESES
	// NOTE: PANICKER ALREADY DID THIS, BUT IT'S NOT PANICKER'S JOB TO FIX EXPRESSIONS
	openParentheses := int(0)
	for _, inputChar := range inputString {
		switch inputChar {
		case '(':
			openParentheses++
		case ')':
			openParentheses--
		}
	}

	// USE DEDICATED ADJUSTMENT STRING FOR A LIGHTWEIGHT APPROACH
	// NOTE:  ADJUSTMENT STRING AVOIDS RECOPYING ENTIRE EXPRESSION
	additionalEndParens := string("")
	for index := 0; index < openParentheses; index++ {
		additionalEndParens += ")"
	}

	// ENSURE FIRST CHAR IS BEGINPARENS
	if isFirstCharParens {
		outputString = inputString + additionalEndParens
	} else {
		outputString = "(" + inputString + additionalEndParens + ")"
	}

	return
}

// PRECEDENCER DOES THE PRECEDENCING.
// SPECIFICALLY, PRECEDENCER WILL RETURN THE PRECEDENCE VALUE OF A GIVEN OPERATOR.
func precedencer(op byte) uint8 {
	switch op {

	// NOTE: BEGINPARENS HAS LOWEST PRECEDENCE. ENDPARENS IS NOT APPLICABLE HERE
	case '(':
		return 0
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		// CANNOT REMOVE DEFAULT CASE OR SYNTAX ERROR IS INCURRED
		// CANNOT CALL PANICKER (FOR NONEXISTENT THREAT) OR SYNTAX ERROR IS INCURRED
		// THEREFORE, PRETEND THE FOLLOWING STATEMENT IS USEFUL.
		panic("I am a very calm function.  I promise I will never panic.")
	}
}

// CALCULATORER CALCULATORS.
// CALCULATORER DOES ALL THE CALCULATORING.
func calculatorer() {

	// POP OPERATOR OFF OPERATOR STACK
	op := operatorStack.Pop().(byte)

	// NOTE:  ANY BEGINPARENS WILL SIMPLY BE DISCARDED FROM THE STACK BEFORE RETURNING
	if op != '(' {

		// STORE THE REFLECT TYPE OF THE RIGHT OPERAND
		rightType := reflect.TypeOf(operandStack.Top())

		// INITIALIZE RIGHT OPERAND VARIABLES FOR BOTH TYPES
		// NOTE:  DEPENDING ON REFLECT TYPE ONLY ONE VARIABLE WILL BE USED
		rightFloat64 := float64(0)
		rightInt := int(0)

		switch rightType {

		// IF REFLECT TYPE OF RIGHT OPERAND MATCHES REFLECT TYPE OF KNOWN FLOAT:
		case reflect.TypeOf(rightFloat64):

			// THEN POP RIGHT OPERAND AS FLOAT
			// NOTE: CANNOT POP OPERAND AS TYPE: REFLECT TYPE OF STACK TOP
			// DOES NOT WORK: right = operandStack.Pop(reflect.TypeOf(operandStack.Top())
			rightFloat64 = operandStack.Pop().(float64)

		// IF REFLECT TYPE OF RIGHT OPERAND MATCHES REFLECT TYPE OF KNOWN INT:
		case reflect.TypeOf(rightInt):

			// THEN POP RIGHT OPERAND AS INT
			// NOTE: CANNOT POP OPERAND AS TYPE: REFLECT TYPE OF STACK TOP
			// DOES NOT WORK: left = operandStack.Pop(reflect.TypeOf(operandStack.Top())
			rightInt = operandStack.Pop().(int)
		}

		// STORE THE REFLECT TYPE OF THE LEFT OPERAND
		leftType := reflect.TypeOf(operandStack.Top())

		// INITIALIZE LEFT OPERAND VARIABLES FOR BOTH TYPES
		// NOTE:  DEPENDING ON REFLECT TYPE ONLY ONE VARIABLE WILL BE USED
		leftFloat64 := float64(0)
		leftInt := int(0)

		switch leftType {

		// IF REFLECT TYPE OF LEFT OPERAND MATCHES REFLECT TYPE OF KNOWN FLOAT:
		case reflect.TypeOf(leftFloat64):

			// THEN POP LEFT OPERAND AS FLOAT
			// NOTE: CANNOT POP OPERAND AS TYPE: REFLECT TYPE OF STACK TOP
			leftFloat64 = operandStack.Pop().(float64)

		// IF REFLECT TYPE OF RIGHT OPERAND MATCHES REFLECT TYPE OF KNOWN INT:
		case reflect.TypeOf(leftInt):

			// THEN POP LEFT OPERAND AS INT
			// NOTE: CANNOT POP OPERAND AS TYPE: REFLECT TYPE OF STACK TOP
			leftInt = operandStack.Pop().(int)
		}

		// DO CALCULATORING BASED ON REFLECT TYPE CASE
		// NOTE:  MULTIPLICATION AND DIVISION WITH ANY FLOAT64 OPERAND WILL NEVER COERCE TO INT!!!

		// REFLECTION CASE:  FLOAT64 <OPERATOR> FLOAT64 = FLOAT64
		if leftType == reflect.TypeOf(leftFloat64) && rightType == reflect.TypeOf(rightFloat64) {
			switch op {
			case '+':
				operandStack.Push(leftFloat64 + rightFloat64)
			case '-':
				operandStack.Push(leftFloat64 - rightFloat64)
			case '*':
				operandStack.Push(leftFloat64 * rightFloat64)
			case '/':
				if rightFloat64 == 0 {
					panicker("/0") // CANNOT DIVIDE BY 0
				} else {
					operandStack.Push(leftFloat64 / rightFloat64)
				}
			}
		// REFLECTION CASE:  FLOAT64 <OPERATOR> INT = FLOAT64
		} else if leftType == reflect.TypeOf(leftFloat64) && rightType == reflect.TypeOf(rightInt) {
			switch op {
			case '+':
				operandStack.Push(leftFloat64 + float64(rightInt))
			case '-':
				operandStack.Push(leftFloat64 - float64(rightInt))
			case '*':
				operandStack.Push(leftFloat64 * float64(rightInt))
			case '/':
				if rightInt == 0 {
					panicker("/0") // CANNOT DIVIDE BY 0
				} else {
					operandStack.Push(leftFloat64 / float64(rightInt))
				}
			}
		// REFLECTION CASE:  INT <OPERATOR> FLOAT64 = FLOAT64
		} else if leftType == reflect.TypeOf(leftInt) && rightType == reflect.TypeOf(rightFloat64) {
			switch op {
			case '+':
				operandStack.Push(float64(leftInt) + rightFloat64)
			case '-':
				operandStack.Push(float64(leftInt) - rightFloat64)
			case '*':
				operandStack.Push(float64(leftInt) * rightFloat64)
			case '/':
				if rightFloat64 == 0 {
					panicker("/0") // CANNOT DIVIDE BY 0
				} else {
					operandStack.Push(float64(leftInt) / rightFloat64)
				}
			}
		// REFLECTION CASE:  INT <OPERATOR> INT = INT
		} else if leftType == reflect.TypeOf(leftInt) && rightType == reflect.TypeOf(rightInt) {
			switch op {
			case '+':
				operandStack.Push(leftInt + rightInt)
			case '-':
				operandStack.Push(leftInt - rightInt)
			case '*':
				operandStack.Push(leftInt * rightInt)
			case '/':
				if rightInt == 0 {
					panicker("/0") // CANNOT DIVIDE BY 0
				} else {
					operandStack.Push(leftInt / rightInt)
				}
			}
		} else {
			// UNNECESSARY STATEMENT FOR NONEXISTENT CONDITION
			panicker("???")
		}
	}
}

func main() {

	// NOTE:  AN INFINITE LOOP IS THE ONLY WAY TO CALCULATOR PROPERLY.
	// PANICKER WILL BREAK THE LOOP AS NEEDED.
	for true {

		fmt.Println("\n\nPLEASE ENTER A SIMPLE ARITHMETIC EXPRESSION OR PRESS ENTER TO EXIT:")
		fmt.Print(">  ")

		// Read a from Stdin.
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputLine := scanner.Text()
		fmt.Print("\n")

		// NOTE:  PANICKER PANICS.
		// IF PANICKER DOESN'T PANIC, WE CAN CONTINUE WITH ASSURANCE THAT NOTHING WILL PANIC.
		// EXCEPT MAYBE FROM DIVIDE BY ZERO.  THAT MIGHT PANIC, WHICH IS NOT PANICKER'S FAULT.
		// PANICKER HAD A BAD FEELING ABOUT YOUR DIVISION BUT WAS TOO POLITE TO PANIC AT THE TIME.
		// PANICKER PREVENTS THE REST OF THE PROGRAM FROM ATTEMPTING TO CALCULATOR YOUR NONSENSE.
		// READ:  DO NOT POPULATE STACK OBJECTS UNLESS PANICKER DOES NOT PANIC.
		panicker(inputLine)

		// FULLY PARENTHESIZE INPUTLINE
		inputLine = parenthesizer(inputLine)
		line := inputLine

		// EVALUATE EXPRESSION IN SINGLE-PASS SCAN
		for i := 0; i < len(line); {
			switch line[i] {

			// FOUND OPERAND
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':

				// INITIALIZE VARIABLES FOR BOTH TYPES
				// NOTE:  BOTH VARIABLES MAY BE USED
				valueInt := int(0)
				valueFloat64 := float64(0)

				// HAS THE DOT BEEN FOUND
				seenDecimalPoint := false

				// EXPLICIT FLOAT64 MULTIPLIER FOR FLOATING POINT EVALUATION
				decimalMultiplier := float64(0.1)

				// CONTINUE FINDING OPERAND
				for {
					if line[i] == '.' {

						// FOUND THE DOT
						seenDecimalPoint = true

						// CONVERT INTEGER VALUE TO FLOAT64
						// NOTE:  NOW SKIP TO NEXT CHARACTER
						valueFloat64 = float64(valueInt)
					} else if seenDecimalPoint {

						// CONTINUE ACCUMULATING DECIMAL VALUE OF FLOAT64 OPERAND
						valueFloat64 = valueFloat64 + (float64(int(line[i]-'0')) * decimalMultiplier)

						// RATCHET THE MULTIPLIER
						decimalMultiplier /= float64(10)
					} else {

						// DOT NOT FOUND SO ACCUMULATE THE INTEGER
						valueInt = valueInt*10 + int(line[i]-'0')
					}

					// SKIP TO NEXT CHAR
					i++

					// IF NEXT CHAR IS EOL OR NOT OPERAND, BREAK THE CONTINUE FINDING OPERAND LOOP
					if i == len(line) ||
						!(('0' <= line[i] && line[i] <= '9') || line[i] == '.') {
						break
					}
				}

				// NOTE:  BREAK JUST OCCURRED
				if seenDecimalPoint {

					// DOT WAS FOUND SO PUSH FLOAT64
					operandStack.Push(valueFloat64)
				} else {
					// DOT WAS NOT FOUND SO PUSH INT
					operandStack.Push(valueInt)
				}

			// FOUND OPERATOR
			case '+', '-', '*', '/':

				// CALCULATOR ALL CURRENT OPERATIONS ALREADY ON THE STACKS PROVIDED THAT:
				// 1. THERE ARE STILL OPERATIONS ON THE STACK TO CALCULATOR, AND
				// 2. ALL EXISTING OPERATIONS BELONG TO THE CURRENT PARENTHETIC SUBEXPRESSION, AND
				// 3. THE EXISTING OPERATION HAS HIGHER PRECEDENCE THAN THE CURRENT FOUND OPERATOR
				for !operatorStack.IsEmpty() && operatorStack.Top().(byte) != '(' &&
					precedencer(operatorStack.Top().(byte)) >= precedencer(line[i]) {
					calculatorer()
				}

				// PUSH THE FOUND OPERATOR ONTO OPERATOR STACK
				operatorStack.Push(line[i])

				// SKIP TO NEXT CHAR
				i++

			// FOUND SPACE CHAR
			case ' ':

				// SKIP TO NEXT CHAR
				// NOTE:  IF A SPACE SPLITS AN OPERAND THEN PANICKER WILL PANIC
				// NOTE:  IF PANICKER PANICS THEN THIS CODE WILL NOT RUN
				i++

			// FOUND BEGINPARENS
			case '(':

				// PUSH BEGINPARENS TO OPERATOR STACK THEN SKIP TO NEXT CHAR
				operatorStack.Push(line[i])
				i++

			// FOUND ENDPARENS
			case ')':

				// CALCULATOR ALL CURRENT OPERATIONS CURRENTLY ON THE STACKS PROVIDED THAT:
				// 1. THERE ARE STILL OPERATIONS ON THE STACK TO CALCULATOR, AND
				// 2. ALL EXISTING OPERATIONS BELONG TO THE CURRENT PARENTHETIC SUBEXPRESSION
				for !operatorStack.IsEmpty() && operatorStack.Top().(byte) != '(' {
					calculatorer()
				}

				// CALCULATOR THE TOP OPERATION ON THE STACKS PROVIDED THAT:
				// 1. THERE IS STILL AN OPERATION ON THE STACK TO CALCULATOR
				// NOTE:  THIS COMBINES THE CURRENT PARENTHETIC SUBEXPRESSION WITH THE PREVIOUS ONE
				if !operatorStack.IsEmpty() {
					calculatorer()
				}

				// SKIP TO NEXT CHAR
				i++
			}
		}

		// NOTE:  END OF EXPRESSION INPUTLINE
		// CALCULATOR ALL REMAINING OPERATIONS UNTIL THE OPERATOR STACK IS EMPTY
		for !operatorStack.IsEmpty() {
			calculatorer()
		}

		// DISPLAY THE PARENTHESIZED EXPRESSION
		fmt.Print(inputLine + "  =  ")

		// INITIALIZE RESULT VARIABLES OF BOTH TYPES
		// NOTE:  DEPENDING ON REFLECT TYPE ONLY ONE VARIABLE WILL BE USED
		resultFloat64 := float64(0)
		resultInt := int(0)

		// DISPLAY RESULT BASED ON REFLECT TYPE OF STACK TOP
		switch reflect.TypeOf(operandStack.Top()) {

		// REFLECTION CASE:  REFLECT TYPE OF RESULT MATCHES REFLECT TYPE OF KNOWN FLOAT64:
		case reflect.TypeOf(resultFloat64):

			// POP RESULT AS FLOAT64 AND DISPLAY TO OUTPUT WITH TYPE
			resultFloat64 = operandStack.Pop().(float64)
			fmt.Print(resultFloat64)
			fmt.Println("  [TYPE: FLOAT64]")

		// REFLECTION CASE:  REFLECT TYPE OF RESULT MATCHES REFLECT TYPE OF KNOWN INT:
		case reflect.TypeOf(resultInt):

			// POP RESULT AS INT AND DISPLAY TO OUTPUT WITH TYPE
			resultInt = operandStack.Pop().(int)
			fmt.Print(resultInt)
			fmt.Println("  [TYPE: INT]")
		}
	}
}
