package main

const visualizerJS = `
const canvas = document.getElementById('stage');
const ctx = canvas.getContext('2d');
function resize() {
  canvas.width = canvas.clientWidth;
  canvas.height = canvas.clientHeight;
}
window.addEventListener('resize', resize);
resize();

const AMBER = '#FFB347';
const CYAN = '#2EE6D6';

const xs = COLONY.Rooms.map(r => r.X), ys = COLONY.Rooms.map(r => r.Y);
const minX = Math.min(...xs), maxX = Math.max(...xs);
const minY = Math.min(...ys), maxY = Math.max(...ys);
const pad = 80;
function pos(r) {
  const x = pad + (canvas.width - 2*pad) * (maxX === minX ? 0.5 : (r.X - minX) / (maxX - minX));
  const y = pad + (canvas.height - 2*pad) * (maxY === minY ? 0.5 : (r.Y - minY) / (maxY - minY));
  return {x, y};
}
const roomPos = {};
const roomInfo = {};
COLONY.Rooms.forEach(r => { roomPos[r.Name] = pos(r); roomInfo[r.Name] = r; });

let ants = {};          // antID -> {x,y,room}
let roomPulse = {};     // roomName -> {color, intensity}
let turnIndex = 0;
let moveCount = 0;
let playing = false;

function log(msg) {
  const el = document.getElementById('termLog');
  el.innerHTML += '<div>&gt; ' + msg + '</div>';
  el.scrollTop = el.scrollHeight;
}

function pulse(room, color) {
  roomPulse[room] = { color, intensity: 1 };
}

function decayPulses() {
  Object.keys(roomPulse).forEach(r => {
    roomPulse[r].intensity -= 0.04;
    if (roomPulse[r].intensity <= 0) delete roomPulse[r];
  });
}

function draw() {
  ctx.clearRect(0,0,canvas.width,canvas.height);

  // tunnels
  ctx.strokeStyle = 'rgba(180,140,90,0.35)';
  ctx.lineWidth = 1.5;
  COLONY.Links.forEach(l => {
    const a = roomPos[l.From], b = roomPos[l.To];
    if (!a || !b) return;
    ctx.beginPath(); ctx.moveTo(a.x,a.y); ctx.lineTo(b.x,b.y); ctx.stroke();
  });

  // rooms (square nodes, per spec)
  COLONY.Rooms.forEach(r => {
    const p = roomPos[r.Name];
    const size = r.IsStart || r.IsEnd ? 26 : 18;
    const pulseInfo = roomPulse[r.Name];

    ctx.save();
    if (pulseInfo) {
      ctx.shadowColor = pulseInfo.color;
      ctx.shadowBlur = 20 * pulseInfo.intensity;
    }
    ctx.fillStyle = r.IsStart || r.IsEnd ? '#1c1c1c' : '#131313';
    ctx.strokeStyle = pulseInfo ? pulseInfo.color : (r.IsStart || r.IsEnd ? AMBER : 'rgba(255,255,255,0.25)');
    ctx.lineWidth = r.IsStart || r.IsEnd ? 2 : 1;
    roundRect(ctx, p.x - size/2, p.y - size/2, size, size, 4);
    ctx.fill(); ctx.stroke();
    ctx.restore();

    ctx.fillStyle = 'rgba(220,235,240,0.85)';
    ctx.font = '11px monospace';
    ctx.textAlign = 'center';
    ctx.fillText(r.Name, p.x, p.y + size/2 + 16);
  });

  // ants (glowing amber circles)
  Object.entries(ants).forEach(([id, a]) => {
    ctx.save();
    ctx.shadowColor = AMBER;
    ctx.shadowBlur = 12;
    ctx.beginPath();
    ctx.arc(a.x, a.y, 5, 0, Math.PI*2);
    ctx.fillStyle = AMBER;
    ctx.fill();
    ctx.restore();
    ctx.fillStyle = AMBER;
    ctx.font = '9px monospace';
    ctx.fillText('L' + id, a.x + 8, a.y - 8);
  });

  decayPulses();
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
  turn.Moves.forEach(m => {
    const from = ants[m.Ant] ? ants[m.Ant].room : COLONY.Rooms.find(r => r.IsStart).Name;
    const target = roomPos[m.Room];
    const targetInfo = roomInfo[m.Room];

    ants[m.Ant] = { x: target.x, y: target.y, room: m.Room };
    moveCount++;

    // departure pulse (cyan) on the room the ant just left
    if (from) pulse(from, CYAN);

    // arrival glow (amber) at ##end
    if (targetInfo && targetInfo.IsEnd) {
      pulse(m.Room, AMBER);
    }

    log('Step ' + (turnIndex+1) + ': L' + m.Ant + '\u2192' + m.Room);
  });
  document.getElementById('moveCount').textContent = moveCount;
}

function stepForward() {
  if (turnIndex >= COLONY.Turns.length) { playing = false; return; }
  applyTurn(COLONY.Turns[turnIndex]);
  turnIndex++;
}

function reset() {
  ants = {};
  roomPulse = {};
  turnIndex = 0;
  moveCount = 0;
  playing = false;
  document.getElementById('termLog').innerHTML = '';
  document.getElementById('moveCount').textContent = 0;
}

document.getElementById('antCount').textContent = COLONY.NumAnts;
document.getElementById('stepBtn').onclick = stepForward;
document.getElementById('playBtn').onclick = () => playing = !playing;
document.getElementById('resetBtn').onclick = reset;

function loop() {
  if (playing) stepForward();
  draw();
  requestAnimationFrame(() => setTimeout(loop, parseInt(document.getElementById('speed').value)));
}
draw();
loop();
`