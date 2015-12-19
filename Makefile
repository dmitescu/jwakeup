comp:
	env GOPATH=$(PWD) go build github.com/dmitescu/jwakeup

clean:
	rm -r *~
