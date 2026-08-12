package main

const visualizerCSS = `
:root {
  --amber: #FFB347;
  --cyan: #00FFFF;
  --offwhite: #F5F5F5;
  --dim: #A3A3A3;
  --dark-log: #737373;
  --tunnel: #393939;
  --charcoal-low: #1c1b1b;
  --radius: 4px;
  --font-ui: 'Hanken Grotesk', sans-serif;
  --font-mono: 'JetBrains Mono', monospace;
}
* { box-sizing: border-box; }
body { margin: 0; background: #000; font-family: var(--font-ui); color: var(--offwhite); overflow: hidden; }
#bg {
  position: fixed; inset: 0;
  background-size: cover; background-position: center;
  opacity: 0.6; z-index: -2;
}
#bgScrim { position: fixed; inset: 0; background: rgba(14,14,14,0.4); z-index: -1; }
#topbar {
  position: absolute; top: 0; left: 0; right: 0;
  display: flex; justify-content: space-between; align-items: center;
  padding: 12px 24px;
  background: rgba(14,14,14,0.5);
  backdrop-filter: blur(12px); -webkit-backdrop-filter: blur(12px);
  border-bottom: 1px solid rgba(255,179,71,0.25);
  font-family: var(--font-mono); font-size: 12px; font-weight: 700; letter-spacing: 0.05em;
  color: var(--amber);
}
#topbar .sys { display: flex; align-items: center; gap: 8px; }
#topbar .filename {
  background: var(--charcoal-low);
  border: 1px solid rgba(255,179,71,0.25);
  border-radius: var(--radius);
  padding: 4px 10px; color: var(--offwhite); font-weight: 500;
}
#hud {
  position: absolute; top: 68px; left: 24px;
  background: rgba(20,20,19,0.55);
  backdrop-filter: blur(12px); -webkit-backdrop-filter: blur(12px);
  border: 1px solid rgba(255,179,71,0.25);
  border-radius: var(--radius);
  padding: 10px 16px;
  font-family: var(--font-mono); font-size: 14px; font-weight: 500; color: var(--amber);
}
#hud div { margin: 4px 0; }
#hud .value { color: var(--offwhite); }
#stage { display: block; width: 100vw; height: 100vh; background: transparent; }
#term {
  position: absolute; bottom: 90px; right: 24px; width: 300px; height: 150px;
  background: rgba(28,27,27,0.8);
  backdrop-filter: blur(12px); -webkit-backdrop-filter: blur(12px);
  border: 1px solid var(--tunnel);
  border-radius: var(--radius);
  padding: 10px 12px;
  font-family: var(--font-mono); font-size: 12px;
  overflow: hidden;
  display: flex; flex-direction: column;
}
#term .termHeader { display: flex; align-items: center; justify-content: space-between; margin-bottom: 6px; }
#term .label { color: var(--amber); font-size: 11px; font-weight: 700; letter-spacing: 0.05em; }
#term .rec { display: flex; align-items: center; gap: 4px; font-size: 9px; color: var(--dim); }
#term .recDot { width: 6px; height: 6px; border-radius: 50%; background: #ff4d4d; animation: recPulse 1.4s infinite; }
@keyframes recPulse { 0%,100% { opacity: 1; } 50% { opacity: 0.25; } }
#termLog { flex: 1; overflow: hidden; display: flex; flex-direction: column; justify-content: flex-end; }
#termLog div { margin: 2px 0; opacity: 0; animation: fadeIn 0.4s forwards; }
@keyframes fadeIn { to { opacity: 1; } }
#termLog div.step { color: var(--amber); font-weight: 700; }
#termLog div.move { color: var(--offwhite); }
#controls {
  position: absolute; bottom: 24px; left: 50%; transform: translateX(-50%);
  display: flex; gap: 16px; align-items: center;
  background: rgba(20,20,19,0.55);
  backdrop-filter: blur(12px); -webkit-backdrop-filter: blur(12px);
  border: 1px solid rgba(255,179,71,0.25);
  border-radius: 8px; padding: 10px 20px;
}
button {
  background: transparent; border: 1px solid rgba(255,179,71,0.3); color: var(--offwhite);
  padding: 8px 14px; border-radius: var(--radius); cursor: pointer;
  font-family: var(--font-ui); font-weight: 500; font-size: 13px;
  transition: transform 0.1s ease;
}
button:active { transform: scale(0.95); }
button:hover { border-color: var(--amber); background: rgba(255,179,71,0.08); }
#playBtn { background: var(--amber); color: #2a1800; font-weight: 700; border: none; }
#playBtn:hover { background: #ffc46b; }
#playBtn.active { background: rgba(255,179,71,0.2); color: var(--amber); border: 1px solid var(--amber); }
input[type=range] { accent-color: var(--amber); }
`