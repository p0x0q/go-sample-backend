test:
	curl -i -H 'Content-Type: application/json' \
	 -d '{"Code":"JP","Name":"Japan"}' http://127.0.0.1:8081/countries
setup-air:
	go get -u github.com/cosmtrek/air
