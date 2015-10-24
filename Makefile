init:	exppath comp

exppath:
	export GOPATH=./src/jwakeup

comp:
	go build jwakeup

clean:
	rm -r *~
