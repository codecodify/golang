//app.js
// apikey=0b2bdeda43b5688921839c8ecb20399b
  let app = require('./sim.js/index')
App(Object.assign(app, {
  onLaunch: function () {
  
  },
  globalData: {
    api: "http://127.0.0.1:8888", // api地址
  }

}))