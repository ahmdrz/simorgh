all: 
	@bash build.sh
install: all
	@cp simorgh-server /usr/local/bin
	@cp simorgh-client /usr/local/bin