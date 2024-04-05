/** @type {import('tailwindcss').Config} */
module.exports = {
  content: {
    relative: true,
    files: [
      "./*.html",
      "./blogs/*.html",
      "/blogs/*/index.html",
    ],
  },
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
}

