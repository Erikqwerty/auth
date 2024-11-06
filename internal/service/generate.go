package service

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate $HOME/go/bin/minimock -i AuthService -o ./mocks/ -s "_minimock.go"
