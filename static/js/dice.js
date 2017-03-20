/**
 * Created by SIRMly on 2016/6/21.
 */
    // 防止下滑
    document.ontouchmove = function (e) {
        e.preventDefault();
    };
    function preload(images, numWrapper) {
        //预加载图片总数
        var n = images.length,
            loadedimg = 0,
            percent = 0;
        for (var i = 0; i < images.length; i++) {
            var imgNew = new Image();
            imgNew.src = images[i];
            imgNew.onload = function(){//每张图片加载成功后执行
                loadedimg++;
                console.log(loadedimg);
                percent = Math.round(loadedimg/n*100 );
                numWrapper.innerText = percent + "%";
                if(percent >= 100){
                    showGamePage();
                }
            };
        };
    };
    var numberWrapper = document.getElementById("loading");
    preload([
        "../static/img/0.jpg",
        "../static/img/dice1.png",
        "../static/img/dice2.png",
        "../static/img/dice3.png",
        "../static/img/dice4.png",
        "../static/img/dice5.png",
        "../static/img/dice6.png",
        "../static/img/play-bg.jpg",
        "../static/img/play-state.png",
        "../static/img/playcon-bg.png",
        "../static/img/reward.png",
        "../static/img/yao-pic.png",
        "../static/img/yao-word.png"
    ], numberWrapper);

    function showGamePage(){
        setTimeout(function (){
            document.getElementById("page-index").style.display = "none";
            document.getElementById("page-game").style.display = "block";
            page2();
        },500)
    }
    function page2(){
        // 出现
        setTimeout(function () {
            for(var i = 0; i < 5; i++) {
                document.getElementsByClassName("cube")[i].className = "cube cubeShow";
            }
        }, 0);

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
            var audio       = document.getElementById("audio-bg");
            var audioTime   = true;
            audio.play();
            for (var i = 0; i < 5; i++) {
                document.getElementsByClassName("cube")[i].className = "cube cubeShow cubeMove" + i;
            }
            setTimeout(function () {
                cubeStop();
            },3000)
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
            document.getElementById("crash").style.display = "block";
            window.removeEventListener('shake', shakeEventDidOccur, false);
        }, false);

        document.getElementById("crash").addEventListener("click", function () {
            this.style.display = "none";
            window.addEventListener('shake', shakeEventDidOccur, false);
        }, false);
    }
