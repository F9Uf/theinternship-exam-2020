const axios = require('axios')
const cheerio = require('cheerio')

// fetch response from url
const fetchHTML = async (url) => {
    const res = await axios.get(url)
    return res.data
}

const getResults = async (url) => {
    const data = await fetchHTML(url)
    const $ = cheerio.load(data)
    const companies = []

    // select div that has "partner" class and loop for each children node
    $("div.partner").each((i, e) => {
        const logoBoxNode = e.firstChild
        const linkLogoNode = logoBoxNode.lastChild
        const imgLogoNode = linkLogoNode.firstChild
        const textBoxNode = e.lastChild

        
        const linkCompany = linkLogoNode.attribs.href
        const logoCompany = imgLogoNode.attribs.src
        const descriptionCompany = $(textBoxNode.firstChild).text()
        
        companies.push({
            link: linkCompany,
            logo: logoCompany,
            description: descriptionCompany
        })
    })

    // return the companies array that is sorted by description length
    return companies.sort((a, b) => {
        return a.description.length - b.description.length
    })
}

module.exports = getResults;