module.exports = {
    devServer: {
        port: 8080,
        proxy: {
        '/api': {
          target: 'http://localhost:8081', //服务端地址
          secure: false,
          changeOrigin: true, // 允许跨域
          pathRewrite: {
            '^/api': ''
          }
        }
      },
    }
  }