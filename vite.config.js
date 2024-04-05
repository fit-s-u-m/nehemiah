export default {
  build: {
    rollupOptions: {
      input: {
        home: '/index.html',
        blogs: './blogs/index.html',
        trig: './blogs/trig/index.html',
        complex: './blogs/complex/index.html',
      }
    }
  }
}
