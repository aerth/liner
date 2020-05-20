PREFIX ?= /usr/local/
NAME := liner

${NAME}: *.go
	go build -o $@
clean:
	rm -f ${NAME}
install:
	install ${NAME} ${PREFIX}/bin/
