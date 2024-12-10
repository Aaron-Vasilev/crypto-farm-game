window.vars = {
  selectedPot: null,
}

function openModal() {
  document.querySelector('dialog').showModal()
}

function closeModal() {
  document.querySelector('dialog').close()
}

function updateHoursValue() {
  const dialog = document.querySelector('dialog')

  const count = dialog.querySelector('#modal-hours-range').value
  const span = dialog.querySelector('#modal-hours-span')

  span.innerHTML = count
}

function getSelectedPotId() {
  return window.vars.selectedPot
}
