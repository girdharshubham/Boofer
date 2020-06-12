test:
	$(shell mkdir -p reports)
	go test -covermode=atomic -coverpkg=./... -coverprofile reports/test.out ./...
view:
	go tool cover -html=reports/test.out	
package:
	go build -o bin/boofer main.go
run:
	bin/boofer
clean:
	$(shell rm -rf reports bin)
