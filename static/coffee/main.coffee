sty = sty || {}

sty.CODE_FIELD_ID = '#code-field'
sty.SHORTEN_FORM_ID = '#shorten-form'
sty.URL_FIELD_ID = '#url-field'


# Kick things off
sty.init = ->
  $(sty.SHORTEN_FORM_ID).on 'submit', sty.onFormSubmit
  $(sty.CODE_FIELD_ID).on 'blur', sty.validForm
  $(sty.URL_FIELD_ID).on 'blur', sty.validForm


# Handle a form submission
# @param {Event} e
sty.onFormSubmit = (e) ->
  e.preventDefault()

  if (!sty.validForm())
    return false

  data =
    url: $(sty.SHORTEN_FORM_ID).val()
    code: $(sty.CODE_FIELD_ID).val()

  $.post(sty.POST_URL, data, sty.onResponse)

  return true


sty.onResponse = (resp) ->
  # TODO

sty.validForm = ->
  # TODO


$ sty.init
