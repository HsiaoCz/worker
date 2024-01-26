const baseURL = "https://cea6-58-19-1-21.ngrok-free.app";

const wxWithDefaultHeadersRequest = (options) => {
  return new Promise((resolve, reject) => {
    wx.request({
      url: baseURL + options.url,
      method: options.method,
      data: options.data,
      header: {
        "ngrok-skip-browser-warning": "69420",
      },
      success:resolve,
      fail:reject,
    });
  });
};

export default wxWithDefaultHeadersRequest