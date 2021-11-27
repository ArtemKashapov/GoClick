console.log('This project is made by Artem and Anton ;-)')

const btn = document.getElementById("btn-click")
const textCounter = document.getElementById("text-count")
const body = document.getElementById('body')

let counter = 0

btn.addEventListener('click', e => {
    counter += 1

    if (counter < 10) {
        textCounter.textContent = "Вы кликнули всего-лишь " + String(counter) + " раз"
    } else {
        textCounter.textContent = "Вы кликнули целых " + String(counter) + " раз"
        body.setAttribute('class', 'lighter')
    }
    // console.log('btn clicked: ', counter)
})