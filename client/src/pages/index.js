import { BrowserRouter, Routes, Route } from "react-router-dom";
import Container from "@mui/material/Container";
import Roll from "./roll";
import NotFound from "./error";

const App = () => (
  <Container maxWidth="sm">
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Roll />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  </Container>
);

export default App;
