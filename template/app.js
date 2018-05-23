$(function() {
  $('.request').each(function(index, el) {
    el = $(el);
    el.find('.request-content button').click(function() {
      var url = el.attr('data-request');
      var method = el.attr('data-method');
      var params = el.find('.request-content input.param');
      var data = {};
      for (var i = 0; i < params.length; i++) {
        var param = $(params[i]);
        var key = param.attr('data-key');
        var value = param.val()
        if(param.hasClass('is-required') && !value) {
          param.addClass('error');
          console.log('缺少必要参数', key);
          return;
        }
        data[key] = value || param.attr('data-default');
      }
      $.ajax({
        url: url,
        type: method,
        data: data,
        success: function(data) {
          console.log(data);
        },
        complete: function(xhr, status) {
          console.log(xhr, status);
          console.log(xhr.getAllResponseHeaders());
        }
      })
      console.log('data', data);
    })

    el.find('.request-content input.param').on('click', '.error', function() {
      this.removeClass('error');
    })
  })
})
