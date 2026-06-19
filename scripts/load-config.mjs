import path from "node:path";
import { pathToFileURL } from "node:url";
import { fileURLToPath } from "node:url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

/** Repository root (parent of scripts/). */
export const root = path.resolve(__dirname, "..");

let cachedConfig;

/** Load and cache fastygo.config.mjs from the repository root. */
export async function loadConfig() {
  if (cachedConfig) {
    return cachedConfig;
  }
  const configPath = path.join(root, "fastygo.config.mjs");
  const mod = await import(pathToFileURL(configPath).href);
  cachedConfig = mod.default;
  return cachedConfig;
}

/** Resolve a config-relative path against the repository root. */
export function resolveFromRoot(relativePath) {
  return path.resolve(root, relativePath);
}

/** Merge server env from config with process.env (config wins for listed keys). */
export function serverEnv(config, baseEnv = process.env) {
  return { ...baseEnv, ...config.server.env };
}

/** Human-readable dev server URL from config. */
export function serverUrl(config) {
  const bind = config.server.env.APP_BIND ?? `${config.server.host}:${config.server.port}`;
  const host = bind.startsWith(":") ? `127.0.0.1${bind}` : bind.split(":")[0];
  const port = bind.includes(":") ? bind.split(":").pop() : String(config.server.port);
  return `http://${host}:${port}/`;
}
