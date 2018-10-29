//页码跳转
function jumpPage() {

}

function appendUID(obj) {
    var id = obj.attributes["articleID"].value;//文章id
    if (getCookie("uuid") == null || getCookie("uuid") === "") {
        setCookie("uuid", generateUUID().replace(/-/g, "_")) //保存唯一标识
    } else {
        var uuid = getCookie("uuid")
    }
    var href = "/blog/article/" + id + "?uuid=" + uuid + ":" + id;
    go(href);
}
