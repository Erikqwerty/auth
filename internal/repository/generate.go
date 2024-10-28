package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate $HOME/go/bin/minimock -i AuthRepository -o ./mocks/ -s "_minimock.go"
