mkdir -p dist

cd grup
GOOS=windows go build -o grup . && mv grup ../dist/grup-win
GOOS=linux go build -o grup . && mv grup ../dist/grup-linux
GOOS=darwin go build -o grup . && mv grup ../dist/grup-darwin
cd ..

cd el
GOOS=windows go build -o el . && mv el ../dist/el-win
GOOS=linux go build -o el . && mv el ../dist/el-linux
GOOS=darwin go build -o el . && mv el ../dist/el-darwin
cd ..

cd kat
GOOS=windows go build -o kat . && mv kat ../dist/kat-win
GOOS=linux go build -o kat . && mv kat ../dist/kat-linux
GOOS=darwin go build -o kat . && mv kat ../dist/kat-darwin
