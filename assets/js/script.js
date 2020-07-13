// document.querySelector(".card-flip").classList.toggle("flip");

/*
 * Holder.js for demo image
 * Just for demo purpose
 */
Holder.addTheme('gray', {
    bg: '#777',
    fg: 'rgba(255,255,255,.75)',
    font: 'Helvetica',
    fontweight: 'normal'
});

$(function () {
    let flip_button = $("#flip")
    let flip = $(".flip")

    // 翻转
    flip_button.on("click", function (e) {
        let state = flip_button.data('state')
        if (state === "front") {
            flip.css("-webkit-transform", "rotateY(180deg)")
            flip.css("transform", "rotateY(180deg)")
            flip_button.data("state", "back")
        } else {
            flip.css("-webkit-transform", "rotateY(0deg)")
            flip.css("transform", "rotateY(0deg)")
            flip_button.data("state", "front")
        }
    })

    // 上一页
    $("#prev").on("click", function () {
        $.get("/remember/prev", function (data) {
            if (data.card_state === 1) {
                alert(data.card_front)
                return
            }
            let state = flip_button.data('state')
            if (state === "back") {
                flip.css("-webkit-transform", "rotateY(0deg)")
                flip.css("transform", "rotateY(0deg)")
                flip_button.data("state", "front")
            }
            $("#card_back").text(data.card_back)
            $("#card_front").text(data.card_front)
            $("card_id").val(data.card_id)
        })
    })

    // 下一页
    $("#next").on("click", function () {
        $.get("/remember/next", function (data) {
            if (data.card_state === 1) {
                alert(data.card_front)
                return
            }
            let state = flip_button.data('state')
            if (state === "back") {
                flip.css("-webkit-transform", "rotateY(0deg)")
                flip.css("transform", "rotateY(0deg)")
                flip_button.data("state", "front")
            }
            $("#card_back").text(data.card_back)
            $("#card_front").text(data.card_front)
            $("card_id").val(data.card_id)
        })
    })

    // 已掌握
    $("#known").on("click", function () {
        let card_id = $("#card_id").val()
        $.get(`/known?id=${card_id}`, function (data) {
            if (data.card_state === 1) {
                alert(data.card_front)
                window.location.reload()
                return
            }
            let state = flip_button.data('state')
            if (state === "back") {
                flip.css("-webkit-transform", "rotateY(0deg)")
                flip.css("transform", "rotateY(0deg)")
                flip_button.data("state", "front")
            }
            $("#card_back").text(data.card_back)
            $("#card_front").text(data.card_front)
        })
    })
})

function getFormData($form) {
    let un_indexed_array = $form.serializeArray();
    let indexed_array = {};

    $.map(un_indexed_array, function (n, i) {
        indexed_array[n['name']] = n['value'];
    });

    return indexed_array;
}