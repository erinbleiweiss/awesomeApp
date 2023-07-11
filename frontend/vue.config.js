const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: 'http://backend:8888/',
    // proxy: 'http://localhost:8888/',
  }
})
