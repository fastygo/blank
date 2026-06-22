export type DevPanelID = "health" | "assets" | "request";

export interface DevOverlayPanelI18n {
  health: {
    intro: string;
    refresh_button: string;
    ok_badge: string;
    http_badge: string;
    down_label: string;
    latency_unit: string;
    separator: string;
  };
  assets: {
    intro: string;
    present: string;
    missing: string;
    stale: string;
    age_prefix: string;
    separator: string;
    load_failed: string;
    missing_css_hint: string;
    stale_css_hint: string;
  };
  request: {
    intro: string;
    request_id_label: string;
    path_label: string;
    locale_label: string;
    copy_request_id: string;
    copied: string;
    copy_failed: string;
    empty_value: string;
    locale_missing: string;
  };
  errors: {
    fetch_failed: string;
    status_json_failed: string;
    disable_failed: string;
  };
}

export interface DevContext {
  requestId: string;
  path: string;
  locale: string;
  i18n: DevOverlayPanelI18n;
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
