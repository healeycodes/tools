mkdir -p dist

cd grup
GOOS=windows go build -o grup . && mv grup ../dist/grup-win
GOOS=linux go build -o grup . && mv grup ../dist/grup-linux
GOOS=darwin go build -o grup . && mv grup ../dist/grup-darwin
cd ..

cd el
GOOS=windows go build -o el . && mv grup ../dist/el-win
GOOS=linux go build -o el . && mv grup ../dist/el-linux
GOOS=darwin go build -o el . && mv grup ../dist/el-darwin
cd ..

cd kat
GOOS=windows go build -o kat . && mv grup ../dist/kat-win
GOOS=linux go build -o kat . && mv grup ../dist/kat-linux
GOOS=darwin go build -o kat . && mv grup ../dist/kat-darwin
