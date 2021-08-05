package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

var src = `
package main
func main() {
	println("Hello, World!")
}
`

//0  *ast.File {
//1  .  Package: test.go:2:1
//2  .  Name: *ast.Ident {
//3  .  .  NamePos: test.go:2:9
//4  .  .  Name: "main"
//5  .  }
//6  .  Decls: []ast.Decl (len = 1) {
//7  .  .  0: *ast.FuncDecl {
//8  .  .  .  Name: *ast.Ident {
//9  .  .  .  .  NamePos: test.go:3:6
//10  .  .  .  .  Name: "main"
//11  .  .  .  .  Obj: *ast.Object {
//12  .  .  .  .  .  Kind: func
//13  .  .  .  .  .  Name: "main"
//14  .  .  .  .  .  Decl: *(obj @ 7)
//15  .  .  .  .  }
//16  .  .  .  }
//17  .  .  .  Type: *ast.FuncType {
//18  .  .  .  .  Func: test.go:3:1
//19  .  .  .  .  Params: *ast.FieldList {
//20  .  .  .  .  .  Opening: test.go:3:10
//21  .  .  .  .  .  Closing: test.go:3:11
//22  .  .  .  .  }
//23  .  .  .  }
//24  .  .  .  Body: *ast.BlockStmt {
//25  .  .  .  .  Lbrace: test.go:3:13
//26  .  .  .  .  List: []ast.Stmt (len = 1) {
//27  .  .  .  .  .  0: *ast.ExprStmt {
//28  .  .  .  .  .  .  X: *ast.CallExpr {
//29  .  .  .  .  .  .  .  Fun: *ast.Ident {
//30  .  .  .  .  .  .  .  .  NamePos: test.go:4:2
//31  .  .  .  .  .  .  .  .  Name: "println"
//32  .  .  .  .  .  .  .  }
//33  .  .  .  .  .  .  .  Lparen: test.go:4:9
//34  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
//35  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
//36  .  .  .  .  .  .  .  .  .  ValuePos: test.go:4:10
//37  .  .  .  .  .  .  .  .  .  Kind: STRING
//38  .  .  .  .  .  .  .  .  .  Value: "\"Hello, World!\""
//39  .  .  .  .  .  .  .  .  }
//40  .  .  .  .  .  .  .  }
//41  .  .  .  .  .  .  .  Ellipsis: -
//42  .  .  .  .  .  .  .  Rparen: test.go:4:25
//43  .  .  .  .  .  .  }
//44  .  .  .  .  .  }
//45  .  .  .  .  }
//46  .  .  .  .  Rbrace: test.go:5:1
//47  .  .  .  }
//48  .  .  }
//49  .  }
//50  .  Scope: *ast.Scope {
//51  .  .  Objects: map[string]*ast.Object (len = 1) {
//52  .  .  .  "main": *(obj @ 11)
//53  .  .  }
//54  .  }
//55  .  Unresolved: []*ast.Ident (len = 1) {
//56  .  .  0: *(obj @ 29)
//57  .  }
//58  }
//0  *ast.Ident {
//1  .  NamePos: test.go:2:9
//2  .  Name: "main"
//3  }
//0  nil
//0  *ast.FuncDecl {
//1  .  Name: *ast.Ident {
//2  .  .  NamePos: test.go:3:6
//3  .  .  Name: "main"
//4  .  .  Obj: *ast.Object {
//5  .  .  .  Kind: func
//6  .  .  .  Name: "main"
//7  .  .  .  Decl: *(obj @ 0)
//8  .  .  }
//9  .  }
//10  .  Type: *ast.FuncType {
//11  .  .  Func: test.go:3:1
//12  .  .  Params: *ast.FieldList {
//13  .  .  .  Opening: test.go:3:10
//14  .  .  .  Closing: test.go:3:11
//15  .  .  }
//16  .  }
//17  .  Body: *ast.BlockStmt {
//18  .  .  Lbrace: test.go:3:13
//19  .  .  List: []ast.Stmt (len = 1) {
//20  .  .  .  0: *ast.ExprStmt {
//21  .  .  .  .  X: *ast.CallExpr {
//22  .  .  .  .  .  Fun: *ast.Ident {
//23  .  .  .  .  .  .  NamePos: test.go:4:2
//24  .  .  .  .  .  .  Name: "println"
//25  .  .  .  .  .  }
//26  .  .  .  .  .  Lparen: test.go:4:9
//27  .  .  .  .  .  Args: []ast.Expr (len = 1) {
//28  .  .  .  .  .  .  0: *ast.BasicLit {
//29  .  .  .  .  .  .  .  ValuePos: test.go:4:10
//30  .  .  .  .  .  .  .  Kind: STRING
//31  .  .  .  .  .  .  .  Value: "\"Hello, World!\""
//32  .  .  .  .  .  .  }
//33  .  .  .  .  .  }
//34  .  .  .  .  .  Ellipsis: -
//35  .  .  .  .  .  Rparen: test.go:4:25
//36  .  .  .  .  }
//37  .  .  .  }
//38  .  .  }
//39  .  .  Rbrace: test.go:5:1
//40  .  }
//41  }
//0  *ast.Ident {
//1  .  NamePos: test.go:3:6
//2  .  Name: "main"
//3  .  Obj: *ast.Object {
//4  .  .  Kind: func
//5  .  .  Name: "main"
//6  .  .  Decl: *ast.FuncDecl {
//7  .  .  .  Name: *(obj @ 0)
//8  .  .  .  Type: *ast.FuncType {
//9  .  .  .  .  Func: test.go:3:1
//10  .  .  .  .  Params: *ast.FieldList {
//11  .  .  .  .  .  Opening: test.go:3:10
//12  .  .  .  .  .  Closing: test.go:3:11
//13  .  .  .  .  }
//14  .  .  .  }
//15  .  .  .  Body: *ast.BlockStmt {
//16  .  .  .  .  Lbrace: test.go:3:13
//17  .  .  .  .  List: []ast.Stmt (len = 1) {
//18  .  .  .  .  .  0: *ast.ExprStmt {
//19  .  .  .  .  .  .  X: *ast.CallExpr {
//20  .  .  .  .  .  .  .  Fun: *ast.Ident {
//21  .  .  .  .  .  .  .  .  NamePos: test.go:4:2
//22  .  .  .  .  .  .  .  .  Name: "println"
//23  .  .  .  .  .  .  .  }
//24  .  .  .  .  .  .  .  Lparen: test.go:4:9
//25  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
//26  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
//27  .  .  .  .  .  .  .  .  .  ValuePos: test.go:4:10
//28  .  .  .  .  .  .  .  .  .  Kind: STRING
//29  .  .  .  .  .  .  .  .  .  Value: "\"Hello, World!\""
//30  .  .  .  .  .  .  .  .  }
//31  .  .  .  .  .  .  .  }
//32  .  .  .  .  .  .  .  Ellipsis: -
//33  .  .  .  .  .  .  .  Rparen: test.go:4:25
//34  .  .  .  .  .  .  }
//35  .  .  .  .  .  }
//36  .  .  .  .  }
//37  .  .  .  .  Rbrace: test.go:5:1
//38  .  .  .  }
//39  .  .  }
//40  .  }
//41  }
//0  nil
//0  *ast.FuncType {
//1  .  Func: test.go:3:1
//2  .  Params: *ast.FieldList {
//3  .  .  Opening: test.go:3:10
//4  .  .  Closing: test.go:3:11
//5  .  }
//6  }
//0  *ast.FieldList {
//1  .  Opening: test.go:3:10
//2  .  Closing: test.go:3:11
//3  }
//0  nil
//0  nil
//0  *ast.BlockStmt {
//1  .  Lbrace: test.go:3:13
//2  .  List: []ast.Stmt (len = 1) {
//3  .  .  0: *ast.ExprStmt {
//4  .  .  .  X: *ast.CallExpr {
//5  .  .  .  .  Fun: *ast.Ident {
//6  .  .  .  .  .  NamePos: test.go:4:2
//7  .  .  .  .  .  Name: "println"
//8  .  .  .  .  }
//9  .  .  .  .  Lparen: test.go:4:9
//10  .  .  .  .  Args: []ast.Expr (len = 1) {
//11  .  .  .  .  .  0: *ast.BasicLit {
//12  .  .  .  .  .  .  ValuePos: test.go:4:10
//13  .  .  .  .  .  .  Kind: STRING
//14  .  .  .  .  .  .  Value: "\"Hello, World!\""
//15  .  .  .  .  .  }
//16  .  .  .  .  }
//17  .  .  .  .  Ellipsis: -
//18  .  .  .  .  Rparen: test.go:4:25
//19  .  .  .  }
//20  .  .  }
//21  .  }
//22  .  Rbrace: test.go:5:1
//23  }
//0  *ast.ExprStmt {
//1  .  X: *ast.CallExpr {
//2  .  .  Fun: *ast.Ident {
//3  .  .  .  NamePos: test.go:4:2
//4  .  .  .  Name: "println"
//5  .  .  }
//6  .  .  Lparen: test.go:4:9
//7  .  .  Args: []ast.Expr (len = 1) {
//8  .  .  .  0: *ast.BasicLit {
//9  .  .  .  .  ValuePos: test.go:4:10
//10  .  .  .  .  Kind: STRING
//11  .  .  .  .  Value: "\"Hello, World!\""
//12  .  .  .  }
//13  .  .  }
//14  .  .  Ellipsis: -
//15  .  .  Rparen: test.go:4:25
//16  .  }
//17  }
//0  *ast.CallExpr {
//1  .  Fun: *ast.Ident {
//2  .  .  NamePos: test.go:4:2
//3  .  .  Name: "println"
//4  .  }
//5  .  Lparen: test.go:4:9
//6  .  Args: []ast.Expr (len = 1) {
//7  .  .  0: *ast.BasicLit {
//8  .  .  .  ValuePos: test.go:4:10
//9  .  .  .  Kind: STRING
//10  .  .  .  Value: "\"Hello, World!\""
//11  .  .  }
//12  .  }
//13  .  Ellipsis: -
//14  .  Rparen: test.go:4:25
//15  }
//0  *ast.Ident {
//1  .  NamePos: test.go:4:2
//2  .  Name: "println"
//3  }
//0  nil
//0  *ast.BasicLit {
//1  .  ValuePos: test.go:4:10
//2  .  Kind: STRING
//3  .  Value: "\"Hello, World!\""
//4  }

//Package: 2:1代表Go解析出package这个词在第二行的第一个
//main是一个ast.Ident标识符，它的位置在第二行的第9个
//此处func main被解析成ast.FuncDecl（function declaration）,而函数的参数（Params）和函数体（Body）自然也在这个FuncDecl中。Params对应的是*ast.FieldList，顾名思义就是项列表；而由大括号“｛｝”组成的函数体对应的是ast.BlockStmt（block statement）
//而对于main函数的函数体中，我们可以看到调用了println函数，在ast中对应的是ExprStmt（Express Statement），调用函数的表达式对应的是CallExpr(Call Expression)，调用的参数自然不能错过，因为参数只有字符串，所以go把它归为ast.BasicLis (a literal of basic type)。
//最后，我们可以看出ast还解析出了函数的作用域，以及作用域对应的对象。
//Go将所有可以识别的token抽象成Node，通过interface方式组织在一起。
//在这里说到token我们需要说一下词法分析，token是词法分析的结果，即将字符序列转换为标记(token)的过程，这个操作由词法分析器完成。这里的标记是一个字符串，是构成源代码的最小单位。

func main() {

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		// Called recursively.
		ast.Print(fset, n)
		return true
	})
	//// Print the AST.
	//ast.Print(fset, f)
}
