package cache

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate $HOME/go/bin/minimock -i RedisClient  -o ./mocks/ -s "_minimock.go"
