// Code generated from CCII.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // CCII

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCCIIVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCCIIVisitor) VisitMain(ctx *MainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCCIIVisitor) VisitObject(ctx *ObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCCIIVisitor) VisitObject_item(ctx *Object_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCCIIVisitor) VisitObject_intern(ctx *Object_internContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCCIIVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCCIIVisitor) VisitArray_item(ctx *Array_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCCIIVisitor) VisitArray_intern(ctx *Array_internContext) interface{} {
	return v.VisitChildren(ctx)
}
