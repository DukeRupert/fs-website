/** @type {import('tailwindcss').Config}*/
const config = {
  content: ["./src/**/*.{html,js,svelte,ts}"],

  theme: {
    fontSize: {
      xs: ["0.75rem", { lineHeight: "1rem" }],
      sm: ["0.875rem", { lineHeight: "1.5rem" }],
      base: ["1rem", { lineHeight: "1.75rem" }],
      lg: ["1.125rem", { lineHeight: "1.75rem" }],
      xl: ["1.25rem", { lineHeight: "2rem" }],
      "2xl": ["1.5rem", { lineHeight: "2.25rem" }],
      "3xl": ["1.75rem", { lineHeight: "2.25rem" }],
      "4xl": ["2rem", { lineHeight: "2.5rem" }],
      "5xl": ["2.5rem", { lineHeight: "3rem" }],
      "6xl": ["3rem", { lineHeight: "3.5rem" }],
      "7xl": ["4rem", { lineHeight: "4.5rem" }],
    },
    extend: {
      colors: {
        primary: {
          // Curious Blue
          50: "#f1f8fe",
          100: "#e2f0fc",
          200: "#bfe0f8",
          300: "#86c8f3",
          400: "#46acea",
          500: "#1e91d9",
          600: "#1072b9",
          700: "#0f5c95",
          800: "#104e7c",
          900: "#134267",
          950: "#0d2a44",
        },
        secondary: {
          // Downriver
          50: "#dadce2",
          100: "#ced1d9",
          200: "#c2c5cf",
          300: "#9da2b3",
          400: "#535d79",
          500: "#091740",
          600: "#08153a",
          700: "#071130",
          800: "#050e26",
          900: "#040b1f",
          950: "#091740",
        },
        tertiary: {
          // Bright-turquoise
          50: "#eefffa",
          100: "#c6fff3",
          200: "#8dffe7",
          300: "#4cfcdb",
          400: "#15e6c5",
          500: "#00cdb0",
          600: "#00a692",
          700: "#028375",
          800: "#07685f",
          900: "#0b564f",
          950: "#003531",
        },
      },
      borderRadius: {
        "4xl": "2.5rem",
      },
    },
  },

  plugins: [require("@tailwindcss/typography")],
};

module.exports = config;
