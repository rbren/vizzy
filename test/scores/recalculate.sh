#! /bin/bash
set -eo pipefail

out=$(go run ./test/scores/main.go)
avg=$(echo $out | jq -r '.average_score')
acc=$(echo $out | jq -r '.accuracy')
good=$(echo $out | jq -r '.percentage_good')
badge=$good
sed -i "s/^# Average.*/# Average: $avg/" ./test/e2e/scores.yaml
sed -i "s/^# Accuracy.*/# Accuracy: $acc%/" ./test/e2e/scores.yaml
sed -i "s/^# Percent Good.*/# Percent Good: $good%/" ./test/e2e/scores.yaml
sed -i "s/[0-9]\+\(\.[0-9]\+\)\?% success rate/$badge% success rate/" ./app/src/components/Home.tsx
sed -i "s/Accuracy [0-9]\+\(\.[0-9]\+\)\?%/Accuracy $badge%/" ./README.md
sed -i "s/Accuracy-[0-9]\+\(\.[0-9]\+\)\?%/Accuracy-$badge%/" ./README.md
echo "Scores recalculated"
echo "Average: $avg"
echo "Accuracy: $acc%"
echo "Percent Good: $good%"
