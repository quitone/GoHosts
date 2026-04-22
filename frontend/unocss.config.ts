import { defineConfig } from 'unocss'
import presetMini from '@unocss/preset-mini'
import presetIcons from '@unocss/preset-icons'
import transformerVariantGroup from '@unocss/transformer-variant-group'
import transformerDirectives from '@unocss/transformer-directives'

export default defineConfig({
  rules: [],
  shortcuts: {
    'flex-center': 'flex items-center justify-center',
    'flex-between': 'flex items-center justify-between',
    'flex-start': 'flex items-center justify-start',
    'flex-end': 'flex items-center justify-end',
    'flex-col': 'flex flex-col',
    'flex-col-center': 'flex flex-col items-center justify-center',
    'grid-center': 'grid place-items-center',
    'text-ellipsis': 'overflow-hidden text-ellipsis whitespace-nowrap',
    'wh-full': 'w-full h-full',
    'absolute-center': 'absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2',
    'transition-base': 'transition-all duration-200 ease-in-out',
    'card-base': 'bg-white dark:bg-gray-800 rounded-lg shadow-md p-4',
    'btn-base': 'px-4 py-2 rounded-md font-medium transition-base hover:opacity-90',
    'btn-primary': 'btn-base bg-blue-500 text-white hover:bg-blue-600',
    'btn-secondary': 'btn-base bg-gray-200 text-gray-700 hover:bg-gray-300 dark:bg-gray-700 dark:text-gray-200',
    'btn-success': 'btn-base bg-green-500 text-white hover:bg-green-600',
    'btn-danger': 'btn-base bg-red-500 text-white hover:bg-red-600',
    'input-base': 'w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700',
  },
  theme: {
    colors: {
      page: '#0f172a',
      // 主色调 - 与 naive-ui 默认主题保持一致
      primary: {
        DEFAULT: '#2080f0',
        hover: '#4098fc',
        active: '#1060c9',
        light: '#e8f3ff',
        dark: '#0e59c4',
      },
      // 次要色调
      info: {
        DEFAULT: '#2080f0',
        hover: '#4098fc',
        active: '#1060c9',
        light: '#e8f3ff',
      },
      // 成功色
      success: {
        DEFAULT: '#18a058',
        hover: '#36ad6a',
        active: '#0c7a43',
        light: '#e8f8ee',
      },
      // 警告色
      warning: {
        DEFAULT: '#f0a020',
        hover: '#fcb040',
        active: '#c97c10',
        light: '#fff7e8',
      },
      // 错误色
      error: {
        DEFAULT: '#d03050',
        hover: '#de576d',
        active: '#ab1f3f',
        light: '#ffeef3',
      },
      // 中性色
      neutral: {
        50: '#fafafa',
        100: '#f5f5f5',
        200: '#eeeeee',
        300: '#e0e0e0',
        400: '#bdbdbd',
        500: '#9e9e9e',
        600: '#757575',
        700: '#616161',
        800: '#424242',
        900: '#212121',
      },
      // 背景色
      background: {
        DEFAULT: '#ffffff',
        secondary: '#f5f5f5',
        hover: '#f0f0f0',
      },
      // 文字颜色
      text: {
        primary: '#333333',
        secondary: '#666666',
        tertiary: '#999999',
        disabled: '#cccccc',
      },
      // 边框颜色
      border: {
        DEFAULT: '#e5e5e5',
        light: '#f0f0f0',
        dark: '#d0d0d0',
      },
    },
    fontSize: {
      xs: ['0.75rem', { lineHeight: '1rem' }],
      sm: ['0.875rem', { lineHeight: '1.25rem' }],
      base: ['1rem', { lineHeight: '1.5rem' }],
      lg: ['1.125rem', { lineHeight: '1.75rem' }],
      xl: ['1.25rem', { lineHeight: '1.75rem' }],
      '2xl': ['1.5rem', { lineHeight: '2rem' }],
      '3xl': ['1.875rem', { lineHeight: '2.25rem' }],
      '4xl': ['2.25rem', { lineHeight: '2.5rem' }],
      '5xl': ['3rem', { lineHeight: '1' }],
    },
    spacing: {
      '0': '0',
      '1': '0.25rem',
      '2': '0.5rem',
      '3': '0.75rem',
      '4': '1rem',
      '5': '1.25rem',
      '6': '1.5rem',
      '8': '2rem',
      '10': '2.5rem',
      '12': '3rem',
      '16': '4rem',
      '20': '5rem',
      '24': '6rem',
      '32': '8rem',
    },
    borderRadius: {
      none: '0',
      sm: '0.125rem',
      DEFAULT: '0.25rem',
      md: '0.375rem',
      lg: '0.5rem',
      xl: '0.75rem',
      '2xl': '1rem',
      '3xl': '1.5rem',
      full: '9999px',
    },
    boxShadow: {
      sm: '0 1px 2px 0 rgba(0, 0, 0, 0.05)',
      DEFAULT: '0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)',
      md: '0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)',
      lg: '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)',
      xl: '0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04)',
      '2xl': '0 25px 50px -12px rgba(0, 0, 0, 0.25)',
      inner: 'inset 0 2px 4px 0 rgba(0, 0, 0, 0.06)',
      none: 'none',
    },
    fontFamily: {
      sans: [
        'Nunito',
        '-apple-system',
        'BlinkMacSystemFont',
        '"Segoe UI"',
        'Roboto',
        '"Helvetica Neue"',
        'Arial',
        '"Noto Sans"',
        'sans-serif',
        '"Apple Color Emoji"',
        '"Segoe UI Emoji"',
        '"Segoe UI Symbol"',
        '"Noto Color Emoji"',
      ],
      serif: ['Georgia', 'Cambria', '"Times New Roman"', 'Times', 'serif'],
      mono: ['Menlo', 'Monaco', 'Consolas', '"Liberation Mono"', '"Courier New"', 'monospace'],
    },
    zIndex: {
      '0': '0',
      '10': '10',
      '20': '20',
      '30': '30',
      '40': '40',
      '50': '50',
      'auto': 'auto',
    },
  },
  presets: [
    presetMini({}),
    presetIcons({
      extraProperties: {
        'display': 'inline-block',
        'vertical-align': 'middle',
      },
    }),
  ],
  transformers: [
    transformerVariantGroup(),
    transformerDirectives(),
  ],
});

