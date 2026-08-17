package main

const visualizerJS = `
const svgNS = "http://www.w3.org/2000/svg";
const linkSvg = document.getElementById('linkSvg');
const nodeLayer = document.getElementById('nodeLayer');
const antLayer = document.getElementById('antLayer');

const xs = COLONY.Rooms.map(r => r.X), ys = COLONY.Rooms.map(r => r.Y);
const minX = Math.min(...xs), maxX = Math.max(...xs);
const minY = Math.min(...ys), maxY = Math.max(...ys);
const PAD = 12;

function pct(r) {
  const px = PAD + (100 - 2*PAD) * (maxX === minX ? 0.5 : (r.X - minX) / (maxX - minX));
  const py = PAD + (100 - 2*PAD) * (maxY === minY ? 0.5 : (r.Y - minY) / (maxY - minY));
  return { x: px, y: py };
}
const roomPct = {};
COLONY.Rooms.forEach(r => { roomPct[r.Name] = pct(r); });
const startName = COLONY.Rooms.find(r => r.IsStart).Name;
const endName = COLONY.Rooms.find(r => r.IsEnd).Name;

COLONY.Links.forEach(l => {
  const a = roomPct[l.From], b = roomPct[l.To];
  if (!a || !b) return;
  const line = document.createElementNS(svgNS, 'line');
  line.setAttribute('x1', a.x + '%'); line.setAttribute('y1', a.y + '%');
  line.setAttribute('x2', b.x + '%'); line.setAttribute('y2', b.y + '%');
  line.setAttribute('stroke', '#524535'); line.setAttribute('stroke-width', '2');
  linkSvg.appendChild(line);
});

const nodeEls = {};
COLONY.Rooms.forEach(r => {
  const p = roomPct[r.Name];
  const isSE = r.IsStart || r.IsEnd;
  const div = document.createElement('div');
  div.className = isSE
    ? 'absolute w-12 h-12 backdrop-blur-sm rounded-full border-2 flex items-center justify-center transform -translate-x-1/2 -translate-y-1/2 ' +
      (r.IsStart ? 'border-secondary node-idle-start' : 'border-primary node-idle-end')
    : 'absolute w-8 h-8 bg-surface-container/80 backdrop-blur-sm rounded-full border border-outline-variant flex items-center justify-center transform -translate-x-1/2 -translate-y-1/2';
  div.style.left = p.x + '%';
  div.style.top = p.y + '%';
  const label = document.createElement('span');
  label.className = 'font-data-mono text-xs absolute -bottom-6 whitespace-nowrap bg-black/50 px-1 rounded ' +
    (r.IsStart ? 'text-secondary' : r.IsEnd ? 'text-primary' : 'text-on-surface-variant');
  label.textContent = r.IsStart ? '##start' : r.IsEnd ? '##end' : r.Name;
  div.appendChild(label);
  nodeLayer.appendChild(div);
  nodeEls[r.Name] = div;
});

function pulseNode(name, isStart) {
  const el = nodeEls[name];
  if (!el) return;
  el.classList.remove(isStart ? 'node-idle-start' : 'node-idle-end');
  el.classList.add(isStart ? 'node-pulse-start' : 'node-pulse-end');
  setTimeout(() => {
    el.classList.remove(isStart ? 'node-pulse-start' : 'node-pulse-end');
    el.classList.add(isStart ? 'node-idle-start' : 'node-idle-end');
  }, 600);
}

let antEls = {};
let turnIndex = 0, moveCount = 0, arrivedCount = 0, playing = false;

function log(msg, kind) {
  const el = document.getElementById('termLog');
  const div = document.createElement('div');
  div.className = kind === 'step' ? 'text-secondary' : kind === 'move' ? 'text-primary' : 'text-outline';
  div.textContent = msg;
  el.appendChild(div);
  while (el.children.length > 6) el.removeChild(el.firstChild);
}

function ensureAnt(id) {
  if (antEls[id]) return antEls[id];
  const dot = document.createElement('div');
  dot.className = 'absolute w-3 h-3 bg-primary rounded-full ant-glow ant-dot transform -translate-x-1/2 -translate-y-1/2';
  const label = document.createElement('span');
  label.className = 'font-data-mono text-[10px] text-primary absolute -top-4 -left-1';
  label.textContent = 'L' + id;
  dot.appendChild(label);
  antLayer.appendChild(dot);
  antEls[id] = dot;
  return dot;
}

function applyTurn(turn) {
  log('> Step ' + (turnIndex + 1), 'step');
  const moves = [];
  turn.Moves.forEach(m => {
    const fromRoom = antEls[m.Ant] ? antEls[m.Ant].dataset.room : startName;
    const dot = ensureAnt(m.Ant);
    const target = roomPct[m.Room];
    dot.style.left = target.x + '%';
    dot.style.top = target.y + '%';
    dot.style.opacity = '1';
    dot.dataset.room = m.Room;
    moveCount++;
    if (fromRoom === startName) pulseNode(startName, true);
    if (m.Room === endName) {
      pulseNode(endName, false);
      setTimeout(() => {
        dot.style.opacity = '0';
        arrivedCount++;
        document.getElementById('antCount').textContent = arrivedCount + '/' + COLONY.NumAnts;
      }, 850);
    }
    moves.push('L' + m.Ant + '-' + m.Room);
  });
  log(moves.join(' '), 'move');
  document.getElementById('moveCount').textContent = moveCount;
}

function stepForward() {
  if (turnIndex >= COLONY.Turns.length) { playing = false; return; }
  applyTurn(COLONY.Turns[turnIndex]);
  turnIndex++;
}

function reset() {
  Object.values(antEls).forEach(el => el.remove());
  antEls = {}; turnIndex = 0; moveCount = 0; arrivedCount = 0; playing = false;
  document.getElementById('termLog').innerHTML = '';
  document.getElementById('moveCount').textContent = 0;
  document.getElementById('antCount').textContent = '0/' + COLONY.NumAnts;
}

document.getElementById('antCount').textContent = '0/' + COLONY.NumAnts;
log('Initializing simulation...', 'idle');
log('Ants ready: ' + COLONY.NumAnts, 'idle');
document.getElementById('stepBtn').onclick = stepForward;
document.getElementById('playBtn').onclick = () => { playing = !playing; };
document.getElementById('resetBtn').onclick = reset;

let lastTick = 0;
function loop(now) {
  const speed = parseInt(document.getElementById('speed').value);
  if (playing && now - lastTick > speed) { stepForward(); lastTick = now; }
  requestAnimationFrame(loop);
}
requestAnimationFrame(loop);
`