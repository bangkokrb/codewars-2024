SOL=./solutions/
BIN=$(SOL)/bin/

all: \
	$(BIN)/bknumerals \
	$(BIN)/dragon \
	$(BIN)/primes \
	$(BIN)/maxsum \
	$(BIN)/taxi \
	$(BIN)/treesearch \
	$(BIN)/paperfold \
	$(BIN)/wordsearch \
	$(BIN)/coins \
	$(BIN)/flag

$(BIN)/bknumerals: $(SOL)/bknumerals/bknumerals.go
	go build -o $(BIN)/bknumerals $(SOL)/bknumerals/bknumerals.go

$(BIN)/dragon: $(SOL)/dragon/dragon.go
	go build -o $(BIN)/dragon $(SOL)/dragon/dragon.go

$(BIN)/primes: $(SOL)/primes/primes.go
	go build -o $(BIN)/primes $(SOL)/primes/primes.go

$(BIN)/maxsum: $(SOL)/maxsum/maxsum.go
	go build -o $(BIN)/maxsum $(SOL)/maxsum/maxsum.go

$(BIN)/taxi: $(SOL)/taxi/taxi.go
	go build -o $(BIN)/taxi $(SOL)/taxi/taxi.go

$(BIN)/treesearch: $(SOL)/treesearch/treesearch.go
	go build -o $(BIN)/treesearch $(SOL)/treesearch/treesearch.go

$(BIN)/paperfold: $(SOL)/paperfold/paperfold.go
	go build -o $(BIN)/paperfold $(SOL)/paperfold/paperfold.go

$(BIN)/wordsearch: $(SOL)/wordsearch/wordsearch.go
	go build -o $(BIN)/wordsearch $(SOL)/wordsearch/wordsearch.go

$(BIN)/coins: $(SOL)/coins/coins.go
	go build -o $(BIN)/coins $(SOL)/coins/coins.go

$(BIN)/flag: $(SOL)/flag/flag.go
	go build -o $(BIN)/flag $(SOL)/flag/flag.go

clean:
	rm $(BIN)/*
