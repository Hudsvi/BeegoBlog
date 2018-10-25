window.onload = function (ev) {
    // 页面加载完时保存所有checkbox数据。
    // 作用是用于筛选时还原复选框列表。
    var cbox = document.getElementById("tags-div");

    //你可能很疑惑，为什么不用cookie保存而使用localStorage/sessionStorage,
    // 你不得不信的是，cookie的个数和总大小限制太局限了！
    //因此cookie只适合储存小文件。
    // @Depracated setCookie("tagList", cbox.innerHTML, 1);

    //storage的用法http://www.w3school.com.cn/html/html5_webstorage.asp
    sessionStorage.tagList = cbox.innerHTML;
};

function saveAsDraft() {
    var form = document.getElementById("formWrite");
    var blogTitle = $('#blogTitle').val();//标题
    var blogContent = xhEditor2.getSource();//正文
    var workId = $('#workId').val();
    var tags = [];//标签
    var cbox = document.getElementsByName("cbox-tag");//

    /*获取选中的标签*/
    for (var i = 0; i < cbox.length; i++) {
        if (cbox[i].checked) {
            tags.push(cbox[i].value);
        }
    }
    if (workId != "-1" && tags.length > 0) {
        //追加参数action,用于区分当前的提交是保存草稿还是正式发表。=save表示保存草稿，=publish表示发表。
        form.action = form.action + "?blogTitle=" + blogContent + "&blogContent=" +
            blogContent + "&workId=" + workId + "&action=save";
        for (i in tags) {
            form.action = form.action + "&tags=" + tags[i];
        }
        form.method = "post";
        form.submit();
    } else {
        alert("请将标签和文集补充完整再提交！");
    }
}

function publish() {

}

function filterTag(obj) {
    var parentID = obj.options[obj.selectedIndex].value;
    var tagsDIV = document.getElementById("tags-div");
    // 因大小限制，cookie方式弃用:var tagList = getCookie("tagList").toString();
    var  tagList=sessionStorage.tagList;
    var tags = tagList.split("<p>");
    var new_tags = "";

    /*如果不筛选，则默认加载显示标签*/
    if (parentID == "0") {
        tagsDIV.innerHTML = tagList;
    } else {
        //split "<p>"后，第一项内容为空，所以不用考虑。
        //过滤掉parentID以外的checkbox
        for (i = 1; i < tags.length; i++) {
            if (tags[i].indexOf("parent-id=\"" + parentID + "\"") > -1) {
                new_tags += "<p>" + tags[i];//split时“<p>”被删除，现在还原。
            }
        }
        tagsDIV.innerHTML = new_tags.toString();
    }

}

//清空xhEditor内容
function clearContent() {
    if (xhEditor2.getSource() != "" && confirm("清空后无法恢复，确定要清空？")) {
        $('#myContent').val("");
    }
}