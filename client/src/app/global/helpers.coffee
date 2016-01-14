
endsWith = (str, suffix) ->
  return str.indexOf(suffix, str.length - suffix.length) != -1

String::endsWith = endsWith
