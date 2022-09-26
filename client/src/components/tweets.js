import { useEffect, useState } from "react";
import axios from "axios";
import Tweet from "./tweet";
import { Typography, CircularProgress } from "@mui/material";

const Tweets = () => {
  const [tweets, setTweets] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isError, setIsError] = useState(false);

  useEffect(() => {
    const fetchTweets = async () => {
      try {
        const { data } = await axios.get(process.env.REACT_APP_ROLL_URL);
        setTweets(data);
        setIsLoading(false);
        setIsError(false);
      } catch {
        setIsError(true);
      }
    };

    fetchTweets();
  }, []);

  if (isError) {
    return <Typography>We can't process your request</Typography>;
  }

  if (isLoading) {
    return <CircularProgress color="inherit" />;
  }

  if (tweets.length === 0) {
    return <Typography>No Tweets to show...</Typography>;
  }

  return (
    <div>
      {tweets.map((tweet) => (
        <Tweet tweet={tweet} key={tweet.data.id} />
      ))}
    </div>
  );
};

export default Tweets;
