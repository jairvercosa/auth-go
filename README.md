# Dependencies on your local environment

Installing `dep`
```
brew install dep
```

Installing `migrations`
```
brew install golang-migrate
```

Set the environment virable for migrations
```
export $DBURL="postgres://noodling_dev:123@localhost/noodling?sslmode=disable"
```