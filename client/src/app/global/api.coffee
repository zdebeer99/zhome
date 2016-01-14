
api = {}

apiPath = (path) ->
  "/api/" + path

class HAApi
  constructor: (@http, @q, @rootScope) ->

  @$inject: ['$http','$q','$rootScope']

  emptyPromise: () ->
    defer = @q.defer()
    defer.resolve()
    return defer.promise

  status: (code, text) ->
    @rootScope.statusCode = code
    @rootScope.statusText = text

  get: (path) ->
    defer = @q.defer()
    @http.get(path)
    .success (data) =>
      if (data.statusCode != 1)
        defer.resolve(data.data)
        @status(0, "")
      else
        @status(data.statusCode, data.statusText)
        defer.reject(data)
    .error (data) =>
      @status 1, data
      defer.reject(data)
    return defer.promise

  post: (path, model) ->
    defer = @q.defer()
    @http.post(path, model)
    .success (data) =>
      if (data.statusCode != 1)
        defer.resolve(data.data)
        @status(0, "")
      else
        @status(data.statusCode, data.statusText)
        defer.reject(data)
    .error (data) =>
      @status 1, data
      defer.reject(data)
    return defer.promise

  _list: (modelName) ->
    @get(apiPath("#{modelName}List"))

  _get: (modelName, id) ->
    @get(apiPath("#{modelName}Get/#{id}"))

  _add: (modelName, model) ->
    @post(apiPath("#{modelName}Add"), model)

  _update: (modelName, id, model) ->
    @post(apiPath("#{modelName}Update/#{id}"), model)

  _remove: (modelName, id) ->
    @get(apiPath("#{modelName}Remove/#{id}"))

  deviceList: () ->
    @_list("device")

  deviceGet: (id) ->
    @_get("device", id)

  deviceAdd: (model) ->
    @_add("device", model)

  deviceUpdate: (id, model) ->
    @_update("device", id, model)

  deviceRemove: (id) ->
    @_remove("device", id)

  channelList: () ->
    @_list("channel")

  channelGet: (id) ->
    @_get("channel", id)

  channelAdd: (model) ->
    @_add("channel", model)

  channelUpdate: (id, model) ->
    @_update("channel", id, model)

  channelRemove: (id) ->
    @_remove("channel", id)

  triggerList: () ->
    @_list("trigger")

  triggerGet: (id) ->
    @_get("trigger", id)

  triggerAdd: (model) ->
    @_add("trigger", model)

  triggerUpdate: (id, model) ->
    @_update("trigger", id, model)

  triggerRemove: (id) ->
    @_remove("trigger", id)

  setValue: (id, value) ->
    @get apiPath("setValue/#{id}/#{value}")

  getValue: (id) ->
    @get apiPath("getValue/#{id}")

  channelStates: () ->
    @get apiPath("channelStates")

  measurements: (chId) ->
    @get apiPath("measurements/#{chId}")
