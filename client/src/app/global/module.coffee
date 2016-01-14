this.global = global = {}

app = angular.module 'global', ['ngRoute','ngMaterial']

app.config ["$mdThemingProvider",($mdThemingProvider)->
  $mdThemingProvider.theme('default')
  .primaryPalette('blue')
  .accentPalette('orange')
]

app.controller "globalCtrl", ["$rootScope", ($scope)->
  $scope.statusCode = 0
  $scope.statusText = "Loaded"
]

app.controller "basicCtrl", ['$scope', ($scope)->
  $scope.title = "Home Automation"
]

app.service 'api', HAApi

class global.BaseController
  @$inject = [
      "$scope"
      "api"
      "$routeParams"
      "$injector"
    ]

  constructor: (@scope, @api, @params, @injector) ->
    @init()

  model: (m)->
    if m?
      @scope.model = m
      return @
    return m
