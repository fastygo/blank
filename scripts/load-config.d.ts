export interface FastyGoConfig {
  server: {
    host: string;
    port: number;
    staticDir: string;
    env: Record<string, string>;
  };
  templ: {
    generate: string;
    watch: string[];
  };
  css: {
    input: string;
    output: string;
    sources: string[];
  };
  js: {
    bundles: Array<{ entry: string; output: string }>;
    manifest: string;
  };
  ui8px: {
    lint: string[];
    aria: {
      paths: string[];
      manifest: string;
    };
  };
}
