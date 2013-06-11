shty = shty || {}

# Kick things off
shty.init = ->
  $('#shorten-form').on 'submit', shty.onFormSubmit

# Handle a form submission
# @param {Event} e
shty.onFormSubmit = (e) ->
  e.preventDefault()
  data =
    url: $('#url-field').val()
    code: $('#shortcode-field').val()

  if not data.url
    shty.handleNoUrl()

  # TODO: Left off here


# Handle a form missing a URL
shty.handleNoUrl = -> # TODO



$ shty.init
