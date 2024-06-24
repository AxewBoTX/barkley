const { addIconSelectors } = require("@iconify/tailwind");

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./web/components/**/*.{templ,go}",
    "./web/routes/**/*.{templ,go}",
    "./handlers/**/*.go",
  ],
  daisyui: {
    themes: [
      {
        main: {
          primary: "#AE7AFF",
          "primary-content": "#000000",
          secondary: "#FAE8A4",
          "secondary-content": "#000000",
          accent: "#98E9AB",
          "accent-content": "#000000",
          neutral: "#5F646D",
          "neutral-content": "#FFFFFF",
          "base-100": "#161616",
          "base-200": "#D5DDDE",
          "base-300": "#3f3f3f",
          "base-content": "#FFFFFF",
          info: "#00B5DB",
          "info-content": "#000C11",
          success: "#00CD80",
          "success-content": "#000F06",
          warning: "#BA1A00",
          "warning-content": "#F7D6CF",
          error: "#CE003C",
          "error-content": "#FCD6D6",
          "--rounded-box": "1rem",
          "--rounded-btn": "0.5rem",
          "--rounded-badge": "1.9rem",
          "--animation-btn": "0.25s",
          "--animation-input": "0.2s",
          "--btn-focus-scale": "0.95",
          "--border-btn": "1px",
          "--tab-border": "1px",
          "--tab-radius": "0.5rem",
        },
      },
    ],
  },
  theme: {
    extend: {
      fontFamily: {},
      colors: {
        "base-background": "#000000",
      },
      boxShadow: {
        "neu-md": "4px 4px 0 #3f3f3f",
        "neu-lg": "6px 6px 0 #3f3f3f",
        "neu-xl": "10px 10px 0 #3f3f3f",
      },
    },
  },
  plugins: [require("daisyui"), addIconSelectors(["mdi"])],
};
