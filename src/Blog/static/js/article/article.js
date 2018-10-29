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

