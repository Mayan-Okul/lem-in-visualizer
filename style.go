package main

const visualizerCSS = `
:root {
  --amber: #FFB347;
  --teal: #2A4D4D;
  --teal-light: #A8CECD;
  --beige: #D2B48C;
  --charcoal: #1A1A1A;
  --charcoal-container: #20201f;
  --charcoal-low: #1c1b1b;
  --text: #e5e2e1;
  --radius: 0.25rem;
  --font-ui: 'Hanken Grotesk', sans-serif;
  --font-mono: 'JetBrains Mono', monospace;
}

* { box-sizing: border-box; }

body {
  margin: 0;
  background: var(--charcoal);
  font-family: var(--font-ui);
  color: var(--text);
  overflow: hidden;
}

#bg {
  position: fixed;
  inset: 0;
  background-size: cover;
  background-position: center;
  filter: brightness(0.7) saturate(1.05);
  transform: scale(1.05);
  z-index: -2;
}

#bgOverlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.4);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  z-index: -1;
}

/* top nav */
#topbar {
  position: absolute;
  top: 0; left: 0; right: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 24px;
  background: rgba(19,19,19,0.6);
  border-bottom: 1px solid var(--teal);
  font-family: var(--font-mono);
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.05em;
  color: var(--amber);
}
#topbar .filename {
  background: var(--charcoal-low);
  border: 1px solid var(--teal);
  border-radius: var(--radius);
  padding: 4px 10px;
  color: var(--text);
  font-weight: 500;
}

/* stats card */
#hud {
  position: absolute;
  top: 68px;
  left: 24px;
  background: var(--charcoal-container);
  border: 1px solid var(--teal);
  border-radius: var(--radius);
  padding: 10px 16px;
  font-family: var(--font-mono);
  font-size: 14px;
  font-weight: 500;
}
#hud div { margin: 4px 0; }
#hud .amber { color: var(--amber); }

#stage {
  display: block;
  width: 100vw;
  height: 100vh;
  background: transparent;
}

/* terminal log */
#term {
  position: absolute;
  bottom: 90px;
  right: 24px;
  width: 300px;
  height: 150px;
  background: var(--charcoal-low);
  border: 1px solid var(--teal);
  border-radius: var(--radius);
  padding: 10px 12px;
  font-family: var(--font-mono);
  font-size: 12px;
  overflow-y: auto;
}
#term::-webkit-scrollbar { width: 4px; }
#term::-webkit-scrollbar-thumb { background: var(--teal-light); border-radius: 2px; }
#term .label {
  color: var(--amber);
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.05em;
  margin-bottom: 6px;
}
#termLog div { color: var(--beige); margin: 3px 0; }
#termLog div.step { color: var(--teal-light); }

/* control bar */
#controls {
  position: absolute;
  bottom: 24px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 16px;
  align-items: center;
  background: var(--charcoal-container);
  border: 1px solid var(--teal);
  border-radius: 0.5rem;
  padding: 10px 20px;
}
button {
  background: transparent;
  border: 1px solid var(--teal);
  color: var(--text);
  padding: 8px 14px;
  border-radius: var(--radius);
  cursor: pointer;
  font-family: var(--font-ui);
  font-weight: 500;
  font-size: 13px;
}
button:hover { border-color: var(--teal-light); background: rgba(168,206,205,0.08); }
#playBtn { background: var(--amber); color: #462a00; font-weight: 700; border: none; }
#playBtn:hover { background: #ffc46b; }

input[type=range] { accent-color: var(--amber); }
`