# Dev Server
You'll need to start the Go backend, as well as the React app:
```bash
go run main.go &
cd app
npm run dev
```

# End-to-end Tests
> Warning: the e2e tests are flaky and cost money. It costs ~$5 to run all e2e tests.
>
> TODO: the `seed` parameter is supposed to make OpenAI responses deterministic-ish. It doesn't
seem to work for large queries like ours.

We have several data sets and prompt sets that we run end-to-end tests on, then evaluate manually.
This gives us a sense for how Vizzy is improving over time.

To run all the tests:
```bash
go run test/e2e/main.go
```

To run only specific cases/visualizations:
```bash
TEST_CASE=countries TEST_VISUALIZATION=0 TEST_SUBVISUALIZATION=0 go run test/e2e/main.go
```

To view the results:
```bash
go run test/server.go
```
and visit `localhost:3333/home`. You can page through each test visualization.

To update the scores, edit `test/e2e/scores.yaml` with your ratings, then run
```bash
go run test/scores/main.go
```
to see the average score.
