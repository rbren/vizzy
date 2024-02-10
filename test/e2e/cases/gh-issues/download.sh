pages=(1 2 3 4 5 6 7 8 9 10)
rm -f raw.json && touch raw.json
for page in ${pages[@]}; do
    echo "https://api.github.com/repos/kubernetes/kubernetes/issues?per_page=100&page=$page"
    curl "https://api.github.com/repos/kubernetes/kubernetes/issues?per_page=100&page=$page" >> raw.json
done

cat ./raw.json | jq -s '. | flatten' > ./data

