const express = require('express')
const getResult = require('./scraping')

const app = express()

const siteUrl = "https://theinternship.io/"

// simple api
app.get('/', (req, res) => {
    res.json({
        message: 'ok'
    })
})

app.get('/companies', async (req, res) => {
    const result = await getResult(siteUrl)

    // map only logo attribute in each company
    // add website path prefix to logo url
    const companies = result.map(
        ({ logo }) => ({ logo: `${siteUrl}${logo}`})
    )
    res.json({
        companies
    })
})

// listening server
app.listen(3000, () => {
    console.log('app listening on port 3000')
})