// Imports the Google Cloud client library
const BigQuery = require('@google-cloud/bigquery')
const axios = require('axios')

const projectId = 'alex-olivier'
const datasetId = 'flighttracker_dev'
const tableId = 'locations'

// Creates a client
const bigquery = new BigQuery({
  projectId: projectId,
  keyFilename: 'google.json'
})

async function run () {
  try {
    const { data } = await axios.get('http://192.168.3.175:8080/data/aircraft.json')
    const rows = data.aircraft.map((input) => {
      input.ts = Math.floor(data.now)
      if (input.flight) input.flight = input.flight.trim()
      return input
    })
    // Inserts data into a table
    await bigquery
    .dataset(datasetId)
    .table(tableId)
    .insert(rows)

    console.log(`Inserted ${rows.length} rows`)
  } catch (err) {
    if (err && err.name === 'PartialFailureError') {
      if (err.errors && err.errors.length > 0) {
        console.log('Insert errors:')
        err.errors.forEach(err => console.error(err))
      }
    } else {
      console.error('ERROR:', err)
    }
  }
}
setInterval(run, 5000)
