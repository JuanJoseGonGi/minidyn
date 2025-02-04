package language

import (
	"testing"
)

func TestConditionalExpression(t *testing.T) {
	es := ConditionalExpression{
		Token: Token{Type: IDENT, Literal: "a"},
	}

	tl := es.TokenLiteral()
	if tl != "a" {
		t.Fatalf("wrong token literal. expected=%q, got=%q", "a", tl)
	}

	es.statementNode()

	if es.String() != "" {
		t.Fatalf("empty expression expected ")
	}
}

func TestIdentifier(t *testing.T) {
	es := Identifier{
		Token: Token{Type: IDENT, Literal: "a"},
		Value: "a",
	}

	tl := es.TokenLiteral()
	if tl != "a" {
		t.Fatalf("wrong token literal. expected=%q, got=%q", "a", tl)
	}

	es.expressionNode()
}

func TestPrefixExpression(t *testing.T) {
	es := PrefixExpression{
		Token:    Token{Type: NOT, Literal: "NOT"},
		Operator: NOT,
		Right: &Identifier{
			Token: Token{Type: IDENT, Literal: "a"},
			Value: "a",
		},
	}

	tl := es.TokenLiteral()
	if tl != "NOT" {
		t.Fatalf("wrong token literal. expected=%q, got=%q", "NOT", tl)
	}

	es.expressionNode()
}

func TestInfixExpression(t *testing.T) {
	ie := InfixExpression{
		Token:    Token{Type: EQ, Literal: "="},
		Operator: "=",
		Left: &Identifier{
			Token: Token{Type: IDENT, Literal: "a"},
			Value: "a",
		},
		Right: &Identifier{
			Token: Token{Type: IDENT, Literal: "b"},
			Value: "b",
		},
	}

	ie.expressionNode()

	tl := ie.TokenLiteral()
	if tl != "=" {
		t.Fatalf("wrong token literal. expected=%q, got=%q", "NOT", tl)
	}
}

func TestIndexExpression(t *testing.T) {
	ie := IndexExpression{
		Token: Token{Type: LBRACKET, Literal: "["},
		Left: &Identifier{
			Token: Token{Type: IDENT, Literal: "a"},
			Value: "a",
		},
		Index: &Identifier{
			Token: Token{Type: IDENT, Literal: ":i"},
			Value: ":i",
		},
	}

	tl := ie.TokenLiteral()
	if tl != "[" {
		t.Fatalf("wrong token literal. expected=%q, got=%q", "[", tl)
	}

	ie.expressionNode()
}

func TestCallExpression(t *testing.T) {
	ce := CallExpression{
		Token: Token{Type: LPAREN, Literal: "("},
		Function: &Identifier{
			Token: Token{Type: IDENT, Literal: "size"},
			Value: "size",
		},
		Arguments: []Expression{
			&Identifier{
				Token: Token{Type: IDENT, Literal: "a"},
				Value: "a",
			},
		},
	}

	tl := ce.TokenLiteral()
	if tl != "(" {
		t.Fatalf("wrong token literal. expected=%q, got=%q", "NOT", tl)
	}

	ce.expressionNode()
}

func TestBetweenExpression(t *testing.T) {
	be := BetweenExpression{
		Token: Token{Type: LPAREN, Literal: "("},
		Left: &Identifier{
			Token: Token{Type: IDENT, Literal: "b"},
			Value: "b",
		},
		Range: [2]Expression{
			&Identifier{
				Token: Token{Type: IDENT, Literal: "a"},
				Value: "a",
			},
			&Identifier{
				Token: Token{Type: IDENT, Literal: "b"},
				Value: "b",
			},
		},
	}

	tl := be.TokenLiteral()
	if tl != "(" {
		t.Fatalf("wrong token literal. expected=%q, got=%q", "NOT", tl)
	}

	if be.String() != "b BETWEEN a AND b" {
		t.Fatalf("wrong string representation. expected=%q, got=%q", "b BETWEEN a AND b", be.String())
	}

	be.expressionNode()
}

func TestInExpression(t *testing.T) {
	ine := InExpression{
		Token: Token{Type: LPAREN, Literal: "("},
		Left: &Identifier{
			Token: Token{Type: IDENT, Literal: "b"},
			Value: "b",
		},
		Range: []Expression{
			&Identifier{
				Token: Token{Type: IDENT, Literal: "a"},
				Value: "a",
			},
			&Identifier{
				Token: Token{Type: IDENT, Literal: "b"},
				Value: "b",
			},
		},
	}

	tl := ine.TokenLiteral()
	if tl != "(" {
		t.Fatalf("wrong token literal. expected=%q, got=%q", "NOT", tl)
	}

	if ine.String() != "(b IN (a, b))" {
		t.Fatalf("wrong string representation. expected=%q, got=%q", "b IN (a, b)", ine.String())
	}

	ine.expressionNode()
}

func TestUpdateExpression(t *testing.T) {
	es := UpdateExpression{
		Token: Token{Type: IDENT, Literal: "SET"},
	}

	tl := es.TokenLiteral()
	if tl != "SET" {
		t.Fatalf("wrong token literal. expected=%q, got=%q", "a", tl)
	}

	if es.String() != "()" {
		t.Fatalf("unexpected expression representation. expected=%q, got=%q", "SET ()", es.String())
	}

	es.expressionNode()
}

func TestActionExpression(t *testing.T) {
	ae := &ActionExpression{
		Token: Token{Type: DELETE, Literal: "SET"},
		Left:  &Identifier{Value: ":x", Token: Token{Type: IDENT, Literal: ":x"}},
		Right: &Identifier{Value: ":x", Token: Token{Type: IDENT, Literal: ":x"}},
	}

	tl := ae.TokenLiteral()
	if tl != "SET" {
		t.Fatalf("wrong token literal. expected=%q, got=%q", "SET", tl)
	}

	if ae.String() != "SET (:x, :x)" {
		t.Fatalf("wrong string representation. expected=%q, got=%q", "SET (:x, :x)", ae.String())
	}

	ae.expressionNode()
}

func TestUpdateStatement(t *testing.T) {
	es := UpdateStatement{
		Token: Token{Type: IDENT, Literal: "SET"},
	}

	tl := es.TokenLiteral()
	if tl != "SET" {
		t.Fatalf("wrong token literal. expected=%q, got=%q", "a", tl)
	}

	if es.String() != "" {
		t.Fatalf("empty expression expected ")
	}

	es.statementNode()
}

func BenchmarkCallExpression(b *testing.B) {
	ce := CallExpression{
		Token: Token{Type: LPAREN, Literal: "("},
		Function: &Identifier{
			Token: Token{Type: IDENT, Literal: "size"},
			Value: "size",
		},
		Arguments: []Expression{
			&Identifier{
				Token: Token{Type: IDENT, Literal: "a"},
				Value: "a",
			},
		},
	}

	for n := 0; n < b.N; n++ {
		if ce.String() != "size(a)" {
			b.Fatal("wrong call expression string")
		}
	}
}
