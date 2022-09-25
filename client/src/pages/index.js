import { BrowserRouter, Routes, Route } from "react-router-dom";
import Roll from "./roll";
import NotFound from "./error";

const App = () => (
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<Roll />} />
      <Route path="*" element={<NotFound />} />
    </Routes>
  </BrowserRouter>
);

export default App;
