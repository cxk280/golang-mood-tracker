$( document ).ready(function() {
        console.log( "ready!" );

        $("#dashboard-header-text").hover(function () {
          console.log('hovering');
          $(".dashboard-icon").attr("src","public/assets/dashboard_active.svg");
        }, function () {
          console.log('no longer hovering');
          $(".dashboard-icon").attr("src","public/assets/dashboard_inactive.svg");
        });

        $("#diary-header-text").hover(function () {
          console.log('hovering');
          $(".diary-icon").attr("src","public/assets/diary_active.svg");
        }, function () {
          console.log('no longer hovering');
          $(".diary-icon").attr("src","public/assets/diary_inactive.svg");
        });

        $("#analytics-header-text").hover(function () {
          console.log('hovering');
          $(".analytics-icon").attr("src","public/assets/analytics_active.svg");
        }, function () {
          console.log('no longer hovering');
          $(".analytics-icon").attr("src","public/assets/analytics_inactive.svg");
        });

        // $("#logout-header-text").hover(function () {
        //   console.log('hovering');
        //   $(".logout-icon").attr("src","public/assets/logout_active.svg");
        // }, function () {
        //   console.log('no longer hovering');
        //   $(".logout-icon").attr("src","public/assets/logout_inactive.svg");
        // });
});
