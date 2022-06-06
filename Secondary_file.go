package main

//The term pass by reference means that there is generally
// a value of this variable and all operations refer to this variable.
// All operations know where this value is.

//This file is written for instruction only

// an example of pass by Reference :

import "fmt"

var outer = "A"

// This function receives a
// pointer refering to a value of type string. 
// In fact, this address points to a string.
func MF(inner *string){

	*inner = "B"

  // Inside the function of existence * and inner means go to the value of inner.
  // inner is an address. Change the value of inner into memory and put B.
}

// Finally we call the function.

// Existence & means create an address from the outer variable and give it an address.

MF(&outer)

// output :
// value of inner => "B"
// value of outer => "B"
