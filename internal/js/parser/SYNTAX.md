# Syntax
## Statement
```text
PROGRAM     = STMT*
STMT        = EXPR ";"
            | // import
            | // class
            | "function" IDENT "(" FUNC_PARAMS ")" STMT
            | "return" EXPR? ";"
            | "if" "(" EXPR ")" STMT ("else" STMT)?
            | "while" "(" EXPR ")"STMT 
            | "for" "(" EXPR? ";" EXPR? ";" EXPR? ")"STMT 
            | "{" STMT* "}"
EXPR        = ASSIGN
ASSIGN      = ANDOR ("=" ANDOR)?
            | "var" IDENT ("=" ANDOR)?
            | "let" IDENT ("=" ANDOR)?
            | "const" IDENT "="ANDOR 
ANDOR       = EQUALITY ("&&" EQUALITY | "||" EQUALITY)*
EQUALITY    = RELATIONAL ("==" RELATIONAL | "!=" RELATIONAL)*
RELATIONAL  = ADD ("<" ADD | "<=" ADD | ">" ADD | ">=" ADD)*
ADD         = MUL ("+" MUL | "-" MUL)*
MUL         = UNARY ("*" UNARY | "/" UNARY | "%" UNARY)*
PRIMARY     = MEMBER
ACCESS      = LITERAL ("["EXPR"]" | "." LITERAL)*
LITERAL     = "(" expr ")"
            | IDENT
            | IDENT "(" CALL_ARGS ")"
            | STRING
            | NUMBER
            | BOOLEAN
            | ARRAY
            | OBJECT
            | NULL
            | SYMBOL
            | UNDEFINED

FUNC_PARAMS = IDENT ("," IDENT)*
CALL_ARGS   = EXPR ("," EXPR)*

IDENT       = [a-zA-Z_]+[a-zA-Z0-9_]*
```