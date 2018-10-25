//用于判断侧栏菜单状态，1为打开，0为关闭
var side_nav_status = 0;
//点击侧栏外面时关闭
window.onclick = function (e) {
    if (!e.target.classList.contains("menuIcon")
        && !e.target.parentNode.classList.contains("menuIcon")
        && !e.target.parentNode.classList.contains("sidenav")
        && !e.target.classList.contains("sidenav")
    ) {
        if (side_nav_status == 1)
            closeNav();
    }
};

//menu动画,以及控制侧滑菜单
function menuAnimation(x) {
    //菜单动画，默认关闭
    // x.classList.toggle("change");

    if (side_nav_status == 0) {
        openNav();

    }
    else {
        closeNav();
    }
}

//打开侧栏菜单
function openNav() {
    document.getElementById("mySidenav").style.width = "12em";
    side_nav_status = 1;
}

//关闭侧栏菜单
function closeNav() {
    document.getElementById("mySidenav").style.width = "0";
    side_nav_status = 0;

}
function go(action) {
    location.href=action;
    showLoading();

}

function showLoading() {
    var loading=document.getElementById("loading");
    loading.style.display="block";
}

function hideLoading() {
    var loading = document.getElementById("loading");
    loading.style.display = "none";
}
function setCookie(c_name, value, expiredays) {
    var exdate = new Date()

    if (expiredays == null) {
        var exdate2 = new Date();
        exdate2.setFullYear(exdate.getFullYear() + 99);//永久生效，除非用户删除cookie
    }
    exdate.setDate(exdate.getDate() + expiredays)
    document.cookie = c_name + "=" + escape(value) +
        ((expiredays == null) ? ";expires=" + exdate2.toGMTString() : ";expires=" + exdate.toGMTString())
}

function getCookie(c_name) {
    if (document.cookie.length > 0) {
        c_start = document.cookie.indexOf(c_name + "=")
        if (c_start != -1) {
            c_start = c_start + c_name.length + 1
            c_end = document.cookie.indexOf(";", c_start)
            if (c_end == -1) c_end = document.cookie.length
            return unescape(document.cookie.substring(c_start, c_end))
        }
    }
    return ""
}