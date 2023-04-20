package main

import (
	"strings"
	"unicode"
)

type TokenType int

const (
	EOF TokenType = iota
	Identifier
	TypeKeyword
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	Colon
	Bang
	AngleBracketLeft
	AngleBracketRight
)

type Token struct {
	Type  TokenType
	Value string
}

func lex(input string) []Token {
	var tokens []Token
	r := strings.NewReader(input)

	for {
		ch, _, err := r.ReadRune()
		if err != nil {
			break
		}

		if unicode.IsSpace(ch) {
			continue
		}

		switch ch {
		case '(':
			tokens = append(tokens, Token{Type: LeftParen})
		case ')':
			tokens = append(tokens, Token{Type: RightParen})
		case '{':
			tokens = append(tokens, Token{Type: LeftBrace})
		case '}':
			tokens = append(tokens, Token{Type: RightBrace})
		case ':':
			tokens = append(tokens, Token{Type: Colon})
		case '!':
			tokens = append(tokens, Token{Type: Bang})
		case '<':
			tokens = append(tokens, Token{Type: AngleBracketLeft})
		case '>':
			tokens = append(tokens, Token{Type: AngleBracketRight})
		default:
			if unicode.IsLetter(ch) {
				r.UnreadRune()
				identifier := ""
				for {
					letter, _, err := r.ReadRune()
					if err != nil || !unicode.IsLetter(letter) {
						r.UnreadRune()
						break
					}
					identifier += string(letter)
				}

				tokenType := Identifier
				if strings.ToLower(identifier) == "type" {
					tokenType = TypeKeyword
				}
				tokens = append(tokens, Token{Type: tokenType, Value: identifier})
			}
		}
	}

	tokens = append(tokens, Token{Type: EOF})
	return tokens
}
