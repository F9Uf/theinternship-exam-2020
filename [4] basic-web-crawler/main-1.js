const getResults = require('./scraping')

const main = async () => {
    const siteUrl = "https://theinternship.io/"
    const companies = await getResults(siteUrl)
    
    // loop to display of logo of company    
    for (let i = 0; i < companies.length; i++) {
        const { logo } = companies[i]
        console.log(logo)
    }
}

main()

