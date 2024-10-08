package ast

type PropertyKind string

const (
	PropertyKindValue  PropertyKind = "value"
	PropertyKindGet    PropertyKind = "get"
	PropertyKindSet    PropertyKind = "set"
	PropertyKindMethod PropertyKind = "method"
)

type (
	ClassLiteral struct {
		Class      Idx
		RightBrace Idx
		Name       *Identifier
		SuperClass *Expression
		Body       []ClassElement
	}

	Property interface {
		Expr
		_property()
	}

	PropertyShort struct {
		Name        *Identifier
		Initializer *Expression
	}

	PropertyKeyed struct {
		Key      *Expression
		Kind     PropertyKind
		Value    *Expression
		Computed bool
	}

	ClassElement interface {
		_classElement()
	}

	FieldDefinition struct {
		Idx         Idx
		Key         *Expression
		Initializer *Expression
		Computed    bool
		Static      bool
	}

	MethodDefinition struct {
		Idx      Idx
		Key      *Expression
		Kind     PropertyKind // "method", "get" or "set"
		Body     *FunctionLiteral
		Computed bool
		Static   bool
	}

	ClassStaticBlock struct {
		Static Idx
		Block  *BlockStatement
	}
)

func (*ClassLiteral) _expr()  {}
func (*PropertyShort) _expr() {}
func (*PropertyKeyed) _expr() {}

func (*PropertyShort) _property() {}
func (*PropertyKeyed) _property() {}
func (*SpreadElement) _property() {}

func (*FieldDefinition) _classElement()  {}
func (*MethodDefinition) _classElement() {}
func (*ClassStaticBlock) _classElement() {}
