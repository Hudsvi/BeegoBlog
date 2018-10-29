//用于生成UID，
function generateUUID() {
    var d = new Date().getTime();
    var uuid = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
        var r = (d + Math.random() * 16) % 16 | 0;
        d = Math.floor(d / 16);
        return (c ==='x' ? r : (r & 0x3 | 0x8)).toString(16);
    });
    return uuid;
}

//
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

