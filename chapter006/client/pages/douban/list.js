// pages/douban/list.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    type: 'in_theaters',
    page: 1,
    size: 20,
    total: 1,
    movies: [],
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    let type = options.type;
    if(type){
      this.setData({type: type});
    }
    this.retrieve()
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  },
  retrieve(){
    let app = getApp()
    let start = (this.data.page - 1) * this.data.size
    wx.showLoading({
      title: '加载中',
    })
    return app.request(`${app.globalData.api}/movie/bang/${this.data.type}?start=${start}&count=${this.data.size}`)
      .then(res => {
        if(res.subjects.length){
          let movies = this.data.movies.concat(res.subjects)
          let total = Math.floor(res.total / this.data.size)
          this.setData({movies: movies, total: total, page: this.data.page})
          wx.setNavigationBarTitle({title: res.title})
          console.log(movies)
        }
      }).catch(err => {
        console.log(err)
      }).finally(() => {
        wx.hideLoading()
      })
  },
  loadMorePage(){
    if(this.data.page > this.data.total) return
    this.data.page++
    this.retrieve()
  }
})