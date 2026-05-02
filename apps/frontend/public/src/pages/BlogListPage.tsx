import { Link } from "react-router-dom";
import { PublicNav, PublicFooter, SectionHeader, BlogCard } from "../components";

const COLORS = ["purple", "cyan", "pink", "green", "orange", "yellow"];

// Mock data (matching LandingPage)
const blogPosts = [
  {
    title: "THE ART OF PIXEL-PERFECT DESIGN SYSTEMS",
    slug: "art-of-pixel-perfect-design-systems",
    excerpt: "Exploring why the 8-bit aesthetic is making a comeback in modern SaaS interfaces and how to build one that scales.",
    category: "DESIGN",
    publishDate: "MAY 01, 2026",
    readingTime: "5 MIN READ",
    featured: true,
  },
  {
    title: "WHY I STILL USE VANILLA CSS IN 2026",
    slug: "why-vanilla-css-in-2026",
    excerpt: "Tailwind is great, but there's a certain magic in writing raw CSS variables and utility classes that you fully own.",
    category: "CODE",
    publishDate: "APR 20, 2026",
    readingTime: "8 MIN READ",
    featured: true,
  },
  {
    title: "RETRO-FUTURISM: A UX PERSPECTIVE",
    slug: "retro-futurism-ux-perspective",
    excerpt: "How nostalgic elements can improve user engagement by creating an emotional connection to digital tools.",
    category: "UX RESEARCH",
    publishDate: "APR 10, 2026",
    readingTime: "6 MIN READ",
    featured: true,
  },
  {
    title: "BUILDING LIGHTWEIGHT APPS IN A BLOATED WORLD",
    slug: "building-lightweight-apps",
    excerpt: "Strategies for keeping your bundle size small and your performance high when everything else is getting heavier.",
    category: "PERFORMANCE",
    publishDate: "MAR 25, 2026",
    readingTime: "10 MIN READ",
    featured: false,
  },
];

export default function BlogListPage() {
  return (
    <div className="public-shell">
      <PublicNav />
      
      <main className="section mt-20 min-h-[60vh]">
        <SectionHeader 
          kicker="~ INDEX" 
          title="TRANSMISSIONS.LOG" 
          subtitle="A collection of thoughts on design, code, and the aesthetics of the digital future."
        />
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 mt-10">
          {blogPosts.map((post, i) => (
            <BlogCard
              key={post.slug}
              title={post.title}
              excerpt={post.excerpt}
              category={post.category}
              publishDate={post.publishDate}
              readingTime={post.readingTime}
              slug={post.slug}
              color={COLORS[i % COLORS.length]}
              index={i}
            />
          ))}
        </div>

        <div className="mt-20 flex justify-center">
          <Link to="/" className="pix-btn pix-btn-ghost hover-wiggle">
            « RETURN_HOME
          </Link>
        </div>
      </main>

      <PublicFooter ownerName="KAI MORIKAWA" />
    </div>
  );
}
