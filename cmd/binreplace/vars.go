package main

var source string // source binary file
var target string // target binary file

var fromhex string // binary sequence to replace
var tohex string   // binary sequence to replace with

var force bool      // overwrite target file if it exists
var mindistance int // minimum distance between two differences
