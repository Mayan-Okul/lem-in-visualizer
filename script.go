package main

const visualizerJS = `
const canvas = document.getElementById('stage');
const ctx = canvas.getContext('2d');
function resize() { canvas.width = canvas.clientWidth; canvas.height = canvas.clientHeight; }
window.addEventListener('resize', resize);
resize();

const AMBER = '#FFB347';
const CYAN = '#00FFFF';
const DIM = '#A3A3A3';
const TUNNEL = '#393939';

const xs = COLONY.Rooms.map(r => r.X), ys = COLONY.Rooms.map(r => r.Y);
const minX = Math.min(...xs), maxX = Math.max(...xs);
const minY = Math.min(...ys), maxY = Math.max(...ys);
const padX = 90, padY = 90;
function pos(r) {
  const x = padX + (canvas.width - 2*padX) * (maxX === minX ? 0.5 : (r.X - minX) / (maxX - minX));
  const y = padY + (canvas.height - 2*padY) * (maxY === minY ? 0.5 : (r.Y - minY) / (maxY - minY));
  return {x, y};
}
const roomPos = {};
const roomInfo = {};
COLONY.Rooms.forEach(r => { roomPos[r.Name] = pos(r); roomInfo[r.Name] = r; });
const startRoomName = COLONY.Rooms.find(r => r.IsStart).Name;
const endRoomName = COLONY.Rooms.find(r => r.IsEnd).Name;

let ants = {};
let startPulse = null;
let endPulse = null;
let turnIndex = 0;
let moveCount = 0;
let arrivedCount = 0;
let playing = false;
const HOP_DURATION = 900;
const MAX_LOG_LINES = 5;

function log(msg, isStep) {
  const el = document.getElementById('termLog');
  const div = document.createElement('div');
  div.className = isStep ? 'step' : 'move';
  div.textContent = (isStep ? '' : '  ') + msg;
  el.appendChild(div);
  while (el.children.length > MAX_LOG_LINES) el.removeChild(el.firstChild);
}

function pulseStart() { startPulse = { intensity: 1 }; }
function pulseEnd() { endPulse = { intensity: 1 }; }
function decay(p) {
  if (!p) return null;
  p.intensity -= 1000/600/60;
  return p.intensity <= 0 ? null : p;
}

function draw() {
  ctx.clearRect(0,0,canvas.width,canvas.height);

  ctx.strokeStyle = TUNNEL;
  ctx.lineWidth = 1;
  COLONY.Links.forEach(l => {
    const a = roomPos[l.From], b = roomPos[l.To];
    if (!a || !b) return;
    ctx.beginPath(); ctx.moveTo(a.x,a.y); ctx.lineTo(b.x,b.y); ctx.stroke();
  });

  COLONY.Rooms.forEach(r => {
    const p = roomPos[r.Name];
    const isStart = r.IsStart, isEnd = r.IsEnd;
    ctx.save();

    if (isStart || isEnd) {
      const size = 40;
      const glow = isStart ? startPulse : endPulse;
      const glowColor = isStart ? CYAN : AMBER;
      if (glow) { ctx.shadowColor = glowColor; ctx.shadowBlur = 30 * glow.intensity; }
      ctx.fillStyle = '#161616';
      ctx.strokeStyle = glow ? glowColor : AMBER;
      ctx.lineWidth = 2;
      roundRect(ctx, p.x - size/2, p.y - size/2, size, size, 4);
      ctx.fill(); ctx.stroke();
      ctx.restore();
      ctx.fillStyle = DIM;
      ctx.font = '11px "JetBrains Mono", monospace';
      ctx.textAlign = 'center';
      ctx.fillText(isStart ? '##start' : '##end', p.x, p.y + size/2 + 16);
    } else {
      const radius = 16;
      ctx.fillStyle = '#1a1a1a';
      ctx.strokeStyle = TUNNEL;
      ctx.lineWidth = 1.5;
      ctx.beginPath(); ctx.arc(p.x, p.y, radius, 0, Math.PI*2);
      ctx.fill(); ctx.stroke();
      ctx.restore();
      ctx.fillStyle = DIM;
      ctx.font = '11px "JetBrains Mono", monospace';
      ctx.textAlign = 'center';
      ctx.fillText(r.Name, p.x, p.y + radius + 16);
    }
  });

  Object.entries(ants).forEach(([id, a]) => {
    ctx.save();
    ctx.shadowColor = AMBER; ctx.shadowBlur = 10;
    ctx.beginPath(); ctx.arc(a.x, a.y, 6, 0, Math.PI*2);
    ctx.fillStyle = AMBER; ctx.fill();
    ctx.restore();
    ctx.fillStyle = AMBER;
    ctx.font = '9px "JetBrains Mono", monospace';
    ctx.fillText('L' + id, a.x + 9, a.y - 9);
  });

  startPulse = decay(startPulse);
  endPulse = decay(endPulse);
}

function roundRect(ctx, x, y, w, h, r) {
  ctx.beginPath();
  ctx.moveTo(x + r, y);
  ctx.arcTo(x + w, y, x + w, y + h, r);
  ctx.arcTo(x + w, y + h, x, y + h, r);
  ctx.arcTo(x, y + h, x, y, r);
  ctx.arcTo(x, y, x + w, y, r);
  ctx.closePath();
}

function applyTurn(turn) {
  log('Step ' + (turnIndex+1) + ':', true);
  const lineParts = [];
  turn.Moves.forEach(m => {
    const fromRoom = ants[m.Ant] ? ants[m.Ant].room : startRoomName;
    const fromPos = ants[m.Ant] ? { x: ants[m.Ant].x, y: ants[m.Ant].y } : roomPos[startRoomName];
    const target = roomPos[m.Room];
    const isArriving = m.Room === endRoomName;

    ants[m.Ant] = {
      x: fromPos.x, y: fromPos.y,
      startX: fromPos.x, startY: fromPos.y,
      targetX: target.x, targetY: target.y,
      startTime: performance.now(),
      room: m.Room,
      arriving: isArriving
    };
    moveCount++;

    if (fromRoom === startRoomName) pulseStart();
    if (isArriving) pulseEnd();

    lineParts.push('L' + m.Ant + '-' + m.Room);
  });
  log(lineParts.join(' '));
  document.getElementById('moveCount').textContent = moveCount;
}

function stepForward() {
  if (turnIndex >= COLONY.Turns.length) { playing = false; return; }
  applyTurn(COLONY.Turns[turnIndex]);
  turnIndex++;
}

function reset() {
  ants = {}; startPulse = null; endPulse = null; turnIndex = 0; moveCount = 0; arrivedCount = 0; playing = false;
  document.getElementById('termLog').innerHTML = '';
  document.getElementById('moveCount').textContent = 0;
  document.getElementById('antCount').textContent = 0;
  document.getElementById('playBtn').classList.remove('active');
}

document.getElementById('antTotal').textContent = COLONY.NumAnts;
document.getElementById('stepBtn').onclick = stepForward;
document.getElementById('playBtn').onclick = () => {
  playing = !playing;
  document.getElementById('playBtn').classList.toggle('active', playing);
};
document.getElementById('resetBtn').onclick = reset;

let lastTurnTick = 0;
function animationLoop(now) {
  Object.entries(ants).forEach(([id, a]) => {
    const t = Math.min(1, (now - a.startTime) / HOP_DURATION);
    a.x = a.startX + (a.targetX - a.startX) * t;
    a.y = a.startY + (a.targetY - a.startY) * t;
    if (t >= 1 && a.arriving) {
      arrivedCount++;
      document.getElementById('antCount').textContent = arrivedCount;
      delete ants[id];
    }
  });
  draw();
  const speed = parseInt(document.getElementById('speed').value);
  if (playing && now - lastTurnTick > speed) { stepForward(); lastTurnTick = now; }
  requestAnimationFrame(animationLoop);
}
requestAnimationFrame(animationLoop);
`