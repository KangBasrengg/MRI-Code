/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        dark: {
          bg: "#080b12",
          card: "rgba(15, 22, 35, 0.65)",
          border: "rgba(255, 255, 255, 0.12)",
          hover: "rgba(30, 41, 60, 0.75)",
        },
        accent: {
          cyan: "#00f2ff",
          blue: "#3b82f6",
          purple: "#8a2be2",
          emerald: "#10b981",
        }
      },
      fontFamily: {
        sans: ['"NovecentoSansWideDemiBold"', '"Inter"', '-apple-system', 'BlinkMacSystemFont', '"Segoe UI"', 'Roboto', 'sans-serif'],
        heading: ['"NovecentoSansWideDemiBold"', 'sans-serif'],
        mono: ['"JetBrains Mono"', '"Fira Code"', 'Consolas', 'monospace'],
      },
      animation: {
        'pulse-slow': 'pulse 4s cubic-bezier(0.4, 0, 0.6, 1) infinite',
        'glow': 'glow 3s ease-in-out infinite alternate',
      },
      keyframes: {
        glow: {
          '0%': { boxShadow: '0 0 15px -5px rgba(0, 242, 255, 0.25)' },
          '100%': { boxShadow: '0 0 25px 5px rgba(0, 242, 255, 0.5)' },
        }
      }
    },
  },
  plugins: [],
}
