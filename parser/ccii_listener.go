// Code generated from CCII.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // CCII

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CCIIListener is a complete listener for a parse tree produced by CCIIParser.
type CCIIListener interface {
	antlr.ParseTreeListener

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterObject_item is called when entering the object_item production.
	EnterObject_item(c *Object_itemContext)

	// EnterObject_intern is called when entering the object_intern production.
	EnterObject_intern(c *Object_internContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterArray_item is called when entering the array_item production.
	EnterArray_item(c *Array_itemContext)

	// EnterArray_intern is called when entering the array_intern production.
	EnterArray_intern(c *Array_internContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitObject_item is called when exiting the object_item production.
	ExitObject_item(c *Object_itemContext)

	// ExitObject_intern is called when exiting the object_intern production.
	ExitObject_intern(c *Object_internContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitArray_item is called when exiting the array_item production.
	ExitArray_item(c *Array_itemContext)

	// ExitArray_intern is called when exiting the array_intern production.
	ExitArray_intern(c *Array_internContext)
}
