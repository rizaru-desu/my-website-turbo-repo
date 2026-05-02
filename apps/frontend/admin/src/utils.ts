export function getFieldError(errors: Array<unknown>) {
  const error = errors[0];

  if (typeof error === "string") {
    return error;
  }

  if (error instanceof Error) {
    return error.message;
  }

  return undefined;
}

// XP/Level calculation from portfolio metrics
export interface XpData {
  level: number;
  currentXp: number;
  xpToNextLevel: number;
  totalXp: number;
  xpBreakdown: {
    fromStars: number;
    fromViews: number;
    fromReads: number;
  };
}

export function calculateXp({
  totalStars,
  totalViews,
  totalReads,
}: {
  totalStars: number;
  totalViews: number;
  totalReads: number;
}): XpData {
  // XP formula: stars give most, views and reads give passive XP
  const fromStars = totalStars * 50;
  const fromViews = Math.floor(totalViews / 100);
  const fromReads = Math.floor(totalReads / 50);
  const totalXp = fromStars + fromViews + fromReads;

  // Level formula: sqrt curve for gradual progression
  const level = Math.max(1, Math.floor(Math.sqrt(totalXp / 100)) + 1);
  const xpToNextLevel = level * 10000;
  const currentXp = totalXp % xpToNextLevel;

  return {
    level,
    currentXp,
    xpToNextLevel,
    totalXp,
    xpBreakdown: {
      fromStars,
      fromViews,
      fromReads,
    },
  };
}
