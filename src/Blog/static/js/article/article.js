function submitForm() {
    var form = document.getElementById("publishComment");//表单
    var id = document.getElementById("uuid");//uuid标签
    var time = document.getElementById("datetime");//时间标签
    var uuid = getCookie("uuid");//uuid
    id.setAttribute("value", uuid);
    var datetime = new Date().Format("yyyy-MM-dd hh:mm:ss");//日期
    time.value = datetime;
    if (uuid !== "") {
        if (xhEditor.getSource() === "") {
            alert("评论内容不能为空！");
        } else {
            form.submit();
            showLoading();
        }
    }
}

//日期格式化
Date.prototype.Format = function (fmt) {
    var o = {
        "M+": this.getMonth() + 1,
        "d+": this.getDate(),
        "h+": this.getHours(),
        "m+": this.getMinutes(),
        "s+": this.getSeconds(),
        "q+": Math.floor((this.getMonth() + 3) / 3),
        "S": this.getMilliseconds()
    };
    if (/(y+)/.test(fmt))
        fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
        if (new RegExp("(" + k + ")").test(fmt))
            fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    return fmt;
}