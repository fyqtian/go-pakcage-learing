package main

import (
	"context"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"log"
)

var ctx = context.Background()

func main() {
	// caution path
	rs, err := load(ctx, "./example", nil, "", nil)
	if err != nil {
		log.Fatalln(err)
	}

	for _, row := range rs {
		j, _ := row.MarshalJSON()
		log.Println(string(j))
		for _, r := range row.Syntax {
			ast.Inspect(r, func(node ast.Node) bool {
				switch x := node.(type) {

				case *ast.GenDecl:
					if x.Tok == token.TYPE {
						for _, s := range x.Specs {
							vSpec := s.(*ast.TypeSpec)
							if _, ok := vSpec.Type.(*ast.StructType); ok {
								log.Println("TypeSpec Name", vSpec.Name.Name)
							}
						}
					}
				case *ast.FuncDecl:
					log.Println("FuncDecl Name", x.Name)
				case *ast.StructType:
					for _, f := range x.Fields.List {
						log.Println("Struct Field Name", f.Names)
					}
				}
				return true
			})
		}
	}

}

func load(ctx context.Context, path string, env []string, tags string, patterns []string) ([]*packages.Package, []error) {
	cfg := &packages.Config{
		Context:    ctx,
		Mode:       packages.LoadAllSyntax,
		Dir:        path,
		Env:        env,
		BuildFlags: []string{"-tags=my"},
	}
	if len(tags) > 0 {
		cfg.BuildFlags[0] += " " + tags
	}
	escaped := make([]string, len(patterns))
	for i := range patterns {
		escaped[i] = "pattern=" + patterns[i]
	}
	pkgs, err := packages.Load(cfg, escaped...)
	if err != nil {
		return nil, []error{err}
	}
	var errs []error
	for _, p := range pkgs {
		for _, e := range p.Errors {
			errs = append(errs, e)
		}
	}
	if len(errs) > 0 {
		return nil, errs
	}
	return pkgs, nil
}
