var index = 0;
function changeSlide(n) {
    var i;
    var x = document.getElementsByClassName("tm-special-item-slider");
    console.log(x)
    index += n;
    if (index > x.length) {index = 1}
    if (index < 1) {index = x.length}
    for (i = 0; i < x.length; i++) {
        x[i].style.display = "none";
    }
    x[index-1].style.display = "block";
}

changeSlide(0);