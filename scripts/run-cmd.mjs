import { spawn } from "node:child_process";
import { log } from "./log.mjs";
import { root } from "./load-config.mjs";

/**
 * Run a command and wait for exit. Throws on non-zero exit code.
 * @param {string} command
 * @param {string[]} args
 * @param {{ env?: NodeJS.ProcessEnv, label?: string }} [options]
 */
export function runCmd(command, args, options = {}) {
  const { env = process.env, label = `${command} ${args.join(" ")}` } = options;
  log(label);

  return new Promise((resolve, reject) => {
    const child = spawn(command, args, {
      cwd: root,
      stdio: "inherit",
      env,
      shell: process.platform === "win32" && !["go", "node", "bun"].includes(command),
      windowsHide: true,
    });

    child.on("error", reject);
    child.on("exit", (code, signal) => {
      if (signal) {
        reject(new Error(`${label} terminated by ${signal}`));
        return;
      }
      if (code !== 0) {
        reject(new Error(`${label} exited with code ${code}`));
        return;
      }
      resolve(undefined);
    });
  });
}

/**
 * Spawn a long-running child process.
 * @param {string} command
 * @param {string[]} args
 * @param {{ env?: NodeJS.ProcessEnv }} [options]
 */
export function spawnCmd(command, args, options = {}) {
  const { env = process.env } = options;
  return spawn(command, args, {
    cwd: root,
    stdio: "inherit",
    env,
    shell: process.platform === "win32" && !["go", "node", "bun"].includes(command),
    windowsHide: true,
  });
}

/** @param {import("node:child_process").ChildProcess} child */
export function killChild(child) {
  if (!child || child.exitCode !== null || child.signalCode) {
    return;
  }
  if (process.platform === "win32") {
    try {
      child.kill();
    } catch {
      /* ignore */
    }
    return;
  }
  child.kill("SIGTERM");
}
