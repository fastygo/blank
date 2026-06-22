export type DevPanelID = "health" | "assets" | "request";

export interface DevContext {
  requestId: string;
  path: string;
  locale: string;
  fetchStatus: () => Promise<StatusPayload>;
  disableOverlay: () => Promise<void>;
}

export interface DevPanel {
  id: DevPanelID;
  title: string;
  mount(root: HTMLElement, context: DevContext): () => void;
}

export interface AssetStatus {
  id: string;
  path: string;
  exists: boolean;
  size: number;
  mtime: string;
  ageSec: number;
  stale?: boolean;
  hint?: string;
}

export interface StatusPayload {
  bind: string;
  assets: AssetStatus[];
  hints?: string[];
  overlay: boolean;
}

export interface ProbeResult {
  path: string;
  ok: boolean;
  status: number;
  latencyMs: number;
  error?: string;
}
