package db

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate $HOME/go/bin/minimock -i TxManager -o ./mocks/ -s "_minimock.go"
