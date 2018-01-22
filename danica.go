package danica

import (
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
)

const (
	crazy_miguel_stuff = `require=(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);var f=new Error("Cannot find module '"+o+"'");throw f.code="MODULE_NOT_FOUND",f}var l=n[o]={exports:{}};t[o][0].call(l.exports,function(e){var n=t[o][1][e];return s(n?n:e)},l,l.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){`
)

// BundleFile bundles the file named in src and outputs it to file named dest.
func BundleFile(src, dest string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()
	d, err := openDest(dest)
	if err != nil {
		return err
	}
	defer d.Close()
	c, err := io.Copy(d, f)
	log.Printf("%d bytes written\n", c)
	return err
}

func BundleContent(content, dest string) error {
	program, err := parser.ParseFile(nil, "", content, 0)
	if err != nil {
		return err
	}
	fnames := findRequires(program)
	d, err := openDest(dest)
	if err != nil {
		return err
	}
	defer d.Close()
	for i := range fnames {
		f, err := os.Open(fnames[i])
		if err != nil {
			return err
		}
		io.Copy(d, f)
		f.Close()
		io.WriteString(d, "\n")
	}
	return err
}

func findRequires(p *ast.Program) (filenames []string) {
	ast.Walk(&enterOnly{func(n ast.Node) {
		ce, ok := n.(*ast.CallExpression)
		if !ok {
			return
		}
		id, ok := ce.Callee.(*ast.Identifier)
		if !ok {
			return
		}
		if id.Name != "require" {
			return
		}
		fname := ce.ArgumentList[0].(*ast.StringLiteral).Literal
		fname = strings.Trim(fname, `"`)
		filenames = append(filenames, fname)
		return
	}}, p)
	return
}

type enterOnly struct {
	f func(n ast.Node)
}

func (e *enterOnly) Enter(n ast.Node) (v ast.Visitor) {
	e.f(n)
	return e
}

func (e *enterOnly) Exit(n ast.Node) {}

func openDest(dest string) (*os.File, error) {
	basename := path.Dir(dest)
	err := os.MkdirAll(basename, os.ModePerm)
	if err != nil {
		return nil, err
	}
	d, err := os.Create(dest)
	if err != nil {
		return nil, err
	}
	io.WriteString(d, crazy_miguel_stuff)
	io.WriteString(d, "\n")
	return d, nil
}
