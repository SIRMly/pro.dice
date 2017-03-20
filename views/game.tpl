<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>吹牛</title>
    <meta name="renderer" content="webkit">
    <meta content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" name="viewport">
    <meta content="no-cache,must-revalidate" http-equiv="Cache-Control">
    <meta content="no-cache" http-equiv="pragma">
    <meta content="0" http-equiv="expires">
    <meta content="telephone=no, address=no" name="format-detection">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black-translucent">
    <link rel="stylesheet" type="text/css" href="/static/css/game.css">
</head>
<body>
<div class="container">
    <audio src="http://fun-x.b0.upaiyun.com/chuiniu/audio/14.mp3" preload="auto" id="audio-bg" style="display: none"></audio>
    <div id="cude-wrapper">
        <img src="http://fun-x.b0.upaiyun.com/chuiniu/img/play-state.png" alt="" width="100%" id="state">
        <div class="cube">
            <div class="side side1">
            </div>
            <div class="side side2">
            </div>
            <div class="side side3">
            </div>
            <div class="side side4">
            </div>
            <div class="side side5">
            </div>
            <div class="side side6">
            </div>
            <div class="crash crash1"></div>
            <div class="crash crash2"></div>
        </div>
        <div class="cube cube2">
            <div class="side side1">
            </div>
            <div class="side side2">
            </div>
            <div class="side side3">
            </div>
            <div class="side side4">
            </div>
            <div class="side side5">
            </div>
            <div class="side side6">
            </div>
            <div class="crash crash1"></div>
            <div class="crash crash2"></div>
        </div>
        <div class="cube cube3">
            <div class="side side1">
            </div>
            <div class="side side2">
            </div>
            <div class="side side3">
            </div>
            <div class="side side4">
            </div>
            <div class="side side5">
            </div>
            <div class="side side6">
            </div>
            <div class="crash crash1"></div>
            <div class="crash crash2"></div>
        </div>
        <br>
        <div class="cube cube4">
            <div class="side side1">
            </div>
            <div class="side side2">
            </div>
            <div class="side side3">
            </div>
            <div class="side side4">
            </div>
            <div class="side side5">
            </div>
            <div class="side side6">
            </div>
            <div class="crash crash1"></div>
            <div class="crash crash2"></div>
        </div>
        <div class="cube cube5">
            <div class="side side1">
            </div>
            <div class="side side2">
            </div>
            <div class="side side3">
            </div>
            <div class="side side4">
            </div>
            <div class="side side5">
            </div>
            <div class="side side6">
            </div>
            <div class="crash crash1"></div>
            <div class="crash crash2"></div>
        </div>
    </div>
    <div id="yao-wrapper">
        <img src="http://fun-x.b0.upaiyun.com/chuiniu/img/yao-pic.png" alt="" id="yao-pic" width="22%">
        <br>
        <img src="http://fun-x.b0.upaiyun.com/chuiniu/img/yao-word.png" alt="" id="yao-word" width="42%">
    </div>
    <a href="javascript:;" id="reward"></a>
</div>
<div id="crash">
    <img src="http://fun-x.b0.upaiyun.com/chuiniu/img/0.jpg" alt="" id="crash-img">
</div>
<script type="text/javascript" src="http://pingjs.qq.com/h5/stats.js" name="MTAH5" sid="500141455" ></script>
</body>
<script type="text/javascript" src="http://fun-x.b0.upaiyun.com/chuiniu/js/shake.js"></script>
<script type="text/javascript" src="http://res.wx.qq.com/open/js/jweixin-1.0.0.js"></script>
<script type="text/javascript">
window.onload = function () {
    // 防止下滑
    document.ontouchmove = function (e) {
        e.preventDefault();
    };
    // 出现
    setTimeout(function () {
        for(var i = 0; i < 5; i++) {
            document.getElementsByClassName("cube")[i].className = "cube cubeShow";
        }
    }, 800);

    var randomNum;
    var cubeAngel = [
        [0, 0],
        [0, -90],
        [0, -180],
        [0, 90],
        [-90, 0],
        [90, 0]
    ];
    var myShakeEvent = new Shake({
        threshold: 10
    });
    myShakeEvent.start();
    window.addEventListener('shake', shakeEventDidOccur, false);
    // 摇动
    function shakeEventDidOccur() {
        socket.send("2");
        var audio       = document.getElementById("audio-bg");
        var audioTime   = true;
        audio.play();
        for (var i = 0; i < 5; i++) {
            document.getElementsByClassName("cube")[i].className = "cube cubeShow cubeMove" + i;
        }
    }
    // 结束
    function  cubeStop(){
        setTimeout(function () {
            for (var i = 0; i < 5; i++) {
                randomNum = Math.floor(Math.random() * 6);
                document.getElementsByClassName("cube")[i].className = "cube cubeShow";
                document.getElementsByClassName("cube")[i].style.transform = "rotateX(" + cubeAngel[randomNum][0] + "deg) rotateY(" + cubeAngel[randomNum][1] + "deg)";
            }
        }, 3000);
        window.removeEventListener('shake', shakeEventDidOccur, false);
        setTimeout(function () {
            window.addEventListener('shake', shakeEventDidOccur, false);
        }, 3200);
    }

    document.getElementById("reward").addEventListener("click", function () {
        var rewardN = Math.floor(Math.random()*3);
        document.getElementById("crash").style.display = "block";
        document.getElementById("crash-img").setAttribute("src", "http://fun-x.b0.upaiyun.com/chuiniu/img/" + rewardN + ".jpg");
        window.removeEventListener('shake', shakeEventDidOccur, false);
    }, false);

    document.getElementById("crash").addEventListener("click", function () {
        this.style.display = "none";
        window.addEventListener('shake', shakeEventDidOccur, false);
    }, false);

    var socketUrl   = "ws://" + window.location.host + ":8088/join/{{ .Name }}";
    var socket      = new WebSocket(socketUrl);
    socket.onopen   = function () {
        // 页面加载完成发送 准备
        socket.send("0");
    };

    socket.onmessage = function (event) {
        switch (event.data) {
            case "3":
                cubeStop();
                break;
        }
    }

    var options = {
        title   : '喝酒没带骰子？有手机就搞定了！赶紧收藏吧。',
        desc    : '喝酒没带骰子？有手机就搞定了！赶紧收藏吧。',
        link    : 'http://game.fun-x.cn',
        imgUrl  : 'http://game.fun-x.cn/static/img/share.jpg'
    };
    wx.config({
        debug       : false,
        appId       : '{{ .Appid }}', // 必填，公众号的唯一标识
        timestamp   : {{ .Timestamp }},
        nonceStr    : '{{ .NonceStr }}', // 必填，生成签名的随机串
        signature   : '{{ .Signature }}',
        jsApiList   : ['onMenuShareAppMessage', 'onMenuShareTimeline']
    });
    wx.ready(function(){
        wx.onMenuShareTimeline({
            title   : options.desc,
            link    : options.link,
            imgUrl  : options.imgUrl
        });

        wx.onMenuShareAppMessage({
            title   : options.title,
            desc    : options.desc,
            link    : options.link,
            imgUrl  : options.imgUrl
        });
    });
};
</script>
</html>
