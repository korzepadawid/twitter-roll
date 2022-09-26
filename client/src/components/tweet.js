import {
  Avatar,
  Card,
  CardHeader,
  Typography,
  CardContent,
  Chip,
  CardActions,
} from "@mui/material";
import PersonIcon from "@mui/icons-material/Person";
import moment from "moment";

const Tweet = ({ tweet }) => (
  <Card sx={{ minWidth: 275, marginTop: 2 }}>
    <CardHeader
      avatar={
        <Avatar
          src={tweet.includes.users[0].profile_image_url}
          alt={tweet.includes.users[0].username}
        />
      }
      action={
        <Chip
          label="Link"
          component="a"
          color="primary"
          target="_blank"
          variant="outlined"
          href={`https://twitter.com/${tweet.includes.users[0].username}/status/${tweet.data.id}`}
          clickable
        />
      }
      title={`@${tweet.includes.users[0].username}`}
      subheader={moment(tweet.data.created_at).calendar()}
    />
    <CardContent>
      <Typography variant="h6">{tweet.data.text}</Typography>
    </CardContent>
    <CardActions>
      {tweet.includes.users.map(({ username }) => (
        <Chip
          key={username}
          label={`@${username}`}
          component="a"
          icon={<PersonIcon />}
          color="primary"
          target="_blank"
          variant="outlined"
          href={`https://twitter.com/${username}`}
          clickable
        />
      ))}
    </CardActions>
  </Card>
);

export default Tweet;
