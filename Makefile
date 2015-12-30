comp:	
	CGO_ENABLED=1 GOPATH=$(PWD) go build -gccgoflags "-pthread" github.com/dmitescu/jwakeup

dependencies:
	GOPATH=$(PWD) go get github.com/mattn/go-sqlite3

clean:
	rm -r *~
