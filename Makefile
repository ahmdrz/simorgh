all: 
	@cd server && go build -i -o ../simorgh-server
	@cd client && go build -i -o ../simorgh-client
install: all
	@cp simorgh-server /usr/local/bin
	@cp simorgh-client /usr/local/bin