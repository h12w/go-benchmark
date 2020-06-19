#!/bin/bash

cat receiver_test.go.tmpl | env Max=20 gomplate  > receiver_test.go
go fmt
GCFLAGS=
go test $GCFLAGS -bench Value
go test $GCFLAGS -bench Pointer

GCFLAGS=-gcflags='-l' 
go test $GCFLAGS -bench Value
go test $GCFLAGS -bench Pointer
