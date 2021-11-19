BINARY_NAME=psone

all:
	go build -o ${BINARY_NAME}
	go install
