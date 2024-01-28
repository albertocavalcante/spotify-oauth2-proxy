/** @type {import('tailwindcss').Config} */

const colors = require('tailwindcss/colors')

module.exports = {
  content: [
    "./static/**/*.html"
  ],
  theme: {
    // https://tailwindcss.com/docs/customizing-colors#using-the-default-colors
    colors: {
      "spotify": {
        DEFAULT: "#1ED760",
      },
      transparent: 'transparent',
      current: 'currentColor',
      black: colors.black,
      white: colors.white,
      gray: colors.gray,
      green: colors.green,
    },
    extend: {},
  },
  plugins: [],
}
