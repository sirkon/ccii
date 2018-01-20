// Code generated from CCII.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // CCII

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCCIIListener is a complete listener for a parse tree produced by CCIIParser.
type BaseCCIIListener struct{}

var _ CCIIListener = &BaseCCIIListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCCIIListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCCIIListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCCIIListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCCIIListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMain is called when production main is entered.
func (s *BaseCCIIListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseCCIIListener) ExitMain(ctx *MainContext) {}

// EnterObject is called when production object is entered.
func (s *BaseCCIIListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseCCIIListener) ExitObject(ctx *ObjectContext) {}

// EnterObject_item is called when production object_item is entered.
func (s *BaseCCIIListener) EnterObject_item(ctx *Object_itemContext) {}

// ExitObject_item is called when production object_item is exited.
func (s *BaseCCIIListener) ExitObject_item(ctx *Object_itemContext) {}

// EnterObject_intern is called when production object_intern is entered.
func (s *BaseCCIIListener) EnterObject_intern(ctx *Object_internContext) {}

// ExitObject_intern is called when production object_intern is exited.
func (s *BaseCCIIListener) ExitObject_intern(ctx *Object_internContext) {}

// EnterArray is called when production array is entered.
func (s *BaseCCIIListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseCCIIListener) ExitArray(ctx *ArrayContext) {}

// EnterArray_item is called when production array_item is entered.
func (s *BaseCCIIListener) EnterArray_item(ctx *Array_itemContext) {}

// ExitArray_item is called when production array_item is exited.
func (s *BaseCCIIListener) ExitArray_item(ctx *Array_itemContext) {}

// EnterArray_intern is called when production array_intern is entered.
func (s *BaseCCIIListener) EnterArray_intern(ctx *Array_internContext) {}

// ExitArray_intern is called when production array_intern is exited.
func (s *BaseCCIIListener) ExitArray_intern(ctx *Array_internContext) {}
