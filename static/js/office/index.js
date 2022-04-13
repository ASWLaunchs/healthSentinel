const Index = {
    data: {
        isExist: false,
        completedValue: 0,
    },
    func: {
        imgUpload: function (event) {
            event.target.nextElementSibling.click()
        },
        progresss: function (isAdded) {
            let progress = document.querySelector("#progress")
            if (isAdded) {
                Index.data.completedValue += 25
            } else if (Index.data.completedValue > 0) {
                Index.data.completedValue -= 25
            }

            if (Index.data.completedValue == 0) {
                progress.classList.remove("bg-success")
            } else if (!progress.classList.contains("bg-success")) {
                progress.classList.add("bg-success")
            }

            progress.style.width = Index.data.completedValue + "%"
        },
        isSnoExist: function (event) {
            console.log(event.target.value)
            if (event.target.value.length > 0 && event.target.value != null) {
                Index.data.isSnoExist = true
            }
        },
        sno: function (event) {
            if (event.target.value.length > 0 && event.target.value != null) {
                event.target.classList.remove("is-invalid")
                event.target.classList.add("is-valid")
                if (!Index.data.isSnoExist) {
                    Index.func.progresss(true)
                }
            } else {
                event.target.classList.remove("is-valid")
                event.target.classList.add("is-invalid")
                Index.func.progresss(false)
                Index.data.isSnoExist = false
            }
        },
        previewFile: function (event) {
            let preview = document.querySelector("img#" + event.target.name);
            let file = document.querySelector('input[name=' + event.target.name + ']').files[0];
            let reader = new FileReader();

            reader.onloadend = function () {
                preview.src = reader.result;
                if (event.target.value.length > 0) {
                    event.target.classList.remove("is-invalid")
                    event.target.classList.add("is-valid")
                    event.target.style.backgroundColor = "rgb(172, 255 ,172)"
                    Index.func.progresss(true)
                } else {
                    event.target.classList.remove("is-valid")
                    event.target.classList.add("is-invalid")
                    event.target.style.backgroundColor = ""
                    Index.func.progresss(false)
                }
            }

            if (file) {
                reader.readAsDataURL(file);
            } else {
                preview.src = "/static/assets/images/nothing.svg";
                event.target.classList.remove("is-valid")
                event.target.classList.add("is-invalid")
                event.target.style.backgroundColor = ""
                Index.func.progresss(false)
            }
        },
        //check sno is incorrect or not.
        check: function (event) {
            let elemet = event.target
            $.get("/check", {
                sno: elemet.value,
            }, function (data, status) {
                if (status) {
                    data.status ? $("#checkRes").fadeIn() : $("#checkRes").fadeOut()
                }
            })
        },
    }
}