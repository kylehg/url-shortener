(function (){

// Constants
var POST_URL = '/'
var ROOT_URL = 'http://shawty.go/'
var SHORTCODE_REGEX = /^[A-Za-z0-9_-]+$/
var $FORM = $('#shorten-form')
var $URL_FIELD = $('#url')
var $CODE_FIELD = $('#code')
var $RESULT_FIELD = $('#result')


// Handle a form submission
function onFormSubmit(e) {
  e.preventDefault()

  if (!isFormValid()) {
    return
  }

  var url = $URL_FIELD.trim().val()
  var code = $CODE_FIELD.trim().val()
  $.ajax({
    type: 'POST',
    url: POST_URL + code,
    data: {url: url},
    success: onResponse,
    dataType: 'json'
  })
}

// Handle a response from the server
function onResponse(data, textStatus, jqXHR) {
  if (jqXHR.status == 303) {
    showError('Sorry, that code is already taken.')
    $CODE_FIELD.addClass('error').focus()
  } else if (jqXHR.status == 200 || jqXHR.status == 201) {
    showSuccess('URL shortened!')
    $RESULT_FIELD.val(POST_URL + data.code).focus()
  } else {
    var errMsg = data && data.meta && data.meta.error_text
    showError(errMsg || 'An unkown error occurred')
  }
}

// Validate the form
function isFormValid() {
  var code = $CODE_FIELD.val().trim()
  if (!code || SHORTCODE_REGEX.test) {
    $CODE_FIELD.removeClass('error')
  } else {
    showError('Shortcode can only contain [A-Za-z0-9_-]')
    $CODE_FIELD.addClass('error')
  }
}

function showError(msg) {
  console.error(msg)
}

function showSuccess(msg) {
  console.log('success', msg)
}


// Kick things off
$(function () {
  // Initialize event handlers
  $FORM.on('submit', onFormSubmit)
  $CODE_FIELD.on('blur', isFormValid)
  $URL_FIELD.on('blur', isFormValid)
})

}())
