set -ex
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

go build -o server.exe server.go
./server.exe