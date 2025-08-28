import type { Config } from 'tailwindcss'

export default <Partial<Config>>{
  theme: {
    extend: {
      colors: {
        primary: '#ff0000',
        red: {
          50: '#fef2f2',
          100: '#fde7e6',
          200: '#fbd0d1',
          300: '#f7aaab',
          400: '#f27a7f',
          500: '#ff0000', // Override the 500 shade as primary
          600: '#d5293e',
          700: '#c01f36',
          800: '#961b31',
          900: '#811a30',
          950: '#470a15',
        },
      },
    },
  },
}
