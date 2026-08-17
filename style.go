package main

const visualizerHead = `
<script src="https://cdn.tailwindcss.com?plugins=forms,container-queries"></script>
<link href="https://fonts.googleapis.com/css2?family=Hanken+Grotesk:wght@400;600;700&family=JetBrains+Mono:wght@500;700&display=swap" rel="stylesheet"/>
<link href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:wght,FILL@100..700,0..1&display=swap" rel="stylesheet"/>
<script>
tailwind.config = {
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        "on-surface-variant": "#d6c3b0", "on-background": "#e5e2e1",
        "primary": "#ffb347", "surface": "#131313",
        "outline": "#9f8e7c", "outline-variant": "#524535",
        "surface-container": "#20201f", "surface-container-high": "#2a2a2a",
        "secondary": "#a8cecd", "background": "#131313",
        "primary-container": "#ffb347", "on-primary-container": "#704700"
      },
      borderRadius: { DEFAULT: "0.125rem", lg: "0.25rem", xl: "0.5rem", full: "0.75rem" },
      spacing: { "sidebar-width": "320px", "unit-base": "8px", "gutter": "16px", "canvas-margin": "24px" },
      fontFamily: {
        "label-caps": ["JetBrains Mono"], "headline-md": ["Hanken Grotesk"],
        "headline-lg": ["Hanken Grotesk"], "body-md": ["Hanken Grotesk"], "data-mono": ["JetBrains Mono"]
      },
      fontSize: {
        "label-caps": ["12px", { lineHeight: "16px", letterSpacing: "0.05em", fontWeight: "700" }],
        "headline-md": ["24px", { lineHeight: "32px", fontWeight: "600" }],
        "headline-lg": ["32px", { lineHeight: "40px", letterSpacing: "-0.02em", fontWeight: "700" }],
        "body-md": ["16px", { lineHeight: "24px", fontWeight: "400" }],
        "data-mono": ["14px", { lineHeight: "20px", fontWeight: "500" }]
      }
    }
  }
}
</script>
`

const visualizerCSS = `
body { background-color: #1a1a1a; }
.glass-panel { background: rgba(42, 42, 42, 0.6); backdrop-filter: blur(8px); border: 1px solid rgba(82, 69, 53, 0.5); }
.node-glow { box-shadow: 0 0 15px rgba(255, 179, 71, 0.4); }
.ant-glow { box-shadow: 0 0 10px rgba(255, 179, 71, 0.9); animation: pulse-glow 1s infinite alternate; }
@keyframes pulse-glow { 0% { box-shadow: 0 0 5px rgba(255, 179, 71, 0.6); } 100% { box-shadow: 0 0 15px rgba(255, 179, 71, 1); } }
.bg-grid-pattern {
  background-image: linear-gradient(to right, rgba(168, 206, 205, 0.05) 1px, transparent 1px),
                     linear-gradient(to bottom, rgba(168, 206, 205, 0.05) 1px, transparent 1px);
  background-size: 40px 40px;
}
.node-idle-start { box-shadow: 0 0 10px rgba(168, 206, 205, 0.4); background-color: rgba(32,32,31,0.8); transition: box-shadow 0.6s ease, background-color 0.6s ease; }
.node-pulse-start { box-shadow: 0 0 25px 8px rgba(168, 206, 205, 0.9); background-color: rgba(168, 206, 205, 0.4); transition: box-shadow 0.1s ease, background-color 0.1s ease; }
.node-idle-end { box-shadow: 0 0 15px rgba(255, 179, 71, 0.4); background-color: rgba(32,32,31,0.8); transition: box-shadow 0.6s ease, background-color 0.6s ease; }
.node-pulse-end { box-shadow: 0 0 25px 8px rgba(255, 179, 71, 0.9); background-color: rgba(255, 179, 71, 0.4); transition: box-shadow 0.1s ease, background-color 0.1s ease; }
.ant-dot { transition: left 0.85s linear, top 0.85s linear, opacity 0.2s ease; }
body { min-height: max(884px, 100dvh); }
`