#test:
	#go run ../main.go -type TestWalk -out ./

gen:
	rm -rf **_gen.go
	go generate
	go vet ./...
