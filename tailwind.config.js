/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./web/components/**/*.{templ,go}",
    "./web/routes/**/*.{templ,go}",
    "./handlers/**/*.go",
  ],
  theme: {
    extend: {
      colors: {
        "base-background": "#030712",
        "base-text": "#030712",
      },
    },
  },
  plugins: [],
};
