(()=>{function N(){let e=globalThis.__FASTYGO_DEV_OVERLAY__;if(e?.i18n)return e.i18n.trim();return document.querySelector("script[data-fastygo-dev-overlay-bundle]")?.dataset.i18n?.trim()??""}function y(){let e=N();if(!e)return null;try{return JSON.parse(e)}catch{return null}}function m(e,t){if(e.includes("%d"))return e.replace("%d",String(t));if(e.includes("%s"))return e.replace("%s",String(t));return e}var H="/__fastygo/dev/status.json",O="/__fastygo/dev/disable";async function w(e){let t=await fetch(H,{headers:{Accept:"application/json"},cache:"no-store"});if(!t.ok)throw Error(m(e.errors.status_json_failed,t.status));return t.json()}async function E(e){let t=await fetch(O,{method:"POST",headers:{Accept:"application/json"},cache:"no-store"});if(!t.ok)throw Error(m(e.errors.disable_failed,t.status));window.location.reload()}function _(){let e=globalThis.__FASTYGO_DEV_OVERLAY__;if(e?.requestId)return e.requestId.trim();return document.querySelector("script[data-fastygo-dev-overlay-bundle]")?.dataset.requestId?.trim()??""}function k(e){return document.documentElement.lang?.trim()||e}function P(){return`${window.location.pathname}${window.location.search}`}function $(e){if(e<1024)return`${e} B`;if(e<1048576)return`${(e/1024).toFixed(1)} KB`;return`${(e/1048576).toFixed(1)} MB`}function L(e){if(e<60)return`${e}s`;if(e<3600)return`${Math.floor(e/60)}m`;return`${Math.floor(e/3600)}h`}async function C(e,t){let r=performance.now();try{let n=await fetch(e,{cache:"no-store"});return{path:e,ok:n.ok,status:n.status,latencyMs:Math.round(performance.now()-r)}}catch(n){return{path:e,ok:!1,status:0,latencyMs:Math.round(performance.now()-r),error:n instanceof Error?n.message:t}}}var u={success:"inline-flex items-center rounded-md border border-transparent px-2 py-0.5 text-xs font-medium text-white bg-emerald-500",error:"inline-flex items-center rounded-md border border-transparent px-2 py-0.5 text-xs font-medium text-white bg-pink-600",warning:"inline-flex items-center rounded-md border border-transparent px-2 py-0.5 text-xs font-medium text-white bg-amber-500",info:"inline-flex items-center rounded-md border border-transparent px-2 py-0.5 text-xs font-medium text-white bg-sky-600"};function B(e){if(!e.exists)return u.error;if(e.stale)return u.warning;return u.success}function R(e,t){if(!e.exists)return t.missing;if(e.stale)return t.stale;return t.present}function j(e,t,r){let n=r.i18n.assets,a=document.createElement("div");a.className="rounded-md border border-border p-3 flex flex-col gap-1";let s=document.createElement("div");s.className="flex items-center justify-between gap-2";let i=document.createElement("code");i.className="text-sm",i.textContent=t.id;let l=document.createElement("span");l.className=B(t),l.textContent=R(t,n),s.append(i,l);let o=document.createElement("p");if(o.className="text-sm text-muted-foreground",!t.exists)o.textContent=t.path;else o.textContent=`${t.path}${n.separator}${$(t.size)}${n.separator}${n.age_prefix}${L(t.ageSec)}`;if(a.append(s,o),t.hint){let c=document.createElement("p");c.className="text-sm text-foreground",c.textContent=t.hint,a.append(c)}e.append(a)}var D={id:"assets",title:"Assets",mount(e,t){let r=t.i18n.assets;e.replaceChildren();let n=document.createElement("p");n.className="text-sm text-muted-foreground",n.textContent=r.intro,e.append(n);let a=document.createElement("div");a.className="flex flex-col gap-2",e.append(a);let s=!1;async function i(){a.replaceChildren();try{let l=await t.fetchStatus();if(s)return;for(let o of l.assets)j(a,o,t);if(l.hints?.length){let o=document.createElement("div");o.className="rounded-md border border-border bg-muted p-3 text-sm",o.textContent=l.hints.join(" "),a.append(o)}}catch(l){let o=document.createElement("p");o.className="text-sm text-destructive",o.textContent=l instanceof Error?l.message:r.load_failed,a.append(o)}}return i(),()=>{s=!0}}};function V(e,t){if(e)return u.success;if(t===0)return u.error;return u.info}function F(e,t){if(e.ok)return m(t.ok_badge,e.status);if(e.status)return m(t.http_badge,e.status);return t.down_label}function Y(e,t,r){let n=r.i18n.health,a=document.createElement("div");a.className="rounded-md border border-border p-3 flex flex-col gap-1";let s=document.createElement("div");s.className="flex items-center justify-between gap-2";let i=document.createElement("code");i.className="text-sm",i.textContent=t.path;let l=document.createElement("span");l.className=V(t.ok,t.status),l.textContent=F(t,n),s.append(i,l);let o=document.createElement("p");if(o.className="text-sm text-muted-foreground",o.textContent=`${t.latencyMs} ${n.latency_unit}`,t.error)o.textContent+=`${n.separator}${t.error}`;a.append(s,o),e.append(a)}var M={id:"health",title:"Health",mount(e,t){let r=t.i18n.health;e.replaceChildren();let n=document.createElement("p");n.className="text-sm text-muted-foreground",n.textContent=r.intro,e.append(n);let a=document.createElement("button");a.type="button",a.className="inline-flex h-8 items-center justify-center rounded-md border border-border bg-background px-3 text-sm hover:bg-accent",a.textContent=r.refresh_button,e.append(a);let s=document.createElement("div");s.className="flex flex-col gap-2",e.append(s);let i=!1;async function l(){s.replaceChildren();let c=await Promise.all(["/healthz","/readyz"].map((p)=>C(p,t.i18n.errors.fetch_failed)));if(i)return;for(let p of c)Y(s,p,t)}return a.addEventListener("click",()=>{l()}),l(),()=>{i=!0}}};function h(e,t,r,n,a=!1){let s=document.createElement("div");s.className="rounded-md border border-border p-3 flex flex-col gap-2";let i=document.createElement("p");i.className="text-sm font-medium",i.textContent=t;let l=document.createElement("code");if(l.className="block break-all text-sm",l.textContent=r||n.empty_value,s.append(i,l),a&&r){let o=document.createElement("button");o.type="button",o.className="inline-flex h-8 w-fit items-center justify-center rounded-md border border-border bg-background px-3 text-sm hover:bg-accent",o.textContent=n.copy_request_id,o.addEventListener("click",async()=>{try{await navigator.clipboard.writeText(r),o.textContent=n.copied}catch{o.textContent=n.copy_failed}}),s.append(o)}e.append(s)}var S={id:"request",title:"Request",mount(e,t){let r=t.i18n.request;e.replaceChildren();let n=document.createElement("p");return n.className="text-sm text-muted-foreground",n.textContent=r.intro,e.append(n),h(e,r.request_id_label,t.requestId,r,!0),h(e,r.path_label,t.path,r),h(e,r.locale_label,t.locale,r),()=>{}}};var q=`
:host {
  position: fixed;
  right: 1rem;
  bottom: 1rem;
  z-index: 30;
  pointer-events: none;
  color-scheme: light dark;
  font-family:
    ui-sans-serif,
    system-ui,
    -apple-system,
    BlinkMacSystemFont,
    "Segoe UI",
    sans-serif;
  font-size: 14px;
  line-height: 1.5;
}

* {
  box-sizing: border-box;
}

button {
  font: inherit;
  cursor: pointer;
}

[hidden],
.hidden {
  display: none !important;
}

.launcher {
  pointer-events: auto;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 2.5rem;
  height: 2.5rem;
  border: 1px solid hsl(214 32% 91%);
  border-radius: 9999px;
  background: hsl(0 0% 100%);
  color: hsl(222 47% 11%);
  box-shadow: 0 10px 30px rgb(15 23 42 / 0.16);
  transition:
    background-color 120ms ease,
    color 120ms ease,
    box-shadow 120ms ease;
}

.launcher:hover {
  background: hsl(210 40% 96%);
}

.launcher:focus-visible,
.button:focus-visible,
.tab:focus-visible {
  outline: 2px solid hsl(217 91% 60%);
  outline-offset: 2px;
}

.panel {
  pointer-events: auto;
  position: absolute;
  right: 0;
  bottom: calc(100% + 0.5rem);
  width: min(92vw, 420px);
  max-height: min(80vh, 640px);
  overflow: auto;
  border: 1px solid hsl(214 32% 91%);
  border-radius: 0.75rem;
  background: hsl(0 0% 100%);
  color: hsl(222 47% 11%);
  padding: 1rem;
  box-shadow:
    0 20px 25px -5px rgb(15 23 42 / 0.16),
    0 8px 10px -6px rgb(15 23 42 / 0.12);
}

.mobile-reload {
  pointer-events: auto;
  position: absolute;
  right: 0;
  bottom: calc(100% + 0.5rem);
  display: flex;
  width: min(88vw, 18rem);
  flex-direction: column;
  gap: 0.75rem;
  border: 1px solid hsl(214 32% 91%);
  border-radius: 0.75rem;
  background: hsl(0 0% 100%);
  color: hsl(222 47% 11%);
  padding: 0.875rem;
  box-shadow:
    0 20px 25px -5px rgb(15 23 42 / 0.16),
    0 8px 10px -6px rgb(15 23 42 / 0.12);
}

.stack,
.flex-col {
  display: flex;
  flex-direction: column;
}

.gap-1 {
  gap: 0.25rem;
}

.gap-2 {
  gap: 0.5rem;
}

.gap-4 {
  gap: 1rem;
}

.w-full {
  width: 100%;
}

.flex,
.inline-flex {
  display: flex;
}

.inline-flex {
  display: inline-flex;
}

.items-center {
  align-items: center;
}

.justify-between {
  justify-content: space-between;
}

.flex-wrap {
  flex-wrap: wrap;
}

.rounded-md {
  border-radius: 0.375rem;
}

.border {
  border-width: 1px;
  border-style: solid;
}

.border-border {
  border-color: hsl(214 32% 91%);
}

.border-transparent {
  border-color: transparent;
}

.bg-background {
  background: hsl(0 0% 100%);
}

.bg-muted {
  background: hsl(210 40% 96%);
}

.bg-primary {
  background: hsl(222 47% 11%);
}

.bg-emerald-500 {
  background: #10b981;
}

.bg-pink-600 {
  background: #db2777;
}

.bg-amber-500 {
  background: #f59e0b;
}

.bg-sky-600 {
  background: #0284c7;
}

.text-primary-foreground,
.text-white {
  color: #fff;
}

.text-foreground {
  color: hsl(222 47% 11%);
}

.text-muted-foreground {
  color: hsl(215 16% 47%);
}

.text-destructive {
  color: #db2777;
}

.text-sm {
  font-size: 0.875rem;
  line-height: 1.25rem;
}

.text-xs {
  font-size: 0.75rem;
  line-height: 1rem;
}

.font-medium {
  font-weight: 500;
}

.h-8 {
  height: 2rem;
}

.w-fit {
  width: fit-content;
}

.px-2 {
  padding-left: 0.5rem;
  padding-right: 0.5rem;
}

.px-3 {
  padding-left: 0.75rem;
  padding-right: 0.75rem;
}

.py-0\\.5 {
  padding-top: 0.125rem;
  padding-bottom: 0.125rem;
}

.p-3 {
  padding: 0.75rem;
}

.break-all {
  word-break: break-all;
}

.block {
  display: block;
}

.button,
.tab {
  display: inline-flex;
  min-height: 2rem;
  align-items: center;
  justify-content: center;
  border: 1px solid hsl(214 32% 91%);
  border-radius: 0.375rem;
  background: hsl(0 0% 100%);
  color: hsl(222 47% 11%);
  padding: 0.375rem 0.75rem;
}

.button:hover,
.tab:hover {
  background: hsl(210 40% 96%);
}

.separator {
  width: 100%;
  height: 1px;
  border: 0;
  background: hsl(214 32% 91%);
}

h2,
p {
  margin: 0;
}

code {
  font-family:
    ui-monospace,
    SFMono-Regular,
    Menlo,
    Monaco,
    Consolas,
    "Liberation Mono",
    monospace;
}

@media (prefers-color-scheme: dark) {
  .launcher,
  .mobile-reload,
  .panel,
  .bg-background,
  .button,
  .tab {
    border-color: hsl(217 33% 18%);
    background: hsl(222 47% 11%);
    color: hsl(210 40% 98%);
  }

  .launcher:hover,
  .button:hover,
  .tab:hover,
  .bg-muted {
    background: hsl(217 33% 18%);
  }

  .text-foreground {
    color: hsl(210 40% 98%);
  }

  .text-muted-foreground {
    color: hsl(215 20% 65%);
  }

  .border-border,
  .separator {
    border-color: hsl(217 33% 18%);
  }

  .separator {
    background: hsl(217 33% 18%);
  }
}
`;var b="fastygo-dev-overlay",T=[M,D,S];function z(e){return T.find((t)=>t.id===e)}function G(e,t){for(let r of e.querySelectorAll("[data-dev-tab]")){let n=r.dataset.devTab===t;r.setAttribute("aria-selected",n?"true":"false"),r.classList.toggle("bg-primary",n),r.classList.toggle("text-primary-foreground",n)}for(let r of e.querySelectorAll("[data-dev-panel]")){let n=r.dataset.devPanel===t;r.toggleAttribute("hidden",!n),r.dataset.devPanelActive=n?"true":"false"}}function U(e,t){let r=[];for(let n of T){let a=e.querySelector(`#fastygo-dev-panel-${n.id}`);if(!a)continue;r.push(n.mount(a,t))}for(let n of e.querySelectorAll("[data-dev-tab]")){let a=()=>{let s=n.dataset.devTab;if(!s||!z(s))return;G(e,s)};n.addEventListener("click",a),r.push(()=>n.removeEventListener("click",a))}return()=>{for(let n of r)n()}}function Q(e){let t=e.querySelector("#fastygo-dev-launcher"),r=e.querySelector("#fastygo-dev-panel");if(!t||!r)return()=>{};let n=()=>{let a=r.classList.toggle("hidden")===!1;t.setAttribute("aria-expanded",a?"true":"false")};return t.addEventListener("click",n),()=>t.removeEventListener("click",n)}function J(){return y()}function K(){return window.matchMedia("(max-width: 767px)").matches}function d(e){return e.replaceAll("&","&amp;").replaceAll("<","&lt;").replaceAll(">","&gt;").replaceAll('"',"&quot;").replaceAll("'","&#39;")}function X(e,t){if(e.includes("%s"))return e.replace("%s",t);return`${e} ${t}`}function g(e,t,r){return`<button
    id="fastygo-dev-tab-${e}"
    class="tab${r?" bg-primary text-primary-foreground":""}"
    type="button"
    role="tab"
    aria-controls="fastygo-dev-panel-${e}"
    aria-selected="${r?"true":"false"}"
    data-dev-tab="${e}"
    data-dev-tab-label="${d(t)}"
  >${d(t)}</button>`}function x(e,t,r,n){return`<section
    class="gap-2"
    id="fastygo-dev-panel-wrap-${e}"
    role="tabpanel"
    aria-labelledby="fastygo-dev-tab-${e}"
    data-dev-panel="${e}"
    data-dev-panel-label="${d(t)}"
    data-dev-panel-active="${n?"true":"false"}"
    ${n?"":"hidden"}
  >
    <p class="text-sm text-muted-foreground">${d(X(r,t))}</p>
    <div id="fastygo-dev-panel-${e}" class="flex flex-col gap-2"></div>
  </section>`}function Z(){return`<svg
    width="20"
    height="20"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    aria-hidden="true"
  >
    <path d="M11 10.27 7 3.34"></path>
    <path d="m11 13.73-4 6.93"></path>
    <path d="M12 22v-2"></path>
    <path d="M12 2v2"></path>
    <path d="M14 12h8"></path>
    <path d="m17 20.66-1-1.73"></path>
    <path d="m17 3.34-1 1.73"></path>
    <path d="M2 12h2"></path>
    <path d="m20.66 17-1.73-1"></path>
    <path d="m20.66 7-1.73 1"></path>
    <path d="m3.34 17 1.73-1"></path>
    <path d="m3.34 7 1.73 1"></path>
    <circle cx="12" cy="12" r="2"></circle>
    <circle cx="12" cy="12" r="8"></circle>
  </svg>`}function W(e){return`
    <style>${q}</style>
    <aside
      id="fastygo-dev-mobile-reload"
      class="mobile-reload hidden"
      role="status"
      aria-live="polite"
    >
      <p>${d(e.mobile_reload_hint)}</p>
      <button
        id="fastygo-dev-mobile-reload-button"
        class="button w-fit"
        type="button"
      >${d(e.mobile_reload_button)}</button>
    </aside>
    <button
      id="fastygo-dev-launcher"
      class="launcher"
      type="button"
      aria-controls="fastygo-dev-panel"
      aria-expanded="false"
      aria-label="${d(e.launcher_aria_label)}"
    >${Z()}</button>
    <section
      id="fastygo-dev-panel"
      class="panel hidden"
      role="region"
      aria-label="${d(e.panel_aria_label)}"
    >
      <div class="stack gap-4 w-full">
        <h2>${d(e.title)}</h2>
        <p class="text-sm text-muted-foreground">${d(e.subtitle)}</p>
        <div class="flex flex-wrap gap-2" role="tablist" aria-label="${d(e.tablist_aria_label)}">
          ${g("health",e.tabs.health,!0)}
          ${g("assets",e.tabs.assets,!1)}
          ${g("request",e.tabs.request,!1)}
        </div>
        ${x("health",e.tabs.health,e.loading_tab,!0)}
        ${x("assets",e.tabs.assets,e.loading_tab,!1)}
        ${x("request",e.tabs.request,e.loading_tab,!1)}
        <hr class="separator" aria-hidden="true" />
        <button
          id="fastygo-dev-hide"
          class="button w-fit"
          type="button"
          aria-label="${d(e.hide_button_aria_label)}"
        >${d(e.hide_button_label)}</button>
      </div>
    </section>
  `}class A extends HTMLElement{cleanup=null;connectedCallback(){if(this.shadowRoot)return;let e=J();if(!e)return;let t=this.attachShadow({mode:"open"});t.innerHTML=W(e);let r={requestId:_(),path:P(),locale:k(e.request.locale_missing),i18n:e,fetchStatus:()=>w(e),disableOverlay:()=>E(e)},n=Q(t),a=U(t,r),s=t.querySelector("#fastygo-dev-hide"),i=t.querySelector("#fastygo-dev-mobile-reload"),l=t.querySelector("#fastygo-dev-mobile-reload-button"),o=window.matchMedia("(max-width: 767px)"),c=()=>{r.disableOverlay()},p=()=>{window.location.reload()},v=(I)=>{i?.classList.toggle("hidden",!I.matches)};s?.addEventListener("click",c),l?.addEventListener("click",p),o.addEventListener("change",v),i?.classList.toggle("hidden",!o.matches),this.cleanup=()=>{s?.removeEventListener("click",c),l?.removeEventListener("click",p),o.removeEventListener("change",v),n(),a()}}disconnectedCallback(){this.cleanup?.(),this.cleanup=null}}function f(){if(K())return;if(!customElements.get(b))customElements.define(b,A);if(!document.querySelector(b))document.body.append(document.createElement(b))}function ee(){if(document.readyState==="loading"){document.addEventListener("DOMContentLoaded",f,{once:!0});return}f()}globalThis.__FASTYGO_DEV_OVERLAY_MOUNT__=f;ee();window.addEventListener("pageshow",f);})();
