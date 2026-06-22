const badgeBase =
  "inline-flex items-center rounded-md border border-transparent px-2 py-0.5 text-xs font-medium text-white";

/** Dev overlay only — explicit Tailwind palette for status chips. */
export const statusBadge = {
  success: `${badgeBase} bg-emerald-500`,
  error: `${badgeBase} bg-pink-600`,
  warning: `${badgeBase} bg-amber-500`,
  info: `${badgeBase} bg-sky-600`,
} as const;

export type StatusTone = keyof typeof statusBadge;
