package token

type TokenType int

const (
	ILLEGAL TokenType = (iota - 1) // Illegal token
	EOF

	IDENT //identifier
	INT   //int literal
	FLOAT //float literal

	EQ         // ==
	NEQ        // !=
	MATCH      // =~
	NOTMATCH   // !~
	ASSIGN     // =
	PLUS       // +
	PLUS_A     // += (PLUS ASSIGN)
	MINUS      // -
	MINUS_A    // -= (MINUS ASSIGN)
	BANG       // !
	ASTERISK   // *
	ASTERISK_A // *= (ASTERISK ASSIGN)
	SLASH      // /  divide
	SLASH_A    // /= (SLASH ASSIGN)
	MOD        // %
	MOD_A      // %= //MOD ASSIGN
	POWER      // ** (POWER)
	QUESTIONM  // ? (QUESTION MARK)

	LT        // <
	LE        // <=
	SHIFT_L   // << (SHIFT LEFT)
	GT        // >
	GE        // >=
	SHIFT_R   // >> (SHIFT RIGHT)
	COMMA     // ,
	SEMICOLON // ;

	LPAREN   // (
	RPAREN   // )
	LBRACE   // {
	RBRACE   // }
	LBRACKET // [
	RBRACKET // ]
	COLON    // :
	DOT      // .
	DOTDOT   // ..  (PARTIAL IMPLEMENTED, ONLY SUPPORT INTEGER/SingleString RANGE, AND ONLY USED IN 'FOR X IN A..B {}' )
	ELLIPSIS //... Function Variadic parameters
	PIPE     // |>
	THINARROW // ->
	FATARROW // =>

	INCREMENT // ++
	DECREMENT // --

	BITAND   // &
	BITOR    // |
	BITXOR   // ^
	BITAND_A // &=
	BITOR_A  // |=
	BITXOR_A // ^=

	CONDAND // &&
	CONDOR  // ||

	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELIF
	ELSIF
	ELSEIF
	ELSE
	RETURN
	INCLUDE
	STRING
	ISTRING
	BYTES
	AND
	OR
	STRUCT
	DO
	WHILE
	BREAK
	CONTINUE

	COMMENT // # (single-line comment)
	REGEX   // REGEX
	FOR
	IN
	WHERE
	GREP
	MAP
	CASE
	IS
	TRY
	CATCH
	FINALLY
	THROW
	DEFER
	SPAWN
	YIELD // NOT IMPLEMENTED
	NIL
	ENUM
	QW
	UNLESS
)

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"elif":     ELIF,
	"elsif":    ELSIF,
	"elseif":   ELSEIF,
	"else":     ELSE,
	"return":   RETURN,
	"include":  INCLUDE,
	"and":      AND,
	"or":       OR,
	"struct":   STRUCT,
	"do":       DO,
	"while":    WHILE,
	"break":    BREAK,
	"continue": CONTINUE,
	"for":      FOR,
	"in":       IN,
	"where":    WHERE,
	"grep":     GREP,
	"map":      MAP,
	"case":     CASE,
	"is":       IS,
	"try":      TRY,
	"catch":    CATCH,
	"finally":  FINALLY,
	"throw":    THROW,
	"defer":    DEFER,
	"spawn":    SPAWN,
	"yield":    YIELD,
	"nil":      NIL,
	"enum":     ENUM,
	"qw":       QW, //“quoted words”
	"unless":   UNLESS,
}

//for debug & testing
func (tt TokenType) String() string {
	switch tt {
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENT"
	case INT:
		return "INT"
	case FLOAT:
		return "FLOAT"
	case EQ:
		return "=="
	case NEQ:
		return "!="
	case MATCH:
		return "=~"
	case NOTMATCH:
		return "!~"
	case ASSIGN:
		return "="
	case PLUS:
		return "+"
	case PLUS_A:
		return "+="
	case MINUS:
		return "-"
	case MINUS_A:
		return "-="
	case BANG:
		return "!"
	case ASTERISK:
		return "*"
	case ASTERISK_A:
		return "*="
	case SLASH:
		return "/"
	case SLASH_A:
		return "/="
	case MOD:
		return "%"
	case MOD_A:
		return "%="
	case POWER:
		return "**"
	case LT:
		return "<"
	case LE:
		return "<="
	case SHIFT_L:
		return "<<"
	case GT:
		return ">"
	case GE:
		return ">="
	case SHIFT_R:
		return ">>"
	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case LBRACKET:
		return "["
	case RBRACKET:
		return "]"
	case COLON:
		return ":"
	case DOT:
		return "."
	case DOTDOT:
		return ".."
	case ELLIPSIS:
		return "..."
	case PIPE:
		return "|>"
	case THINARROW:
		return "->"
	case FATARROW:
		return "=>"
	case INCREMENT:
		return "++"
	case DECREMENT:
		return "--"
	case BITAND:
		return "&"
	case BITOR:
		return "|"
	case BITXOR:
		return "^"
	case BITAND_A:
		return "&="
	case BITOR_A:
		return "|="
	case BITXOR_A:
		return "^="
	case CONDAND:
		return "&&"
	case CONDOR:
		return "||"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case IF:
		return "IF"
	case ELIF:
		return "elif"
	case ELSIF:
		return "ELSIF"
	case ELSEIF:
		return "ELSEIF"
	case ELSE:
		return "ELSE"
	case RETURN:
		return "RETURN"
	case INCLUDE:
		return "INCLUDE"
	case STRING:
		return "STRING"
	case ISTRING:
		return "ISTRING"
	case BYTES:
		return "BYTES"
	case AND:
		return "AND"
	case OR:
		return "OR"
	case STRUCT:
		return "STRUCT"
	case DO:
		return "DO"
	case WHILE:
		return "WHILE"
	case BREAK:
		return "BREAK"
	case CONTINUE:
		return "CONTINUE"
	case COMMENT:
		return "#"
	case REGEX:
		return "/"
	case FOR:
		return "FOR"
	case IN:
		return "IN"
	case WHERE:
		return "WHERE"
	case GREP:
		return "GREP"
	case MAP:
		return "MAP"
	case CASE:
		return "CASE"
	case IS:
		return "IS"
	case TRY:
		return "TRY"
	case CATCH:
		return "CATCH"
	case FINALLY:
		return "FINALLY"
	case THROW:
		return "THROW"
	case QUESTIONM:
		return "?"
	case DEFER:
		return "DEFER"
	case NIL:
		return "nil"
	case ENUM:
		return "enum"
	case QW:
		return "qw"
	case UNLESS:
		return "unless"
	default:
		return "unknown"
	}
}

type Token struct {
	Pos     Position
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
