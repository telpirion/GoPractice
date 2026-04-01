"use strict";

window.onload = function () {
    console.log("Ready to submit a map!");
    let agreeTOS = document.getElementById("chx-tos");
    let submitBtn = document.getElementById("btn-submit");
    let fileUpload = document.getElementById("file-upload");
    agreeTOS.addEventListener("change", (e) => {
        if (e.target.checked) {
            submitBtn.removeAttribute("disabled")
        } else {
            submitBtn.setAttribute("disabled", true);
        }
    });
    fileUpload.addEventListener("change", (e) => {
        if (e.target.files.length > 0) {
            let fileName = e.target.files[0].name;
            let lblFileUpload = document.getElementById("lbl-file-upload");
            lblFileUpload.textContent = fileName;

            document.getElementById("file-name").value = fileName;
        }
    })
}