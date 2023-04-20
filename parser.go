package main

import (
	"errors"
	"fmt"
)

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) Parse() (*TypeDef, error) {
	typeDef := &TypeDef{}

	if p.expect(TypeKeyword) {
		p.consume()

		if !p.expect(Identifier) {
			return nil, errors.New("expected a type name")
		}

		typeDef.TypeName = p.current().Value
		p.consume()

		if p.expect(LeftParen) {
			p.consume()

			if !p.expect(Identifier) {
				return nil, errors.New("expected a tag value")
			}

			typeDef.Tag = p.current().Value
			p.consume()

			if !p.expect(RightParen) {
				return nil, errors.New("expected closing parenthesis")
			}
			p.consume()
		}

		if !p.expect(LeftBrace) {
			return nil, errors.New("expected opening brace")
		}
		p.consume()

		for !p.expect(RightBrace) {
			field, err := p.parseField()
			if err != nil {
				return nil, err
			}
			typeDef.Fields = append(typeDef.Fields, *field)
		}
		p.consume()
	}

	return typeDef, nil
}

func (p *Parser) parseField() (*Field, error) {
	field := &Field{}

	if !p.expect(Identifier) {
		return nil, errors.New("expected field name")
	}

	field.Name = p.current().Value
	p.consume()

	if !p.expect(Colon) {
		return nil, errors.New("expected colon")
	}
	p.consume()

	if !p.expect(Identifier) && !p.expect(AngleBracketLeft) {
		return nil, errors.New("expected field type")
	}

	field.Type = p.current().Value
	p.consume()

	if p.current().Type == AngleBracketLeft {
		p.consume()

		if !p.expect(Identifier) {
			return nil, errors.New("expected array type")
		}

		arrayType := p.current().Value
		field.Type = fmt.Sprintf("array<%s>", arrayType)
		p.consume()

		if !p.expect(AngleBracketRight) {
			return nil, errors.New("expected closing angle bracket")
		}
		p.consume()
	}

	if p.expect(Bang) {
		p.consume()
		field.Nullable = true
	}

	return field, nil
}

func (p *Parser) current() Token {
	return p.tokens[p.pos]
}

func (p *Parser) consume() {
	p.pos++
}

func (p *Parser) expect(tokenType TokenType) bool {
	return p.current().Type == tokenType
}
