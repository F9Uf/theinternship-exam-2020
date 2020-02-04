const express = require('express')
const getResult = require('./scraping')

const app = express()

const siteUrl = "https://theinternship.io/"

app.get('/', (req, res) => {
    res.json({
        message: 'ok'
    })
})

app.get('/companies', async (req, res) => {
    const result = await getResult(siteUrl)
    const companies = result.map(
        ({ logo }) => ({ logo: `${siteUrl}${logo}`})
    )
    res.json({
        companies
    })
})

app.listen(3000, () => {
    console.log('app listening on port 3000')
})