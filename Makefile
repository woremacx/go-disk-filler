PROG = go-disk-filler
SRC = main.go

all: $(PROG) $(PROG).exe


$(PROG): $(SRC)
	go build

$(PROG).exe: $(SRC)
	GOOS=windows GOARCH=386 go build
