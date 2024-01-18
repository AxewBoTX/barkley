/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx,svelte}"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: [
      {
        main: {
          primary: "#0061ff",
          secondary: "#0079ff",
          accent: "#00ddbe",
          neutral: "#1f2625",
          "base-100": "#202528",
          info: "#20afff",
          success: "#009a5b",
          warning: "#fa8100",
          error: "#ff003e",
        },
      },
    ],
  },
  plugins: [require("daisyui")],
};
