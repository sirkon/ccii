// Code generated from CCII.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // CCII

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CCIIParser.
type CCIIVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CCIIParser#main.
	VisitMain(ctx *MainContext) interface{}

	// Visit a parse tree produced by CCIIParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by CCIIParser#object_item.
	VisitObject_item(ctx *Object_itemContext) interface{}

	// Visit a parse tree produced by CCIIParser#object_intern.
	VisitObject_intern(ctx *Object_internContext) interface{}

	// Visit a parse tree produced by CCIIParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by CCIIParser#array_item.
	VisitArray_item(ctx *Array_itemContext) interface{}

	// Visit a parse tree produced by CCIIParser#array_intern.
	VisitArray_intern(ctx *Array_internContext) interface{}
}
