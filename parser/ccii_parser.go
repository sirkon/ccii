// Code generated from CCII.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // CCII

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 64, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 5, 2, 20, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 33, 10, 4, 3, 5, 3, 5, 3,
	5, 7, 5, 38, 10, 5, 12, 5, 14, 5, 41, 11, 5, 5, 5, 43, 10, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 5, 7, 52, 10, 7, 3, 8, 3, 8, 3, 8, 7,
	8, 57, 10, 8, 12, 8, 14, 8, 60, 11, 8, 5, 8, 62, 10, 8, 3, 8, 2, 2, 9,
	2, 4, 6, 8, 10, 12, 14, 2, 2, 2, 66, 2, 19, 3, 2, 2, 2, 4, 21, 3, 2, 2,
	2, 6, 32, 3, 2, 2, 2, 8, 42, 3, 2, 2, 2, 10, 44, 3, 2, 2, 2, 12, 51, 3,
	2, 2, 2, 14, 61, 3, 2, 2, 2, 16, 20, 5, 4, 3, 2, 17, 20, 5, 10, 6, 2, 18,
	20, 7, 9, 2, 2, 19, 16, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 18, 3, 2, 2,
	2, 20, 3, 3, 2, 2, 2, 21, 22, 7, 3, 2, 2, 22, 23, 5, 8, 5, 2, 23, 24, 7,
	4, 2, 2, 24, 5, 3, 2, 2, 2, 25, 26, 7, 9, 2, 2, 26, 27, 7, 5, 2, 2, 27,
	33, 7, 9, 2, 2, 28, 29, 7, 9, 2, 2, 29, 33, 5, 4, 3, 2, 30, 31, 7, 9, 2,
	2, 31, 33, 5, 10, 6, 2, 32, 25, 3, 2, 2, 2, 32, 28, 3, 2, 2, 2, 32, 30,
	3, 2, 2, 2, 33, 7, 3, 2, 2, 2, 34, 39, 5, 6, 4, 2, 35, 36, 7, 6, 2, 2,
	36, 38, 5, 6, 4, 2, 37, 35, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3,
	2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42,
	34, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 9, 3, 2, 2, 2, 44, 45, 7, 7, 2,
	2, 45, 46, 5, 14, 8, 2, 46, 47, 7, 8, 2, 2, 47, 11, 3, 2, 2, 2, 48, 52,
	5, 4, 3, 2, 49, 52, 5, 10, 6, 2, 50, 52, 7, 9, 2, 2, 51, 48, 3, 2, 2, 2,
	51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 13, 3, 2, 2, 2, 53, 58, 5,
	12, 7, 2, 54, 55, 7, 6, 2, 2, 55, 57, 5, 12, 7, 2, 56, 54, 3, 2, 2, 2,
	57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 62, 3,
	2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 53, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62,
	15, 3, 2, 2, 2, 9, 19, 32, 39, 42, 51, 58, 61,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'='", "','", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "SCALAR",
}

var ruleNames = []string{
	"main", "object", "object_item", "object_intern", "array", "array_item",
	"array_intern",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CCIIParser struct {
	*antlr.BaseParser
}

func NewCCIIParser(input antlr.TokenStream) *CCIIParser {
	this := new(CCIIParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CCII.g4"

	return this
}

// CCIIParser tokens.
const (
	CCIIParserEOF    = antlr.TokenEOF
	CCIIParserT__0   = 1
	CCIIParserT__1   = 2
	CCIIParserT__2   = 3
	CCIIParserT__3   = 4
	CCIIParserT__4   = 5
	CCIIParserT__5   = 6
	CCIIParserSCALAR = 7
)

// CCIIParser rules.
const (
	CCIIParserRULE_main          = 0
	CCIIParserRULE_object        = 1
	CCIIParserRULE_object_item   = 2
	CCIIParserRULE_object_intern = 3
	CCIIParserRULE_array         = 4
	CCIIParserRULE_array_item    = 5
	CCIIParserRULE_array_intern  = 6
)

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CCIIParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CCIIParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *MainContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *MainContext) SCALAR() antlr.TerminalNode {
	return s.GetToken(CCIIParserSCALAR, 0)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.ExitMain(s)
	}
}

func (s *MainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CCIIVisitor:
		return t.VisitMain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CCIIParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CCIIParserRULE_main)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CCIIParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.Object()
		}

	case CCIIParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(15)
			p.Array()
		}

	case CCIIParserSCALAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(16)
			p.Match(CCIIParserSCALAR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CCIIParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CCIIParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) Object_intern() IObject_internContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_internContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_internContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.ExitObject(s)
	}
}

func (s *ObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CCIIVisitor:
		return t.VisitObject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CCIIParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CCIIParserRULE_object)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Match(CCIIParserT__0)
	}
	{
		p.SetState(20)
		p.Object_intern()
	}
	{
		p.SetState(21)
		p.Match(CCIIParserT__1)
	}

	return localctx
}

// IObject_itemContext is an interface to support dynamic dispatch.
type IObject_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_itemContext differentiates from other interfaces.
	IsObject_itemContext()
}

type Object_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_itemContext() *Object_itemContext {
	var p = new(Object_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CCIIParserRULE_object_item
	return p
}

func (*Object_itemContext) IsObject_itemContext() {}

func NewObject_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_itemContext {
	var p = new(Object_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CCIIParserRULE_object_item

	return p
}

func (s *Object_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_itemContext) AllSCALAR() []antlr.TerminalNode {
	return s.GetTokens(CCIIParserSCALAR)
}

func (s *Object_itemContext) SCALAR(i int) antlr.TerminalNode {
	return s.GetToken(CCIIParserSCALAR, i)
}

func (s *Object_itemContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *Object_itemContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *Object_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.EnterObject_item(s)
	}
}

func (s *Object_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.ExitObject_item(s)
	}
}

func (s *Object_itemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CCIIVisitor:
		return t.VisitObject_item(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CCIIParser) Object_item() (localctx IObject_itemContext) {
	localctx = NewObject_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CCIIParserRULE_object_item)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Match(CCIIParserSCALAR)
		}
		{
			p.SetState(24)
			p.Match(CCIIParserT__2)
		}
		{
			p.SetState(25)
			p.Match(CCIIParserSCALAR)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
			p.Match(CCIIParserSCALAR)
		}
		{
			p.SetState(27)
			p.Object()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(28)
			p.Match(CCIIParserSCALAR)
		}
		{
			p.SetState(29)
			p.Array()
		}

	}

	return localctx
}

// IObject_internContext is an interface to support dynamic dispatch.
type IObject_internContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_internContext differentiates from other interfaces.
	IsObject_internContext()
}

type Object_internContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_internContext() *Object_internContext {
	var p = new(Object_internContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CCIIParserRULE_object_intern
	return p
}

func (*Object_internContext) IsObject_internContext() {}

func NewObject_internContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_internContext {
	var p = new(Object_internContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CCIIParserRULE_object_intern

	return p
}

func (s *Object_internContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_internContext) AllObject_item() []IObject_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObject_itemContext)(nil)).Elem())
	var tst = make([]IObject_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObject_itemContext)
		}
	}

	return tst
}

func (s *Object_internContext) Object_item(i int) IObject_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObject_itemContext)
}

func (s *Object_internContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_internContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_internContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.EnterObject_intern(s)
	}
}

func (s *Object_internContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.ExitObject_intern(s)
	}
}

func (s *Object_internContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CCIIVisitor:
		return t.VisitObject_intern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CCIIParser) Object_intern() (localctx IObject_internContext) {
	localctx = NewObject_internContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CCIIParserRULE_object_intern)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CCIIParserSCALAR {
		{
			p.SetState(32)
			p.Object_item()
		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CCIIParserT__3 {
			{
				p.SetState(33)
				p.Match(CCIIParserT__3)
			}
			{
				p.SetState(34)
				p.Object_item()
			}

			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CCIIParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CCIIParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) Array_intern() IArray_internContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_internContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_internContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CCIIVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CCIIParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CCIIParserRULE_array)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(CCIIParserT__4)
	}
	{
		p.SetState(43)
		p.Array_intern()
	}
	{
		p.SetState(44)
		p.Match(CCIIParserT__5)
	}

	return localctx
}

// IArray_itemContext is an interface to support dynamic dispatch.
type IArray_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_itemContext differentiates from other interfaces.
	IsArray_itemContext()
}

type Array_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_itemContext() *Array_itemContext {
	var p = new(Array_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CCIIParserRULE_array_item
	return p
}

func (*Array_itemContext) IsArray_itemContext() {}

func NewArray_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_itemContext {
	var p = new(Array_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CCIIParserRULE_array_item

	return p
}

func (s *Array_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_itemContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *Array_itemContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *Array_itemContext) SCALAR() antlr.TerminalNode {
	return s.GetToken(CCIIParserSCALAR, 0)
}

func (s *Array_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.EnterArray_item(s)
	}
}

func (s *Array_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.ExitArray_item(s)
	}
}

func (s *Array_itemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CCIIVisitor:
		return t.VisitArray_item(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CCIIParser) Array_item() (localctx IArray_itemContext) {
	localctx = NewArray_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CCIIParserRULE_array_item)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CCIIParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Object()
		}

	case CCIIParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Array()
		}

	case CCIIParserSCALAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.Match(CCIIParserSCALAR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArray_internContext is an interface to support dynamic dispatch.
type IArray_internContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_internContext differentiates from other interfaces.
	IsArray_internContext()
}

type Array_internContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_internContext() *Array_internContext {
	var p = new(Array_internContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CCIIParserRULE_array_intern
	return p
}

func (*Array_internContext) IsArray_internContext() {}

func NewArray_internContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_internContext {
	var p = new(Array_internContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CCIIParserRULE_array_intern

	return p
}

func (s *Array_internContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_internContext) AllArray_item() []IArray_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArray_itemContext)(nil)).Elem())
	var tst = make([]IArray_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArray_itemContext)
		}
	}

	return tst
}

func (s *Array_internContext) Array_item(i int) IArray_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArray_itemContext)
}

func (s *Array_internContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_internContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_internContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.EnterArray_intern(s)
	}
}

func (s *Array_internContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CCIIListener); ok {
		listenerT.ExitArray_intern(s)
	}
}

func (s *Array_internContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CCIIVisitor:
		return t.VisitArray_intern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CCIIParser) Array_intern() (localctx IArray_internContext) {
	localctx = NewArray_internContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CCIIParserRULE_array_intern)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CCIIParserT__0)|(1<<CCIIParserT__4)|(1<<CCIIParserSCALAR))) != 0 {
		{
			p.SetState(51)
			p.Array_item()
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CCIIParserT__3 {
			{
				p.SetState(52)
				p.Match(CCIIParserT__3)
			}
			{
				p.SetState(53)
				p.Array_item()
			}

			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}
