gen:
	go run ../../cmd/microgen/main.go   -file api/service.go -out . -main

clean:
	find service/ -name \*.microgen.* -type f -delete
	find transport/ -name \*.microgen.* -type f -delete
	#rm cmd/person_service/main.go