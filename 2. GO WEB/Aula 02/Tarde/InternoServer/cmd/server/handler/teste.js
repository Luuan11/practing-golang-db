'use strict';

const fs = require('fs');
const https = require('https');

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', function(inputStdin) {
    inputString += inputStdin;
});

process.stdin.on('end', function() {
    inputString = inputString.split('\n');

    main();
});

function readLine() {
    return inputString[currentLine++];
}

async function getCountryName(code) {
    // write your code here
    // API endpoint: https://jsonmock.hackerrank.com/api/countries?page=<PAGE_NUMBER>
    
    try {
        let page = 1;
        let countryName = null;
        while (!countryName) {
            const data = await fetchData(`https://jsonmock.hackerrank.com/api/countries?page=${page}`);
            const countries = data.data;
            const foundCountry = countries.find(country => country.alpha2Code === code);
            if (foundCountry) {
                countryName = foundCountry.name;
            } else {
                if (page >= data.total_pages) {
                    // If the code is not found and there are no more pages, break the loop
                    break;
                }
                page++;
            }
        }
        return countryName;
    } catch (error) {
        console.error('Error fetching country data:', error);
        return null;
    }
}

function fetchData(url) {
    return new Promise((resolve, reject) => {
        https.get(url, (res) => {
            let data = '';
            res.on('data', (chunk) => {
                data += chunk;
            });
            res.on('end', () => {
                resolve(JSON.parse(data));
            });
        }).on('error', (error) => {
            reject(error);
        });
    });
}

async function main() {
  const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

  const code = readLine().trim();

  const name = await getCountryName(code);

  ws.write(`${name}\n`);

}
