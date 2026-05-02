import { calculateXp, type XpData } from "../../utils";
import type { Project, Post } from "../../types";

interface XpBarProps {
  projects: Project[];
  posts: Post[];
  name?: string;
}

export default function XpBar({ projects, posts, name = "KAI" }: XpBarProps) {
  // Calculate totals from data
  const totalStars = projects.reduce((sum, p) => sum + p.stars, 0);
  const totalViews = projects.reduce((sum, p) => sum + p.views, 0);
  const totalReads = posts.reduce((sum, p) => sum + p.reads, 0);

  const xp: XpData = calculateXp({ totalStars, totalViews, totalReads });
  const progressPercent = (xp.currentXp / xp.xpToNextLevel) * 100;

  // Calculate segments for the pixel bar (16 segments like the image)
  const segmentCount = 16;
  const segmentProgress = Math.floor((progressPercent / 100) * segmentCount);

  return (
    <div className="xp-bar-container">
      <div className="xp-header">
        <span className="xp-name">{name}</span>
        <span className="xp-level">LVL {xp.level}</span>
      </div>
      <div className="xp-label">XP TO LVL {xp.level + 1}</div>
      <div className="xp-bar-frame">
        <div className="xp-bar-fill">
          {Array.from({ length: segmentCount }).map((_, i) => (
            <div
              key={i}
              className={`xp-segment ${i < segmentProgress ? "filled" : "empty"}`}
            />
          ))}
        </div>
      </div>
      <div className="xp-values">
        <span className="xp-current">{xp.currentXp.toLocaleString()}</span>
        <span className="xp-separator"> / </span>
        <span className="xp-max">{xp.xpToNextLevel.toLocaleString()}</span>
      </div>
    </div>
  );
}
