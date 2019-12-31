// pages/douban/splash.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    subjects: [],
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    let app = getApp()
    app.request(`${app.globalData.api}/movie/bang/coming_soon?start=0&count=3`).then(data => {
      this.setData({
        subjects: data.subjects
      })
    }).catch(error => {
      console.log(error)
    })

    wx.setStorage({
      data: true,
      key: 'has_shown_splash',
    });
    setTimeout(() => {
      wx.redirectTo({
        url: 'index',
      })
    }, 10000);
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

  }
})