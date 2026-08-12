package main

const visualizerCSS = `
:root {
  --amber: #FFB347;
  --cyan: #2EE6D6;
  --surface: #131313;
  --surface-soft: rgba(19,19,19,0.75);
  --text: #cfefff;
  --radius: 4px;
}

* { box-sizing: border-box; }

body {
  margin: 0;
  background: #05080c;
  font-family: 'Courier New', monospace;
  color: var(--text);
  overflow: hidden;
}

#bg {
  position: fixed;
  inset: 0;
  background-size: cover;
  background-position: center;
  filter: brightness(0.75) saturate(1.1);
  transform: scale(1.05);
  z-index: -2;
}

#bgOverlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.4);
  backdrop-filter: blur(6px);
  -webkit-backdrop-filter: blur(6px);
  z-index: -1;
}

/* top nav */
#topbar {
  position: absolute;
  top: 0; left: 0; right: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 20px;
  background: rgba(10,10,10,0.5);
  border-bottom: 1px solid rgba(255,179,71,0.25);
  font-size: 12px;
  letter-spacing: 1px;
  color: var(--amber);
}
#topbar .filename {
  background: var(--surface-soft);
  border: 1px solid rgba(255,255,255,0.15);
  border-radius: var(--radius);
  padding: 4px 10px;
  color: var(--text);
}

/* stats card */
#hud {
  position: absolute;
  top: 60px;
  left: 20px;
  background: var(--surface-soft);
  border: 1px solid rgba(46,230,214,0.3);
  border-radius: var(--radius);
  padding: 10px 16px;
  font-size: 13px;
  backdrop-filter: blur(4px);
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
  right: 20px;
  width: 300px;
  height: 150px;
  background: var(--surface-soft);
  border: 1px solid rgba(255,179,71,0.3);
  border-radius: var(--radius);
  padding: 8px 10px;
  font-size: 11px;
  overflow-y: auto;
  backdrop-filter: blur(4px);
}
#term .label {
  color: var(--amber);
  font-size: 10px;
  letter-spacing: 1px;
  margin-bottom: 4px;
}
#termLog div { color: var(--cyan); margin: 2px 0; }

/* control bar */
#controls {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 16px;
  align-items: center;
  background: var(--surface-soft);
  border: 1px solid rgba(255,255,255,0.15);
  border-radius: var(--radius);
  padding: 10px 20px;
  backdrop-filter: blur(6px);
}
button {
  background: transparent;
  border: 1px solid rgba(255,179,71,0.4);
  color: var(--text);
  padding: 8px 14px;
  border-radius: var(--radius);
  cursor: pointer;
  font-family: inherit;
  font-size: 12px;
}
button:hover { background: rgba(255,179,71,0.15); }
#playBtn { background: var(--amber); color: #131313; font-weight: bold; }
#playBtn:hover { background: #ffc46b; }

input[type=range] { accent-color: var(--amber); }
`