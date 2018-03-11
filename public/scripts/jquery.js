$( document ).ready(function() {

        //Perhaps later add a function to preload all images. See below...
        // https://stackoverflow.com/questions/476679/preloading-images-with-jquery

        if (top.location.pathname === '/dashboard') {
            $(".dashboard-icon").attr("src","public/assets/dashboard_active.svg");
            $("#dashboard-header-text").css("color","#7FC347");

            $("#dashboard-header-text").hover(function () {
              $(".dashboard-icon").attr("src","public/assets/dashboard_hover.svg");
            }, function () {
              $(".dashboard-icon").attr("src","public/assets/dashboard_active.svg");
            });

            $("#diary-header-text").hover(function () {
              $(".diary-icon").attr("src","public/assets/diary_hover.svg");
            }, function () {
              $(".diary-icon").attr("src","public/assets/diary_inactive.svg");
            });

            $("#analytics-header-text").hover(function () {
              $(".analytics-icon").attr("src","public/assets/analytics_hover.svg");
            }, function () {
              $(".analytics-icon").attr("src","public/assets/analytics_inactive.svg");
            });

            $("#logout-header-text").hover(function () {
              $(".logout-icon").attr("src","public/assets/logout_hover.svg");
            }, function () {
              $(".logout-icon").attr("src","public/assets/logout_inactive.svg");
            });
        }

        if (top.location.pathname === '/diaryPage') {
            $(".diary-icon").attr("src","public/assets/diary_active.svg");
            $("#diary-header-text").css("color","#7FC347");

            $("#dashboard-header-text").hover(function () {
              $(".dashboard-icon").attr("src","public/assets/dashboard_hover.svg");
            }, function () {
              $(".dashboard-icon").attr("src","public/assets/dashboard_inactive.svg");
            });

            $("#diary-header-text").hover(function () {
              $(".diary-icon").attr("src","public/assets/diary_hover.svg");
            }, function () {
              $(".diary-icon").attr("src","public/assets/diary_active.svg");
            });

            $("#analytics-header-text").hover(function () {
              $(".analytics-icon").attr("src","public/assets/analytics_hover.svg");
            }, function () {
              $(".analytics-icon").attr("src","public/assets/analytics_inactive.svg");
            });

            $("#logout-header-text").hover(function () {
              $(".logout-icon").attr("src","public/assets/logout_hover.svg");
            }, function () {
              $(".logout-icon").attr("src","public/assets/logout_inactive.svg");
            });
        }

        if (top.location.pathname === '/analytics') {
            $(".analytics-icon").attr("src","public/assets/analytics_active.svg");
            $("#analytics-header-text").css("color","#7FC347");

            $("#dashboard-header-text").hover(function () {
              $(".dashboard-icon").attr("src","public/assets/dashboard_hover.svg");
            }, function () {
              $(".dashboard-icon").attr("src","public/assets/dashboard_inactive.svg");
            });

            $("#diary-header-text").hover(function () {
              $(".diary-icon").attr("src","public/assets/diary_hover.svg");
            }, function () {
              $(".diary-icon").attr("src","public/assets/diary_inactive.svg");
            });

            $("#analytics-header-text").hover(function () {
              $(".analytics-icon").attr("src","public/assets/analytics_hover.svg");
            }, function () {
              $(".analytics-icon").attr("src","public/assets/analytics_active.svg");
            });

            $("#logout-header-text").hover(function () {
              $(".logout-icon").attr("src","public/assets/logout_hover.svg");
            }, function () {
              $(".logout-icon").attr("src","public/assets/logout_inactive.svg");
            });
        }

        if ((top.location.pathname === '/') ||  (top.location.pathname === '/signup')) {
            $("#dashboard-header-text").hover(function () {
              $(".dashboard-icon").attr("src","public/assets/dashboard_hover.svg");
            }, function () {
              $(".dashboard-icon").attr("src","public/assets/dashboard_inactive.svg");
            });

            $("#diary-header-text").hover(function () {
              $(".diary-icon").attr("src","public/assets/diary_hover.svg");
            }, function () {
              $(".diary-icon").attr("src","public/assets/diary_inactive.svg");
            });

            $("#analytics-header-text").hover(function () {
              $(".analytics-icon").attr("src","public/assets/analytics_hover.svg");
            }, function () {
              $(".analytics-icon").attr("src","public/assets/analytics_inactive.svg");
            });

            $("#logout-header-text").hover(function () {
              $(".logout-icon").attr("src","public/assets/logout_hover.svg");
            }, function () {
              $(".logout-icon").attr("src","public/assets/logout_inactive.svg");
            });
        }

});
