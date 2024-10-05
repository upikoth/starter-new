cd "$1" || exit
git init -b main
git add .
git commit -m "🐣"
git remote add origin "$2"
git push -u origin main
