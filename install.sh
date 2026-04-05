echo "state: building icefetch"
mkdir -p ./build
go build -o icefetch ./src/main.go 
mv ./icefetch ./build/icefetch

echo "state: installing icefetch for user"

mkdir -p ~/.config/icefetch

cp -r ./ascii/* ~/.config/icefetch/ascii
cp ./config/icefetch.toml ~/.config/icefetch/icefetch.toml
cp ./build/icefetch ~/.local/bin/icefetch

echo "state: cleaning user install"
rm -rf ./build/*


echo "installed icefetch for the user sucessfully :3"


