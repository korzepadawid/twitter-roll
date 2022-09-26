# twitter-roll

Twitter bot for filtering tweets from your favourite accounts and hashtags.

![image](https://i.ibb.co/Px7z4Zy/icon.png)

## Table of content
- [Tech](#tech)
- [Overall](#Overall)
- [Launch](#launch)

## Tech
- [Go](https://go.dev/)
- [Testify](https://github.com/stretchr/testify)
- [Twitter Stream API](https://developer.twitter.com/en/docs/tutorials/stream-tweets-in-real-time)
- [React.js](https://reactjs.org/)
- [MUI](https://mui.com/)

## Overall

![image](https://i.ibb.co/6Xmtyt2/appss.png)
- Consumed Twitter HTTP Streaming API, to receive tweets with low latency.
- Used Go concurrency mechanisms, such as goroutines, mutexes, and channels.
- Implemented a frontend with React.js to expose stored tweets.
## Launch
```
$ git clone git@github.com:korzepadawid/twitter-roll.git
```

```
$ cd twitter-roll
```

You can get your Twitter `Bearer Token` [here](https://developer.twitter.com/en/portal/dashboard).

```
$ echo "TWITTER_BEARER_TOKEN=<YOUR_BEARER_TOKEN>" >> app.env 
```

[Twitter Query builder](https://developer.twitter.com/apitools/query?query=) might be helpful here. You can paste whatever rule you want. 

```
$ echo "TWITTER_RULE=<STREAM_RULE>" >> app.env 
```

Specifies how many tweets can be exposed to a user.

```
$ echo "ROLL_CAPACITY=<ROLL_CAPACITY>" >> app.env 
```

API Endpoint.

```
$ echo "REACT_APP_ROLL_URL=http://localhost:8080/roll" >> ./client/.env 
```


