const puppeteer = require('puppeteer');

const query = 'colton morris'
query.replace(' ', '+');

const queryUrl = `https://www.google.com/search?q=${query}`

const browser = await puppeteer.launch();
const page = await browser.newPage();
await page.goto(queryUrl);

const allResultsSelector = '.resultStats';
await page.waitForSelector(allResultsSelector);

const links = await page.evaluate(resultsSelector => {
  const anchors = Array.from(document.querySelectorAll(allResultsSelector));

  console.log('anchors: ',anchors);
  return anchors.map(anchor => {
    const title = anchor.textContent.split('|')[0].trim();
    return `${title} - ${anchor.href}`;
  });
}, resultsSelector);
console.log('right here: ',links.join('\n'));

await browser.close();
