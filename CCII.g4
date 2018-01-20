grammar CCII;

main
    : object
    | array
    | SCALAR
    ;

object
    : '{' object_intern '}'
    ;

object_item
    : SCALAR '=' SCALAR
    | SCALAR object
    | SCALAR array
    ;

object_intern
    : (object_item (',' object_item)*)?
    ;

array
    : '[' array_intern ']'
    ;

array_item
    : object
    | array
    | SCALAR
    ;

array_intern
    : (array_item (',' array_item)*)?
    ;

SCALAR
    : ~( '=' | '{' | '}' | '[' | ']' | '\\' | ',' | ':' | '$' | '\'' | '"' | '&' | '>' | '<' )+
    ;
