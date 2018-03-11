const config = require('config')
const net = require('net')
const BigQuery = require('@google-cloud/bigquery')

// Creates a client
const bigquery = new BigQuery({
  projectId: config.google.projectId,
  keyFilename: 'google.json'
})

let batch = []

const client = new net.Socket()
client.connect(config.pi.port, config.pi.ip, function () {
  console.log('Connected')
})

client.on('data', function (data) {
  const rows = data.toString().split('\n')
  const formatted = rows.filter((a) => {
    return a !== ''
  }).map((r) => {
    const tsv = r.split('\t')
    let output = {}
    for (let i = 0; i < tsv.length; i = i + 2) {
      if (tsv[i + 1]) {
        if (tsv[i] === 'clock') tsv[i] = 'timestamp'
        output[tsv[i]] = tsv[i + 1]
      }
    }

    if (output.timestamp) output.timestamp = parseInt(output.timestamp)
    if (output.ident) output.ident = output.ident.trim()
    if (output.squawk) output.squawk = parseInt(output.squawk)
    if (output.alt) output.alt = parseInt(output.alt)
    if (output.speed) output.speed = parseInt(output.speed)
    if (output.lat) output.lat = parseFloat(output.lat)
    if (output.lon) output.lon = parseFloat(output.lon)
    if (output.heading) output.heading = parseInt(output.heading)

    return output
  })
  // await streamToBQ(formatted)
  batch = batch.concat(formatted.filter((f) => {
    return !!f.timestamp
  }))
})

client.on('close', function () {
  console.log('Connection closed')
})

async function streamToBQ () {
  try {
    const rows = batch.slice(0)
    batch = []
    await bigquery
      .dataset(config.google.datasetId)
      .table(config.google.tableId)
      .insert(rows)

    console.log(`${new Date()} - Inserted ${rows.length} rows`)
  } catch (err) {
    if (err && err.name === 'PartialFailureError') {
      if (err.errors && err.errors.length > 0) {
        console.log(`${new Date()} - Insert errors:`)
        err.errors.forEach(err => console.error(err))
      }
    } else {
      console.error(`${new Date()} - ERROR:`, err)
    }
  }
}

setInterval(streamToBQ, 5000)
