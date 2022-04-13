const Feedback = {
    data: {

    },
    func: {
        feedback: function () {
            let feedBack = document.querySelector("#feedback")
            let container = document.querySelector(".container")
            feedBack.style.transform = "translateY(0%)"
            container.style.opacity = '0.2'
        },
        feedbackClose: function () {
            let feedBack = document.querySelector("#feedback")
            let container = document.querySelector(".container")
            feedBack.style.transform = "translateY(120%)"
            container.style.opacity = '1'
        }
    }
}