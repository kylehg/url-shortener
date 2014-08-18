(function () {

// Constants
var POST_URL = '/'
var ROOT_URL = 'http://shawty.go/'
var SHORTCODE_REGEX = /^[A-Za-z0-9_-]+$/
var $form = $('#shorten-form')
var $urlField = $('#url')
var $codeField = $('#code')
var $resultField = $('#result')


// Handle a form submission
function onFormSubmit(e) {
  e.preventDefault()

  if (!isFormValid()) {
    return
  }

  var url = $urlField.trim().val()
  var code = $codeField.trim().val()
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
    $codeField.addClass('error').focus()
  } else if (jqXHR.status == 200 || jqXHR.status == 201) {
    showSuccess('URL shortened!')
    $resultField.val(POST_URL + data.code).focus()
  } else {
    var errMsg = data && data.meta && data.meta.error_text
    showError(errMsg || 'An unkown error occurred')
  }
}

// Validate the form
function isFormValid() {
  var code = $codeField.val().trim()
  if (!code || SHORTCODE_REGEX.test) {
    $codeField.removeClass('error')
  } else {
    showError('Shortcode can only contain [A-Za-z0-9_-]')
    $codeField.addClass('error')
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
  $form.on('submit', onFormSubmit)
  $codeField.on('blur', isFormValid)
  $urlField.on('blur', isFormValid)
})

}())
