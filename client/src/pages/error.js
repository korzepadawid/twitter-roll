import { Typography } from "@mui/material";
import { Link } from "react-router-dom";

const NotFound = () => (
  <>
    <Typography variant="h3" component="h1">
      Have you got lost?
    </Typography>
    <Typography>
      <Link to="/">Go back to main page</Link>
    </Typography>
  </>
);

export default NotFound;
