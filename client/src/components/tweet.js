const Tweet = ({ tweet }) => {
  const usersWithoutAuthor = tweet.includes.users.filter(
    (user) => user.id !== tweet.data.author_id
  );
  return (
    <div>
      <img src={tweet.includes.users[0].profile_image_url} alt="avatar" />
      <h3>
        {tweet.includes.users[0].name} {`@${tweet.includes.users[0].username}`}
      </h3>
      <h2>{tweet.data.text}</h2>
      <p>{tweet.data.created_at}</p>
      <a
        target="_blank"
        href={`https://twitter.com/${tweet.includes.users[0].username}/status/${tweet.data.id}`}
        rel="noreferrer"
      >
        Link
      </a>
      <br />
      {usersWithoutAuthor.map((user) => (
        <a
          key={user.id}
          target="_blank"
          href={`https://twitter.com/${user.username}`}
          rel="noreferrer"
        >
          {user.username}
        </a>
      ))}
    </div>
  );
};

export default Tweet;
