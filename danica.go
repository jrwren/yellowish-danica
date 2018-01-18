package danica

import (
	"io"
	"log"
	"os"
	"path"
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
	basename := path.Dir(dest)
	log.Println("creating directory " + basename)
	err = os.MkdirAll(basename, os.ModePerm)
	if err != nil {
		return err
	}
	d, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()
	io.WriteString(d, crazy_miguel_stuff)
	io.WriteString(d, "\n")
	c, err := io.Copy(d, f)
	log.Printf("%d bytes written\n", c)
	return err
}

func BundleContent(content, dest string) error {
	return nil
}
