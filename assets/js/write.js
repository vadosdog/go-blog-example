$(function () {
    $("#content").bind("input change", () => [
        load()
    ])

    load()
})

function load() {
    $.post("/getHtml", {md: $("#content").val()}, (response) => {
        $("#md_html").html(response.html)
    })
}