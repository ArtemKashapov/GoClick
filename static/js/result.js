let tmpCounter = 0
document.getElementById('download-btn').addEventListener('click', e => {
    document.getElementById('download-btn').setAttribute('href', '#')
   if (tmpCounter == 0) {
       alert('А волшебное слово?')
   } else if (tmpCounter == 1) {
       alert('Попробуйте нажать еще раз 5, тогда точно все получится!')
   } else if (tmpCounter > 5) {
       document.getElementById('download-btn').setAttribute('href', '/') // TODO: Вставить ссылку на скачивание!
   }
   
   tmpCounter += 1
})