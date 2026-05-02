import { Routes, Route } from "react-router-dom";
import { LandingPage, TestimonialFormPage } from "./pages";

function App() {
  return (
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/testimonials/new" element={<TestimonialFormPage />} />
      <Route path="*" element={<LandingPage />} />
    </Routes>
  );
}

export default App;
