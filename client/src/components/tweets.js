import { useEffect, useState } from "react";
import axios from "axios";
import Tweet from "./tweet";

const Tweets = () => {
  const [tweets, setTweets] = useState([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchTweets = async () => {
      const { data } = await axios.get(
        "https://mocki.io/v1/a8c39b14-b6bb-4843-8fb7-980e82b3b845"
      );
      console.log(data);
      setTweets(data);
      setIsLoading(false);
    };

    fetchTweets();
  }, []);

  if (isLoading) {
    return <p>Loading...</p>;
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
