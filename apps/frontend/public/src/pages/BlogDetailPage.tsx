import { useParams, Link } from "react-router-dom";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { PublicNav, PublicFooter, SectionHeader } from "../components";

// Mock data (matching LandingPage)
const blogPosts = {
  "art-of-pixel-perfect-design-systems": {
    title: "THE ART OF PIXEL-PERFECT DESIGN SYSTEMS",
    category: "DESIGN",
    publishDate: "MAY 01, 2026",
    readingTime: "5 MIN READ",
    content: `
# The Return of the Pixel

In a world of high-resolution displays and ultra-smooth gradients, there's a surprising trend emerging: the intentional return to pixelated aesthetics. But this isn't just nostalgia; it's a calculated design choice.

## Why Pixels Matter

Pixels represent the fundamental building blocks of the digital world. By embracing them, we acknowledge the medium we're working in. 

### Key Benefits:
- **Clarity**: Pixel-based icons are often more legible at small sizes than their vector counterparts.
- **Personality**: It stands out in a sea of "clean and corporate" SaaS designs.
- **Constraint**: Working with a grid forces you to make decisive design choices.

\`\`\`css
.pixel-border {
  border: 4px solid #282a36;
  box-shadow: 
    inset -4px -4px 0px 0px rgba(0,0,0,0.1),
    inset 4px 4px 0px 0px rgba(255,255,255,0.1);
}
\`\`\`

> "Constraint is the mother of creativity."

## Building Your Own System

When building a pixel-flavored system, consistency is key. You can't mix pixel sizes. If your "pixel" is 4x4 real pixels, everything in your UI must follow that 4px grid.

Stay tuned for part two where we dive into the code!
`,
  },
  "why-vanilla-css-in-2026": {
    title: "WHY I STILL USE VANILLA CSS IN 2026",
    category: "CODE",
    publishDate: "APR 20, 2026",
    readingTime: "8 MIN READ",
    content: "# Vanilla CSS is still king\n\nContent coming soon...",
  },
  "retro-futurism-ux-perspective": {
    title: "RETRO-FUTURISM: A UX PERSPECTIVE",
    category: "UX RESEARCH",
    publishDate: "APR 10, 2026",
    readingTime: "6 MIN READ",
    content: "# Retro-Futurism in UX\n\nContent coming soon...",
  },
};

const mockComments = [
  {
    id: "1",
    displayName: "ALEX RIDER",
    date: "MAY 02, 2026",
    body: "This is exactly what I needed! I've been struggling to get the pixel-perfect look on my new dashboard. The CSS snippet is super helpful.",
  },
  {
    id: "2",
    displayName: "SARAH CHEN",
    date: "MAY 03, 2026",
    body: "Love the point about constraints. It's so true that limiting your options can often lead to better outcomes.",
  },
];

export default function BlogDetailPage() {
  const { slug } = useParams();
  const post = blogPosts[slug as keyof typeof blogPosts];

  if (!post) {
    return (
      <div className="public-shell">
        <PublicNav />
        <main className="section mt-20 min-h-[60vh] flex flex-col items-center justify-center text-center">
          <h1 className="text-pink mb-4" style={{ fontFamily: "var(--font-pixel)", fontSize: "2rem" }}>404: ENTRY_NOT_FOUND</h1>
          <p className="text-muted mb-8" style={{ fontFamily: "var(--font-terminal)" }}>The requested transmission could not be retrieved from the archives.</p>
          <Link to="/blog" className="pix-btn pix-btn-ghost hover-wiggle">« BACK_TO_INDEX</Link>
        </main>
        <PublicFooter ownerName="KAI MORIKAWA" />
      </div>
    );
  }

  return (
    <div className="public-shell">
      <PublicNav />
      
      <main className="section mt-20 max-w-[800px] mx-auto">
        <div className="mb-10">
          <Link to="/blog" className="text-purple hover:underline mb-4 inline-block" style={{ fontFamily: "var(--font-pixel)", fontSize: "0.5rem" }}>
            « BACK_TO_LOG
          </Link>
          <div className="flex gap-4 items-center mb-4">
             <span className="pixel-badge text-line">{post.category}</span>
             <span className="text-muted" style={{ fontFamily: "var(--font-terminal)", fontSize: "0.9rem" }}>{post.publishDate}</span>
             <span className="text-muted" style={{ fontFamily: "var(--font-terminal)", fontSize: "0.9rem" }}>• {post.readingTime}</span>
          </div>
          <h1 className="text-text leading-tight mb-8" style={{ fontSize: "2.5rem" }}>{post.title}</h1>
        </div>

        <article className="prose-retro mb-20">
          <ReactMarkdown remarkPlugins={[remarkGfm]}>
            {post.content}
          </ReactMarkdown>
        </article>

        <section className="comments-section border-t-4 border-panel pt-12">
          <SectionHeader kicker="~ FEEDBACK" title="REPLIES.TXT" />
          
          <div className="space-y-8 mt-10">
            {mockComments.map((comment) => (
              <div key={comment.id} className="card accent-purple">
                <div className="flex justify-between items-center mb-4">
                  <div className="text-purple" style={{ fontFamily: "var(--font-pixel)", fontSize: "0.54rem" }}>{comment.displayName}</div>
                  <div className="text-muted" style={{ fontFamily: "var(--font-terminal)", fontSize: "0.8rem" }}>{comment.date}</div>
                </div>
                <p className="m-0 text-text leading-relaxed" style={{ fontFamily: "var(--font-terminal)", fontSize: "1rem" }}>
                  {comment.body}
                </p>
              </div>
            ))}
          </div>

          <div className="mt-12">
            <h3 className="text-cyan mb-6" style={{ fontFamily: "var(--font-pixel)", fontSize: "0.8rem" }}>&gt; LEAVE_A_REPLY</h3>
            <div className="card accent-cyan p-0">
              <textarea 
                className="w-full bg-transparent border-none text-text p-4 focus:ring-0 min-h-[150px]"
                placeholder="TYPE YOUR MESSAGE HERE..."
                style={{ fontFamily: "var(--font-terminal)", fontSize: "1rem" }}
              />
              <div className="p-4 border-t-2 border-panel flex justify-end">
                <button className="pix-btn pix-btn-cyan hover-wiggle">SUBMIT_REPLY</button>
              </div>
            </div>
          </div>
        </section>

        <div className="mt-20 flex justify-center pb-20">
          <Link to="/blog" className="pix-btn pix-btn-ghost hover-wiggle">
            « BACK_TO_LOG
          </Link>
        </div>
      </main>

      <PublicFooter ownerName="KAI MORIKAWA" />
    </div>
  );
}
