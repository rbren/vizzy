/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./public/**/*.{js,jsx,ts,tsx}",
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  safelist: [ { pattern: /alert-+/ } ],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  "colors": {
    "brand": {
      50: "#EFEEF7",
      100: "#DDD9ED",
      200: "#BEB7DC",
      300: "#9C91C9",
      400: "#7D6FB9",
      500: "#5E4FA2",
      600: "#4C4082",
      700: "#382F60",
      800: "#262041",
      900: "#120F1F",
      950: "#0A0811"
    }
  },
  daisyui: {
    themes: [
      {
        myDataTheme: {                           // Custom theme name
          "primary": "#5E4FA2",
          "secondary": "#AFAFD0",
          "accent": "#222222",
          "neutral": "#3F3F3F",
          "base-100": "#222222",
          "info": "#22d3ee",
          "success": "#a3e635",
          "warning": "#facc15",
          "error": "#C62828",
        },
      },
    ],
  },
}

