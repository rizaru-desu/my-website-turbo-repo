import { Routes, Route } from "react-router-dom";
import { LandingPage, TestimonialFormPage, BlogListPage, BlogDetailPage } from "./pages";

function App() {
  return (
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/blog" element={<BlogListPage />} />
      <Route path="/blog/:slug" element={<BlogDetailPage />} />
      <Route path="/testimonials/new" element={<TestimonialFormPage />} />
      <Route path="*" element={<LandingPage />} />
    </Routes>
  );
}

export default App;
