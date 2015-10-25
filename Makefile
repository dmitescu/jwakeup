init:	exppath comp

exppath:
	export GOPATH=$(PWD)

comp:
	go build jwakeup

clean:
	rm -r *~
