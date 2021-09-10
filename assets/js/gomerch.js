function openDeleteConfirmation(id, name) {
    currentProduct = { id, name };
    $('#deleteConfirmationModalMessage').html(`Are you sure you want to delete <b>${currentProduct.name}</b>?`);
    $('#deleteConfirmationModal').modal({});
}

function onDelete(id) {
    window.location = "/delete?id=" + id;
}
