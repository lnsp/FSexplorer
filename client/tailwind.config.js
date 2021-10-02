module.exports = {
  theme: {
    fontFamily: {
      sans: [
        'Inter',
        '-apple-system',
        'BlinkMacSystemFont',
        'Segoe UI',
        'Roboto',
        'Oxygen',
        'Ubuntu',
        'Cantarell',
        'Open Sans',
        'Helvetica Neue',
        'sans-serif',
      ],
    },
  },
  variants: {
    extend: {
      rotate: ['group-hover'],
    }
  },
  plugins: [
    require('@tailwindcss/forms')
  ]
}
