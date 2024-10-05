npm i -g degit
cd "$1" || exit

degit https://github.com/upikoth/starter-vue3 "$2"

cd "$2" || exit
sed -i "" "s/starter-go/$3/" package.json
