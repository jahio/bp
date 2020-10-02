import React, { useEffect, useState } from 'react';
import Chart from 'react-apexcharts'
import './App.css';
import ApiService from './service'



function Graph(props) {
  const [entries, setEntries] = useState([]);

  useEffect(() => {
    const getEntries = async () => {
      const e = await ApiService.getAllEntries()
      setEntries(e)
    }
    getEntries()
  }, [])

  let state = {
    options: {
      chart: {
        id: 'BloodPressure'
      },
      xaxis: {
        categories: []
      }
    },
    series: [
      {
        name: "systolic",
        data: []
      },
      {
        name: "diastolic",
        data: []
      },
      {
        name: "heartrate",
        data: []
      }
    ]
  }

  // Cycle through the entries and update the graph state
  entries.forEach((e) => {
    // Append to the x axis
    let at = new Date(e.created_at * 1000).toLocaleString(navigator.languages[0])
    state.options.xaxis.categories.push(at)

    // Append to the measurement arrays
    state.series.forEach((series) => {
      switch(series.name) {
        case "systolic":
          series.data.push(e.systolic)
          break;
        case "diastolic":
          series.data.push(e.diastolic)
          break;
        case "heartrate":
          series.data.push(e.heartrate)
          break;
        default:
          console.log(e)
          break;
      }
    })
  })

  return(
    <div id="graph">
      <Chart options={state.options} series={state.series} type="line" width="100%" height="500" />
    </div>
  )
}

export default Graph;
