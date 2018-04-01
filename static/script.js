var showCodeSnippet = function() {
    $(this).addClass("active").siblings().removeClass("active");
    var targetID = $(this).attr("target-id");
    $("#" + targetID).show().siblings().hide();
};

$('.lan').on("click", showCodeSnippet);
