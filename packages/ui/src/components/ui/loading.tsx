import type { ComponentType } from "react";
import LottieModule from "lottie-react";
import pixelDuckAnimation from "../../assets/pixel-duck.json";

type LottieProps = {
  animationData: unknown;
  loop?: boolean;
  className?: string;
};

const Lottie = (
  typeof LottieModule === "function"
    ? LottieModule
    : (LottieModule as { default?: unknown }).default
) as ComponentType<LottieProps>;

export function Loading() {
  return (
    <div className="ui-loading" role="status" aria-live="polite">
      <div className="ui-loading-panel">
        <Lottie
          animationData={pixelDuckAnimation}
          loop={true}
          className="ui-loading-animation"
        />
        <p className="ui-loading-text">Loading...</p>
      </div>
    </div>
  );
}
