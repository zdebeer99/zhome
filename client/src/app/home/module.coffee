app = angular.module('app',['global'])

viewpath = (view)->
  "/html/home/#{view}.html"

app.config ["$routeProvider", ($route)->
  $route
    .when "/",
      templateUrl: viewpath("home"),
      controller: "homeCtrl"
    .when "/charts/:id",
      templateUrl: viewpath("charts"),
      controller: "chartCtrl"
    .otherwise redirectTo: "/"
]
