
class HomeCtrl extends global.BaseController

  init: ()->
    @interval = @injector.get('$interval')
    @scope.valueChanged = @valueChanged
    @scope.channelTemplate = @channelTemplate
    @scope.refresh = @refresh
    @scope.fromNow = (date)->
      moment(date).fromNow()
    @initTimer()
    @load()


  initTimer:()->
    @stop = @interval(()=>
      @load()
    ,60*1000)
    @scope.$on '$destroy', ()=>
      @interval.cancel(@stop)
      @stop = undefined


  load: ()->
    @api.channelStates().then (data)=>
      model = []
      for ch in data
        if ch.info.enabled
          model.push ch
      @model(model)

  valueChanged: (item)=>
    @api.setValue item.info.name, item.state.value.value

  channelTemplate: (chType)->
    viewpath "ch#{chType}"

  refresh: (item)=>
    @api.getValue item.info.name
    .then (data)->
      item.state = data

app.controller("homeCtrl", HomeCtrl)
