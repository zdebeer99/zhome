$ = require('gulp-load-plugins')()
gulp = require 'gulp'
coffee = require 'gulp-coffee'
sourcemaps = require("gulp-sourcemaps")
sequence = require 'run-sequence'
del = require 'del'
nodemon = require 'gulp-nodemon'
_ = require 'lodash'

DEBUG = true

paths =
  src : './src/'
  output : '../static/'
  build : './build/'
  temp : './build/tmp/'

files =
  coreJS: [
    'bower_components/jquery/dist/jquery.min.js'
    'bower_components/lodash/lodash.min.js'
    #'bower_components/fastclick/lib/fastclick.js'
    #'bower_components/viewport-units-buggyfill/viewport-units-buggyfill.js'
    #'bower_components/tether/tether.js'
    #'bower_components/hammerjs/hammer.js'
    'bower_components/angular/angular.min.js'
    'bower_components/angular-route/angular-route.min.js'
    'bower_components/angular-aria/angular-aria.min.js'
    'bower_components/angular-animate/angular-animate.min.js'
    'bower_components/angular-messages/angular-messages.min.js'
    'bower_components/angular-material/angular-material.min.js'
    'bower_components/moment/min/moment.min.js'

    #'bower_components/angular-ui-router/release/angular-ui-router.js'
  ]
  angularmaterialCSS: [
    'bower_components/angular-material/angular-material.min.css'
  ]
  vendor: [
    'bower_components/morris.js/morris.min.js'
    'bower_components/raphael/raphael-min.js'
    'bower_components/angular-morris-chart/src/angular-morris-chart.min.js'
  ]

gulp.task 'default', ['buildwatch']
gulp.task 'build', ['core', 'app', 'views']
gulp.task 'core', ['coreJS','angularmaterial:css', 'vendor']

gulp.task 'buildwatch', (cb) ->
  sequence 'build','watch', cb

#update client side scripts
#gulp.task 'app', ['app:global','app:home','app:action','app:css']

gulp.task 'app', ['app:global','app:home','app:css']

gulp.task 'views', ->
  gulp.src ["!#{paths.src}app/**/_*.jade","#{paths.src}app/**/*.jade"]
  .pipe $.jade(pretty: true)
  .pipe gulp.dest paths.build + 'html/'
  .pipe gulp.dest dest('html')

# compile script for single page application 'home'
gulp.task 'app:global', ->
  build_coffee "global",["app/global"]

gulp.task 'app:home', ->
  build_coffee "home",[
    "app/home/module.coffee"
    "app/home"
  ]

gulp.task 'app:css', ->
  gulp.src paths.base + 'css/*.css'
  .pipe $.concat 'app.css'
  .pipe gulp.dest paths.build
  .pipe gulp.dest dest "css"

gulp.task 'coreJS', ->
  gulp.src files.coreJS
  .pipe $.concat 'core.js'
  .pipe gulp.dest paths.build
  .pipe gulp.dest dest "js"

gulp.task 'vendor', ->
  gulp.src files.vendor
  .pipe gulp.dest dest "js"

###
  angular - material
###

gulp.task 'angularmaterial:css', ->
  gulp.src files.angularmaterialCSS
  .pipe gulp.dest paths.build + "css"
  .pipe gulp.dest dest("css")

###
  Gulp General Tasks
###

gulp.task 'watch', ->
  gulp.watch source('app/**/*.coffee'), ['app']
  gulp.watch source('app/**/*.jade'), ['views']

gulp.task 'clean', (cb) ->
  del [
    paths.build
    dest "**/*"
    "./src/build/"
    ], {force: true} , cb

###
# Gulp helper functions
###

dest = (name) ->
  paths.output + name

source = (name) ->
  paths.src + name

source_coffee = (names, path) ->
  (paths.src + path + name + '.coffee' for name in names)

prefix_path = (names, path) ->
  (path + name for name in names)

#concat and build files from
build_coffee = (name, modulePath) ->
  f1 = []
  for mp1 in modulePath
    if endsWith(mp1,".coffee")
      f1.push paths.src + mp1
    else
      f1.push paths.src + mp1 + "/**/*.coffee"
  gulp.src f1
  .pipe $.concat(name + ".coffee")
  .pipe gulp.dest paths.temp
  .pipe coffee()
  .pipe gulp.dest paths.build + "js"
  .pipe gulp.dest paths.output + "js"

endsWith = (str, suffix) ->
  return str.indexOf(suffix, str.length - suffix.length) != -1

String::endsWith = endsWith
