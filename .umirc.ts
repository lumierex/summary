import { defineConfig } from 'dumi';

export default defineConfig({
  title: 'notes',
  // favicon: '',
  // logo: '',
  publicPath: process.env.NODE_ENV === 'production' ? '/go-notes/' : '/',
  outputPath: 'docs-dist',
  mode: 'site',
  // more config: https://d.umijs.org/config
});
