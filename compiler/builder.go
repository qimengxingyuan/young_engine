package compiler

import (
	"errors"
	"fmt"
	"github.com/qimengxingyuan/young_engine/executor"
	"github.com/qimengxingyuan/young_engine/token"
)

// planner
type planner func(builder *Builder, curPre *precedence) (*executor.Node, error)

func planValue(builder *Builder, curPre *precedence) (*executor.Node, error) {
	if !builder.parser.hasNext() {
		return nil, nil
	}

	tok := builder.parser.next()
	switch tok.Kind {
	case token.OpenParen:
		// 最高优先级
		ret, err := builder.Build()
		if err != nil {
			return nil, err
		}

		// advance past the CloseParen token. We know that it's a CLAUSE_CLOSE, because at parse-time we check for unbalanced parens.
		builder.parser.next()

		// the stage we got represents all of the logic contained within the parens
		// but for technical reasons, we need to wrap this stage in a "noop" stage which breaks long chains of precedence.
		// see github #33.
		node := executor.NewNode(nil, ret, executor.NOOP, nil)
		return node, nil
	case token.Identifier:
		node := executor.NewNode(nil, nil, executor.VALUE, tok.Value)
		return node, nil
	case token.IntegerLiteral:
		node := executor.NewNodeWithType(nil, nil, executor.LITERAL, tok.Value, executor.TypeInteger)
		return node, nil
	case token.FloatLiteral:
		node := executor.NewNodeWithType(nil, nil, executor.LITERAL, tok.Value, executor.TypeFloat)
		return node, nil
	case token.BoolLiteral:
		node := executor.NewNodeWithType(nil, nil, executor.LITERAL, tok.Value, executor.TypeBool)
		return node, nil
	case token.StringLiteral:
		node := executor.NewNodeWithType(nil, nil, executor.LITERAL, tok.Value, executor.TypeString)
		return node, nil
	case token.Subtraction:
		ret, err := curPre.plan(builder)
		if err != nil {
			return nil, err
		}
		node := executor.NewNodeWithPrefixFix(ret, executor.NEGATIVE, nil)
		return node, nil
	case token.Addition: // token.Not,
		ret, err := curPre.plan(builder)
		if err != nil {
			return nil, err
		}
		node := executor.NewNodeWithPrefixFix(ret, executor.POSITIVE, nil)
		return node, nil
	case token.Not:
		builder.parser.rewind()
		return nil, nil
	default:
		errorMsg := fmt.Sprintf("Unable to plan token kind: '%s', value: '%v'", tok.Kind.String(), tok.Value)
		return nil, errors.New(errorMsg)
	}
}

type precedence struct {
	validKindsToSymbols map[token.Kind]executor.Symbol // 当前优先级的token类型
	nextPrecedence      *precedence                    // 更高优先级的
	planner             planner
}

var (
	// * / %
	multiplicative = &precedence{
		validKindsToSymbols: executor.MultiKindsToSymbol,
		planner:             planValue,
	}

	// + -
	additive = &precedence{
		validKindsToSymbols: executor.AddKindsToSymbol,
		nextPrecedence:      multiplicative,
	}

	// > >= < <= == !=
	comparator = &precedence{
		validKindsToSymbols: executor.CompareKindsToSymbol,
		nextPrecedence:      additive,
	}

	//!
	logicalNot = &precedence{
		validKindsToSymbols: executor.NotKindsToSymbol,
		nextPrecedence:      comparator,
	}

	// &&
	logicalAnd = &precedence{
		validKindsToSymbols: executor.AndKindsToSymbol,
		nextPrecedence:      logicalNot,
	}

	// ||
	logicalOr = &precedence{
		validKindsToSymbols: executor.OrKindsToSymbol,
		nextPrecedence:      logicalAnd,
	}

	lowestPrecedence = logicalOr
)

func (p *precedence) plan(builder *Builder) (*executor.Node, error) {
	var err error
	var leftNode, rightNode *executor.Node

	if p.nextPrecedence != nil {
		leftNode, err = p.nextPrecedence.plan(builder)
		if err != nil {
			return nil, err
		}
	} else if p.planner != nil {
		leftNode, err = p.planner(builder, p)
		if err != nil {
			return nil, err
		}
	}

	for builder.parser.hasNext() {
		tok := builder.parser.next()

		if tok.Kind.IsEof() {
			break
		}

		symbol, exist := p.validKindsToSymbols[tok.Kind]
		if !exist {
			break
		}

		rightNode, err = p.plan(builder)
		if err != nil {
			return nil, err
		}

		node := executor.NewNode(leftNode, rightNode, symbol, nil)
		return node, nil
	}
	builder.parser.rewind()
	return leftNode, nil
}

type Builder struct {
	rootPlanner *precedence
	parser      *Parser
}

func NewBuilder(p *Parser) *Builder {
	return &Builder{
		rootPlanner: lowestPrecedence,
		parser:      p,
	}
}

func (b *Builder) Build() (*executor.Node, error) {
	if b.parser == nil {
		return nil, errors.New("parse is nil")
	}

	if b.rootPlanner != nil {
		// TODO 树优化
		return b.rootPlanner.plan(b)
	}

	return nil, errors.New("build failed")
}
